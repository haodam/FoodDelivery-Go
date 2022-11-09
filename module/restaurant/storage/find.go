package restaurantstorage

import (
	restaurantmodel "FoodDelivery/module/restaurant/model"
	"context"
)

func (s *sqlStore) FindRestaurantWithCondition(ctx context.Context, condition map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {

	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}