package handlers

import "net/http"

func RegisterLoginRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /signup", signUp)
	router.HandleFunc("POST /signin", signIn)
}

func signUp(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello from Sign Up")) }

func signIn(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Hello from Sign In")) }
