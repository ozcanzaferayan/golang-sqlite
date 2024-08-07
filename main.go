package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err:= gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!= nil {
		log.Fatal("Failed to connect")
	}

	log.Println("Sqlite Opened")

	db.AutoMigrate(&User{})

	user := User{Name: "Zafer", Email: "ozcanzaferayan@gmail.com"}
	db.Create(&user);

	var readUser User
	db.First(&readUser, user.ID)
	log.Println("User: ", readUser.Name, readUser.Email)

	db.Model(&user).Update("Email", "zafer@example.com")

	// db.Delete(user)

}