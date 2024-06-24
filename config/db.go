package config

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
)

type Student struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(120)"`
	Age       int8
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt
}

func NewDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go-sql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Somethign wrong i can feel it", err)
		return nil
	}

	fmt.Println("Database connected")

	db.AutoMigrate(&Student{})
	db.Debug()

	return db
}
