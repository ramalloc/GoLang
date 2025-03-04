package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
    dsn := "youruser:yourpassword@tcp(localhost:3306)/yourdatabase?charset=utf8&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database: " + err.Error())
    }
    fmt.Println("✅ Database connected successfully!")
    db = database
}

func GetDB() *gorm.DB {
    return db
}
