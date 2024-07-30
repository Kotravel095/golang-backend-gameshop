package controller

import (
	_itemshopService "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/service"
)

type ItemShopControllerImpl struct{
	itemshopService _itemshopService.ItemshopService
}

func NewItemshopService(itemshopService _itemshopService.ItemshopService) ItemShopController {
	return &ItemShopControllerImpl{itemshopService}
}