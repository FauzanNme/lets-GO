package main

import (
	"fmt"
	"lets-go/config"
)

func main() {
	fmt.Println("connecting")
	db := config.NewDatabase()

	if db == nil {
		fmt.Println("failed to connect database")
		return
	}

	fmt.Println("Main func has been trigerred")

	data := config.Student{}

	db.Unscoped().Where("id = ?", 1).Delete(&data)
	db.Where("Name = ?", "Afif").Delete(&data)
}