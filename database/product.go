package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

var ProductList []Product

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

	ProductList = append(ProductList, pd1, pd2, pd3)
}
