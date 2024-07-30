package repository

type ItemshopRepositoryImpl struct{}

func NewItemshopRepositoryImpl() ItemshopRepository {
	return &ItemshopRepositoryImpl{}
}