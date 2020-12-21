package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erhansakarya/gorm_sqlite/controller"
	"github.com/erhansakarya/gorm_sqlite/migration"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", controller.Ping).Methods("GET")
	myRouter.HandleFunc("/users", controller.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", controller.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", controller.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", controller.UpdateUser).Methods("PUT")
	http.Handle("/", myRouter)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM")

	migration.InitialMigration()

	handleRequests()
}
