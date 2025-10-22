package main

import (
	"fmt"
	"gym-tracker-backend/src/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env file")
	}
	port := os.Getenv("PORT")

	// Load handlers
	mux.HandleFunc("/", home)
	handlers.RegisterLoginRoutes(mux)

	fmt.Printf("Running on http://localhost%s\n", port)
	err = http.ListenAndServe(port, mux)
	log.Fatalln(err)
}
