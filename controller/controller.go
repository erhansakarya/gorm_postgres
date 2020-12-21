package controller

import (
	"fmt"
	"net/http"
)

// Ping returns Pong
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

// AllUsers gets all users
func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all users endpoint")
}

// NewUser creates a new user
func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "new user endpoint")
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete user endpoint")
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update user endpoint")
}
