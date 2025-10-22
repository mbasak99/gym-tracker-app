package main

import (
	"fmt"
	"gym-tracker-backend/src/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Home(c echo.Context) error {
	// fmt.Fprintln(w, "Hello World!")
	return c.String(http.StatusOK, "Hello World!")
}

func main() {
	// mux := http.NewServeMux()
	mux := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Failed to load .env file")
	}
	port := os.Getenv("PORT")

	// Load middlewares
	mux.Use(middleware.Logger())
	mux.Use(middleware.Recover())

	// Load handlers
	mux.GET("/", Home)
	handlers.RegisterLoginRoutes(mux) // TODO: update and add the login/signup routes to new echo routes

	fmt.Printf("Running on http://localhost%s\n", port)
	// err = http.ListenAndServe(port, mux)
	err = mux.Start(port)
	log.Fatalln(err)
}
