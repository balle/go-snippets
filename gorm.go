package main

/*
Example code for using the Object Relation Mapper GORM
Code is borrowed from the official documention found on
https://gorm.io/docs/

Use https://github.com/smallnest/gen to generate GORM models
for an exiting database
*/

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

type Product struct {
	gorm.Model
	Name       string
	CategoryID int
	Category   Category
	Price      uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{}, &Category{})

	// Create
	category := Category{Name: "Toy"}
	db.Create(&category)
	db.Create(&Product{Name: "Super Mario",
		Category: category,
		Price:    42})

	// Read
	var product Product
	db.First(&product, 1)                         // find product with integer primary key
	db.First(&product, "name = ?", "Super Mario") // find product with code D42

	var products []Product
	db.Find(&products)

	for _, product := range products {
		fmt.Printf("%v\n", product)
	}

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Name: "Super Metroid"}) // non-zero fields

	// Delete - delete product
	db.Delete(&product, 1)
}
