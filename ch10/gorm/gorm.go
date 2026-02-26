package main

import (
	"gorm.io/gorm"
)

//升级包：go get -u 升级到最新的次要版本或者修订版本
//go get -u=patch 升级到最新的修订版本
//go get xxx

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//func main() {
//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	ctx := context.Background()
//
//	// Migrate the schema
//	db.AutoMigrate(&Product{})
//
//	// Create
//	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
//
//	// Read
//	product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx)       // find product with integer primary key
//	products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // find product with code D42
//
//	// Update - update product's price to 200
//	err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
//	// Update - update multiple fields
//	err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})
//
//	// Delete - delete product
//	err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
//}
