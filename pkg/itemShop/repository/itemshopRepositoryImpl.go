package repository

import (
	"github.com/Kotravel095/golang-backend-gameshop/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemshopException  "github.com/Kotravel095/golang-backend-gameshop/pkg/itemShop/exception"
)

type ItemshopRepositoryImpl struct{
	db *gorm.DB
	logger echo.Logger
}

func NewItemshopRepositoryImpl(db *gorm.DB, Logger echo.Logger) ItemshopRepository {
	return &ItemshopRepositoryImpl{db, Logger}
}

func (r *ItemshopRepositoryImpl) Listing() ([]*entities.Item, error) {
	itemList := make([]*entities.Item, 0)

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Error("Fail to list items: %s", err.Error())
		return nil, &_itemshopException.ItemListing{}
	}

	return itemList, nil
}