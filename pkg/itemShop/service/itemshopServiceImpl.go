package service

import (
	_itemshopRepository "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/repository"
)

type itemshopServiceImpl struct{
	itemshopRepository _itemshopRepository.ItemshopRepository
}

func NewItemshopServiceImpl(itemshopRepository _itemshopRepository.ItemshopRepository) ItemshopService {
	return &itemshopServiceImpl{itemshopRepository}
}