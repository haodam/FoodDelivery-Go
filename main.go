package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
	//Phone string `json:"phone" gorm:"column:phone;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/fooddelivery?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//log.Println(db, err)
	db.AutoMigrate(&Restaurant{})

	newRestaurant := Restaurant{Name: "Trau Ngon Quan", Addr: "So 1 Nguyen Van Cu"}

	db.Table("restaurants").Create(&newRestaurant)
}
