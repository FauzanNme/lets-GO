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

	students := []config.Student{}

	db.Find(&students)

	for _, data := range students {
		fmt.Println("ID:", data.ID)
		fmt.Println("Name:", data.Name)
	}
}
