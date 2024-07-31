package service

import (
    _itemshopRepository "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/repository"
    _itemShopModel "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/model"
)

type itemshopServiceImpl struct {
    itemshopRepository _itemshopRepository.ItemshopRepository
}

func NewItemshopServiceImpl(
    itemshopRepository _itemshopRepository.ItemshopRepository,
) ItemshopService {
    return &itemshopServiceImpl{itemshopRepository}
}

func (s *itemshopServiceImpl) Listing() ([]*_itemShopModel.Item, error) {
    itemList, err := s.itemshopRepository.Listing()
    if err != nil {
        return nil, err
    }

    itemModelList := make([]*_itemShopModel.Item, 0)
    for _, item := range itemList {
        itemModelList = append(itemModelList, item.ToItemModel())
    }
    return itemModelList, nil
}
