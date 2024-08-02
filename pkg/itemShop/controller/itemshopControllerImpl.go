package controller

import (
	"net/http"
	"github.com/Kotravel095/golang-backend-gameshop/pkg/custom"
	"github.com/labstack/echo/v4"

	_itemshopService "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/service"
    // _itemshopModel "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/model"
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
    // itemFilter := new()
    // pctx.Bind()

    itemModelList, err := c.itemshopService.Listing()
    if err != nil {
        return custom.Error(pctx, http.StatusInternalServerError, err.Error())
    }
    return pctx.JSON(http.StatusOK, itemModelList)
}
