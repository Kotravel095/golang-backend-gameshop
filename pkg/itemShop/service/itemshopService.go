package service

import (
    _itemShopModel "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/model"
)

type ItemshopService interface {
    Listing() ([]*_itemShopModel.Item, error)
}
