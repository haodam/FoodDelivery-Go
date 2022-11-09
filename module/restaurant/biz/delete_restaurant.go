package restaurantbiz

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
	"errors"
)

type DeleteRestaurant interface {
	Delete(ctx context.Context, id int) error
	FindRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurant
}

func NewDeleteRestaurant(store DeleteRestaurant) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {

	oldData, err := biz.store.FindRestaurantWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
