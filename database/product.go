package database

var productList []Product

func List() []Product {
	return productList
}

func Get(productId int) *Product {
	for _, product := range productList {
		if product.ID == productId {
			return &product
		}
	}
	return nil
}

func Store(product Product) Product {
	product.ID = len(productList) + 1
	productList = append(productList, product)
	return product
}

func Update(product Product) {
	for idx, prod := range productList {
		if prod.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product
	for _, product := range productList {
		if product.ID != productId {
			tempList = append(tempList, product)
		}
	}
	productList = tempList
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Apple iPhone 14",
		Description: "The latest iPhone with advanced features.",
		Price:       999.99,
		ImageUrl:    "https://example.com/iphone14.jpg",
	}
	product2 := Product{
		ID:          2,
		Title:       "Samsung Galaxy S22",
		Description: "A powerful smartphone with a stunning display.",
		Price:       899.99,
		ImageUrl:    "https://example.com/galaxys22.jpg",
	}
	product3 := Product{
		ID:          3,
		Title:       "Google Pixel 6",
		Description: "Experience the best of Google with the Pixel 6.",
		Price:       599.99,
		ImageUrl:    "https://example.com/pixel6.jpg",
	}
	product4 := Product{
		ID:          4,
		Title:       "OnePlus 9",
		Description: "Flagship performance at an affordable price.",
		Price:       729.99,
		ImageUrl:    "https://example.com/oneplus9.jpg",
	}

	productList = append(productList, product1)
	productList = append(productList, product2)
	productList = append(productList, product3)
	productList = append(productList, product4)
}
