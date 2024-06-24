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

	student := []config.Student{}

	db.Find(&student)

	for _, data := range student {
		fmt.Println("ID:", data.ID)
	}
}