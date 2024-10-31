package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var listStr = []string{"John", "Doe", "Smith"}

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", r))

}

func main() {
	InitialMigration()
	initializeRouter()
}

func main1() {
	fmt.Println("Starting server...")
	log.Println("logging here")
	//var name = "John"
	//age := 50
	//fmt.Println(name)
	//fmt.Println(age)

	//fmt.Println(listStr)
	//fmt.Println("Hello world!")
	//step := 5
	//var val1 int = 2
	//var val2 int = 23
	//fmt.Println(val1 == val2)
	//testPackage.LogData(step)
	//fmt.Println("list : ", listStr)

	//for _, val := range listStr {
	//	fmt.Println("name :", val)
	//}

	//fmt.Println(addItem(listStr, "Test"))
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/test", testUsers)
	http.HandleFunc("/post", postUsers)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		//return
		log.Fatal("error: error loading api")
	}

	//var newError error = returnsError("2345")
	//if newError != nil {
	//	fmt.Println(newError)
	//}
}

func postUsers(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	errResponse := json.NewDecoder(request.Body).Decode(&req)

	if errResponse != nil {
		log.Println(errResponse)
		//http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	fmt.Println(req)
	fmt.Println(request.Body)

}

func returnsError(password string) error {
	fmt.Println("====================_____+++++++++++++++++")
	var secretPassword string = "1234"
	if password == secretPassword {
		return nil
	} else {
		return errors.New("Invalid password")
	}
}

func testUsers(writer http.ResponseWriter, request *http.Request) {
	for _, item := range listStr {
		fmt.Fprintln(writer, item)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greetings = "hello User"
	var greetingsByte []byte = []byte(greetings)
	_, err := writer.Write(greetingsByte)
	if err != nil {
		fmt.Fprintln(writer, "Error here o")
	}
}

func addItem(list []string, item string) []string {
	var updatedList []string = append(list, item)
	return updatedList
}
