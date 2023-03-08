package main

import (
	"fmt"
	"gomysql/models"
	"gomysql/db"
)

func main(){
	db.Connect()

	//fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	user := models.CreateUser("Jonathan", "papá", "jonathanñ@gmail.com")
	fmt.Println(user)
	//users := models.ListUsers()
	//fmt.Println(users)
	// user := models.GetUser(1)
	// fmt.Println(user)

	// user.Delete()

	//db.TruncateTable("users")
	fmt.Println(models.ListUsers())

	//db.Ping()
	//db.TruncateTable("users")
	db.Close()
}