package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList

	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}

	return &product, nil
}

func generateInitialProducts(r *productRepo) {
	pd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       120,
		ImgUrl:      "http://orange.com",
	}

	pd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green",
		Price:       120,
		ImgUrl:      "http://apple.com",
	}

	pd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       120,
		ImgUrl:      "http://banana.com",
	}

	r.productList = append(r.productList, pd1, pd2, pd3)
}
