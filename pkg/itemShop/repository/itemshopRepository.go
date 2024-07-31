package repository

import "github.com/Kotravel095/golang-backend-gameshop/entities"

type ItemshopRepository interface{
	Listing() ([]*entities.Item, error)
}