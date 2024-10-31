package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB
var err error

// const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
const DNS = "host=localhost user=postgres password=1234 dbname=godb port=5432 sslmode=disable TimeZone=UTC"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content_type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(writer).Encode(users)

}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content_type", "application/json")
	params := mux.Vars(request)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(writer).Encode(user)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "aplication/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(writer).Encode(user)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content_type", "application/json")
	params := mux.Vars(request)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(request.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(writer).Encode(user)

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content_type", "application/json")
	params := mux.Vars(request)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(writer).Encode("user is deleted successfully")
}
