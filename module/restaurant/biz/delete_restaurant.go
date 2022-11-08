package restaurantbiz

import "context"

type DeleteRestaurant interface {
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurant
}

func NewDeleteRestaurant(store DeleteRestaurant) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	if err := biz.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
