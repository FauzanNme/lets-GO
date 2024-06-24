package config

import (
	"time"

	"gorm.io/driver/postgres"
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
	// dsn := "host=localhost user=root password= dbname=learning port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=root password= dbname=learning port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Somethign wrong i can feel it", err)
		return nil
	}

	fmt.Println("Database connected")

	db.AutoMigrate(&Student{})
	db.Debug()

	return db
}
