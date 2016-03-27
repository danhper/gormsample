package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// import whatever DB
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/tuvistavie/gormsample/managers"
	"github.com/tuvistavie/gormsample/models"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	managers.Boot(db)

	myUser := models.User{Name: "abc"}
	err = managers.UserManager.CreateUser(myUser)
	fmt.Println(err)
}
