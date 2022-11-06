package main

import (
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
	log.Println(db, err)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// POST / Restaurant
	v1 := r.Group("/v1")

	restaurant := v1.Group("/restaurants")

	restaurant.POST("", func(c *gin.Context) {
		var data Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Create(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

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

	restaurant.GET("", func(c *gin.Context) {
		var data []Restaurant

		db.Order("id desc").Find(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

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
