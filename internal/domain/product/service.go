package internal

type Service interface {
	GetProducts() *[]Product
	GetProduct() *Product
	AddProduct() *Product
	UpdateProduct() *Product
	DeleteProduct() error
}

type service struct {
	storage Storage
}

func NewRepository(storage Storage) Service {
	return &service{storage: storage}
}

func (*service) AddProduct() *Product{
	return nil
}

func (s *service) DeleteProduct() error {
	return nil
}

func (s *service) GetProduct() *Product{
	return s.storage.GetProduct()
}

func (s *service) GetProducts() *[]Product {
	return s.storage.GetProducts()
}

func (s *service) UpdateProduct() *Product {
	return nil
}
