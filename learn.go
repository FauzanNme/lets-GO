package main

import(
	"fmt"
	"time"

	"gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
	ID uint
	Name string
	Age int
	CreatedAt time.Time
	UpdateAt time.Time
}

type Student struct {
	Name string
	Age int
	Class string
	CreatedAt time.Time
	UpdateAt time.Time
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-sql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Something wrong i can feel it", err)
	}

	fmt.Println("Database connected")

	db.AutoMigrate(&User{})

	    data_1 := []User{
     {Name: "Afif",
     Age: 16,},
     {Name: "Joko",
     Age: 16,},
     {Name: "Eko",
     Age: 16,},
     {Name: "Bambang",
     Age: 16,},
    }

	  db.Create(&data_1)

	db.AutoMigrate(&Student{})

	data_student := []Student{}

	fmt.Println("Cetak data student umur atas 10")
	
	db.Raw("SELECT name, age, class FROM students WHERE age > ?", 10).Scan(data_student)
	for _, el := range data_student {
		fmt.Println("Name:", el.Name)
		fmt.Println("Umur:", el.Age)
		fmt.Println("Kelas:", el.Class)
	}
	
	fmt.Println("Cetak data student di class tertentu")
	
	db.Where("class = ?", "RPL").Find(&data_student)
	for _,el := range data_student {
		fmt.Println("Name:", el.Name)
		fmt.Println("Umur:", el.Age)
		fmt.Println("Kelas:", el.Class)
		
	}

}