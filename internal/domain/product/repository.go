package internal

type ProductStorage interface {
	GetAll() *[]Product
	GetById(id string) *Product
	Add(newProduct Product) *Product
	// UpdateProduct() *Product
	DeleteById(id string) error
}

type InMemoryStorage struct {
	data []Product
}

func (s *InMemoryStorage) GetAll() *[]Product {
	return &s.data
}

func (s *InMemoryStorage) GetById(id string) *Product {
	for _, product := range s.data {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func (s *InMemoryStorage) Add(newProduct Product) *Product {
	s.data = append(s.data, newProduct)
	return nil
}

func (s *InMemoryStorage) DeleteById(id string) error {
	for index, product := range s.data {
		if product.ID == id {
			s.data = append(s.data[:index], s.data[index+1:]...)
		}
	}

	return nil
}

func NewInMemoryStorage() ProductStorage {
	return &InMemoryStorage{
		data: []Product{
			{
				ID: "1",
				Cover: []Cover{
					{
						Url:  "http://localhost:8090/shop/demo/product_2.png",
						Type: "photo",
					},
					{
						Url:  "http://localhost:8090/shop/demo/product_9.png",
						Type: "photo",
					},
					{
						Url:  "http://localhost:8090/shop/demo/product_7.png",
						Type: "photo",
					},
				},
				Title:       "Золоте кольє",
				Description: "Романтичний та тендітний акцент твоїх образів.",
				Price:       1699,
			},
			{
				ID: "2",
				Cover: []Cover{
					{
						Url:  "http://localhost:8090/shop/demo/product_2.png",
						Type: "photo",
					},
					{
						Url:  "http://localhost:8090/shop/demo/product_9.png",
						Type: "photo",
					},
					{
						Url:  "http://localhost:8090/shop/demo/product_7.png",
						Type: "photo",
					},
				},
				Title:       "Золоте кольє",
				Description: "Романтичний та тендітний акцент твоїх образів.",
				Price:       1699,
			},
			{
				ID: "3",
				Cover: []Cover{
					{
						Url:  "http://localhost:8090/shop/demo/product_2.png",
						Type: "photo",
					},
					{
						Url:  "http://localhost:8090/shop/demo/product_9.png",
						Type: "photo",
					},
					{
						Url:  "http://localhost:8090/shop/demo/product_7.png",
						Type: "photo",
					},
				},
				Title:       "Золоте кольє",
				Description: "Романтичний та тендітний акцент твоїх образів.",
				Price:       1699,
			},
		},
	}
}
