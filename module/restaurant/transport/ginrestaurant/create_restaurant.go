package ginrestaurant

import (
	"FoodDelivery/common"
	"FoodDelivery/component/appctx"
	restaurantbiz "FoodDelivery/module/restaurant/biz"
	restaurantmodel "FoodDelivery/module/restaurant/model"
	restaurantstorage "FoodDelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMaiDBConnection()
		go func() {
			defer common.AppRecover()
			arr := []int{}
			log.Println(arr[0])
		}()
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		//db.Create(&data)
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
