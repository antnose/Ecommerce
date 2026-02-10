package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

var productList []Product

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       120,
		ImgURL:      "http://orange.com",
	}

	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green",
		Price:       120,
		ImgURL:      "http://apple.com",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       120,
		ImgURL:      "http://banana.com",
	}

	productList = append(productList, pd1, pd2, pd3)
}
