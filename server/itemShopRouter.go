package server

import (
    _itemshopService "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/service"
    _itemshopRepository "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/repository"
    _itemshopController "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/controller"
)

func (s *echoServer) initItemShopRouter() {
    router := s.app.Group("/v1/item-shop")

    itemshopRepository := _itemshopRepository.NewItemshopRepositoryImpl(s.db, s.app.Logger)
    itemshopService := _itemshopService.NewItemshopServiceImpl(itemshopRepository)
    itemshopController := _itemshopController.NewItemshopControllerImpl(itemshopService)

    router.GET("/listing", itemshopController.Listing)
}
