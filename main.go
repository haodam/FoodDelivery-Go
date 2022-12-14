package main

import (
	"FoodDelivery/component/appctx"
	"FoodDelivery/middleware"
	"FoodDelivery/module/restaurant/transport/ginrestaurant"
	"FoodDelivery/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	//dsn := "root:@tcp(127.0.0.1:3306)/fooddelivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(db, err)

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/static", "./static")

	// POST / Restaurant
	v1 := r.Group("/v1")
	v1.POST("upload", ginupload.UploadImage(appContext))
	restaurant := v1.Group("/restaurants")

	restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))

	//GET BY ID
	restaurant.GET("/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data Restaurant

		db.Where("id =?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//GET LIST
	restaurant.GET("", ginrestaurant.ListRestaurant(appContext))

	//UPDATE
	restaurant.PATCH("/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Where("id =?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//DELETE
	restaurant.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//newRestaurant := Restaurant{Name: "Trau Ngon Quan", Addr: "So 1 Tuw Son"}
//if err := db.Create(&newRestaurant).Error; err != nil {
//	log.Fatalln(err)
//}
//
//var myRestaurant Restaurant
//
//if err := db.Where("id=?", 1).First(&myRestaurant).Error; err != nil {
//	log.Fatalln(err)
//}
//log.Println(myRestaurant)

//myRestaurant.Name = "Trau Tra Giang"

//newName := ""
//updateData := RestaurantUpdate{Name: &newName}
//
//if err := db.Where("id=?", 1).Updates(&updateData).Error; err != nil {
//	log.Fatalln(err)
//}
//log.Println(myRestaurant)
//
//if err := db.Table(Restaurant{}.TableName()).Where("id=?", 2).Delete(nil).Error; err != nil {
//	log.Fatalln(err)
//}
//log.Println(myRestaurant)

//}
