package restaurantmodel

import (
	"FoodDelivery/common"
	"errors"
	"strings"
)

type RestaurantType string

const (
	TypeNormal  RestaurantType = "normal"
	TypePremium RestaurantType = "premium"
	EntityName                 = "Restaurant"
)

type Restaurant struct {
	common.SQLModel
	Name  string         `json:"name" gorm:"column:name;"`
	Addr  string         `json:"addr" gorm:"column:addr;"`
	Type  RestaurantType `json:"type" gorm:"column:type;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Image  `json:"cover" gorm:"column:cover;"`
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel
	Name  string        `json:"name" gorm:"column:name;"`
	Addr  string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image `json:"logo" gorm:"column:logo;"`
	Cover *common.Image `json:"cover" gorm:"column:cover;"`
}

func (data *RestaurantCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name  *string       `json:"name" gorm:"column:name;"`
	Addr  *string       `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image `json:"logo" gorm:"column:logo;"`
	Cover *common.Image `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
