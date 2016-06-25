package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Define a basic Model
type (
	User struct {
		gorm.Model

		Email string
		Passwd string
	}
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=daemongear dbname=daemongear sslmode=disable password=daemongear")

	if err != nil {
		fmt.Println(err.Error())
		panic("failed connect to database")
	}

	db.LogMode(true)
	db.AutoMigrate(&User{})

	fmt.Println("Clean Database...")
	db.Delete(User{})

	fmt.Println("Creating Users")
	db.Create(&User{Email: "jorgee.araneda@gmail.com", Passwd: "1234"})
	db.Create(&User{Email: "gonzalo.bahamondez.c@gmail.com", Passwd: "5678"})
	db.Create(&User{Email: "robertoosorio.s@gmail.com", Passwd: "19283"})

	var user User 
	fmt.Println("Find User by Email")
	db.First(&user, "email=  ?", "robertoosorio.s@gmail.com")

	if &user == nil {
		panic("User not found, can't continue")
	}

	fmt.Printf("User Found : %v\n", user)

	var users []User
	db.Find(&users)

	if &users == nil {
		panic("Users not found, can't continue")	
	}

	fmt.Printf("User Found : %v\n", users)
}