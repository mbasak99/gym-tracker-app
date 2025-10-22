package handlers

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var hashedPassword []byte

func RegisterLoginRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /signup", signUp)
	router.HandleFunc("POST /login", logIn)
}

func errorHandler(w http.ResponseWriter, output string, err error) {
	if err != nil {
		fmt.Fprint(w, output, err)
	} else {
		fmt.Fprint(w, output)
	}
}

func logIn(w http.ResponseWriter, r *http.Request) {
	// grab email and password from req
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, "Failed to parse request body. %+v\n", err)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	// TODO: check email exist in db and convert hashed password in table and compare
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) // change hashedPassword to hash from db
	if err != nil {
		errorHandler(w, "Invalid credentials. Please try again.\n", nil)
	}

	fmt.Fprintf(w, "Email: %s, Password: %s, Same?: %t", email, password, err == nil)
}

func signUp(w http.ResponseWriter, r *http.Request) {
	// grab email and password from req
	err := r.ParseForm()
	if err != nil {
		errorHandler(w, "Failed to parse request body. %+v\n", err)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	// TODO: check if user already exists

	// hash the password and store it to db
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errorHandler(w, "Failed to hash password. %+v\n", err)
		return
	}
	// hashedPassword = hashedPass

	// w.Write([]byte("Hello from Log In"))
	fmt.Fprintf(w, "Email: %s, Password: %s, HashedPW: %s", email, password, string(hashedPass))
}
