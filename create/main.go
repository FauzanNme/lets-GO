package main

import (
	"fmt"
	"lets-go/config"
)

func main() {
	fmt.Println("Connecting to database...")
	db := config.NewDatabase()

	if db == nil {
		fmt.Println("Failed to connect to database")
		return
	}

	fmt.Println("Main function has been triggered")

	students := []config.Student{
		{
			Name: "Afif",
			Age:  17,
		},
		{
			Name: "Fauzan",
			Age:  10,
		},
		{
			Name: "Apip",
			Age:  19,
		},
	}
	db.Create(&students)

	fmt.Println("Data has been created")
}
