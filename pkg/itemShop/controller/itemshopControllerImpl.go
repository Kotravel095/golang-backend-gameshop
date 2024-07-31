package controller

import (
    "net/http"
    _itemshopService "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/service"
    "github.com/labstack/echo/v4"
)

type ItemShopControllerImpl struct {
    itemshopService _itemshopService.ItemshopService
}

func NewItemshopControllerImpl(
    itemshopService _itemshopService.ItemshopService,
) ItemShopController {
    return &ItemShopControllerImpl{itemshopService}
}

func (c *ItemShopControllerImpl) Listing(pctx echo.Context) error {
    itemModelList, err := c.itemshopService.Listing()
    if err != nil {
        return pctx.String(http.StatusInternalServerError, err.Error())
    }
    return pctx.JSON(http.StatusOK, itemModelList)
}
