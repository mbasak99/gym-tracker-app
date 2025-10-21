package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello World!"))
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

	fmt.Printf("Running on http://localhost%s\n", port)
	err = http.ListenAndServe(port, mux)
	log.Fatalln(err)
}
