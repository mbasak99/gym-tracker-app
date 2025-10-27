package handlers

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	hashedPassword []byte
	db             *pgx.Conn
)

func RegisterLoginRoutes(router *echo.Echo, dbConn *pgx.Conn) {
	// router.HandleFunc("POST /signup", signUp)
	// router.HandleFunc("POST /login", logIn)
	db = dbConn
	router.POST("/signup", signUp)
	router.POST("/login", logIn)
}

func logIn(c echo.Context) error {
	// grab email and password from req
	err := c.Request().ParseForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read parameters.")
	}
	email := c.FormValue("email")
	password := c.FormValue("password")

	// TODO: check email exist in db and convert hashed password in table and compare
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) // change hashedPassword to hash from db
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials.")
	}

	fmt.Printf("Email: %s, Password: %s, Same?: %t", email, password, err == nil)
	return c.String(http.StatusOK, "Successfully logged in user.")
}

func signUp(c echo.Context) error {
	// grab email and password from req
	err := c.Request().ParseForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read parameters.")
	}
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		// check values are valid
		fmt.Printf("Email: %s | Password: %s\n", email, password)
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials.")
	}

	// TODO: check if user already exists

	// hash the password and store it to db
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password.")
	}
	hashedPassword = hashedPass // TODO: remove this

	// w.Write([]byte("Hello from Log In"))
	fmt.Printf("Email: %s, Password: %s, HashedPW: %s", email, password, string(hashedPass))
	return c.String(http.StatusCreated, "Successfully signed up user.")
}
