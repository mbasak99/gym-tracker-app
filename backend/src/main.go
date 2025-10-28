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
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), // just "host.docker.internal"
		os.Getenv("DB_PORT"), // "5432"
		os.Getenv("DB_USER"), // "monark"
		os.Getenv("DB_PASS"), // your password
		os.Getenv("DB_NAME"))
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Panicf("Failed to connect to database. %+v", err)
	}
	defer conn.Close(ctx)

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
