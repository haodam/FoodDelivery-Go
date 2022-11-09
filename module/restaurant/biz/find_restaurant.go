package restaurantbiz

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

type FindRestaurantStore interface {
	FindRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFindRestaurant(store FindRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}
