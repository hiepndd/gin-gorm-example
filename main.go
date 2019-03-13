package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//http://mindbowser.com/golang-go-with-gorm-2/
// Customer contain infor of Customer
type Customer struct {
	gorm.Model
	CustomerID   int `gorm:"primary_key"`
	CustomerName string
	Contacts     []Contact `gorm:"ForeignKey:CustID"`
}

// Contact contain info contact of Customer
type Contact struct {
	gorm.Model
	ContactID   int `gorm:"primary_key"`
	CountryCode int
	MobileNo    uint
	CustID      int
}

func main() {

	db, err := gorm.Open("mysql", "localhost:123456@/example?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

}
