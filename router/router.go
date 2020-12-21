package router

import (
	"log"
	"net/http"

	"github.com/erhansakarya/gorm_postgres/controller"
	"github.com/gorilla/mux"
)

// HandleRequests routes handlers
func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", controller.Ping).Methods("GET")
	myRouter.HandleFunc("/users", controller.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", controller.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", controller.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", controller.UpdateUser).Methods("PUT")
	http.Handle("/", myRouter)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
