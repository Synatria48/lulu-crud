package crud

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *sql.DB {
	var err error
	// db, err = gorm.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open(mysql.Open("lulu:lulu123@tcp(localhost:3306)/db_lulu?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic("Failed to connect to database")
	}

	sql, err := db.DB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		panic("Failed to connect to database")
	}

	// AutoMigrate will attempt to automatically migrate the schema, creating the table "users" if it doesn't exist
	db.AutoMigrate(&User{})

	return sql
}
