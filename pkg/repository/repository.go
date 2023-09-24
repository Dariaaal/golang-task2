package repository

type ProductsList interface {

}

type Repository struct {
	ProductsList
}

func NewRepository() *Repository {
	return &Repository {
		
	}
}