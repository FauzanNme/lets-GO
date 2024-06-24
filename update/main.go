package main

import (
	"fmt"
	"lets-go/config"
)

func main() {
	fmt.Println("Connecting to database...")
	db := config.NewDatabase()

	if db == nil {
		fmt.Println("Failed connecting to database")
		return
	}

	fmt.Println("Main function has been trigerred.")

	db.Where("id IN ?", []uint{3}).Updates(&config.Student{
		Name: "Neir",
	})
}