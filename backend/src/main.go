package main

import (
	"context"
	"fmt"
	"gym-tracker-backend/src/handlers"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ENV struct {
	db *pgx.Conn
}

func Home(c echo.Context) error {
	// fmt.Fprintln(w, "Hello World!")
	return c.String(http.StatusOK, "Hello World!")
}

func main() {
	// mux := http.NewServeMux()
	mux := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
	port := os.Getenv("PORT")

	// TODO: Load DB conn here
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panicf("Failed to connect to database. %+v", err)
	}
	defer conn.Close(ctx)

	// Test DB connection and version
	var version string
	err = conn.QueryRow(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		log.Panicf("Couldn't retrieve DB version. %+v", err)
	}
	log.Printf("Successfully connected to OrioleDB version: %s\n", version)

	// Load middlewares
	mux.Use(middleware.Logger())
	mux.Use(middleware.Recover())

	// Load handlers
	mux.GET("/", Home)
	handlers.RegisterLoginRoutes(mux, conn) // TODO: update and add the login/signup routes to new echo routes

	fmt.Printf("Running on http://localhost%s\n", port)
	// err = http.ListenAndServe(port, mux)
	err = mux.Start(port)
	log.Fatalln(err)
}
