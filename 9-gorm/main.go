package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type Product struct {
    ID uint
    Code string
    Price uint
}

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&Product{})

    // Create
    db.Create(&Product{Code: "D42", Price: 100})

    // Read
    var product Product
    db.First(&product, 1)
    db.First(&product, "code = ?", "D42")

    db.Model(&product).Updates(Product {
        Price: 200,
        Code: "F42",
    })

    // Delete - delete product
    db.Delete(&product, 1)
}
