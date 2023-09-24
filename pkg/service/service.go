package service

import "github.com/dariaaal/golang-task2/pkg/repository"

type ProductsList interface {
}

type Service struct {
	ProductsList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
