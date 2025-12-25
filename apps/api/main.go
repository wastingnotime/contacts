package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type contact struct {
	ID          string `json:"id" bson:"id"`
	FirstName   string `json:"firstName" bson:"firstName" validate:"required"`
	LastName    string `json:"lastName" bson:"lastName" validate:"required"`
	PhoneNumber string `json:"phoneNumber" bson:"phoneNumber" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// environment
	godotenv.Load()
	environment := os.Getenv("ENVIRONMENT")
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")
	mongoCollection := "contacts"

	// api
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	if environment == "development" {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())
	e.HideBanner = true

	// database
	if mongoURI == "" || mongoDB == "" || mongoCollection == "" {
		log.Panic("MONGO_URI, MONGO_DB, and MONGO_COLLECTION must be set")
	}

	mongoClient, err := NewMongoClient(mongoURI)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("mongo disconnect error: %v", err)
		}
	}()

	// define handlers
	h := &handler{repo: NewMongoContactRepository(mongoClient, mongoDB, mongoCollection)}

	// routes
	e.POST("/contacts", h.CreateContact)
	e.GET("/contacts", h.GetContacts)
	e.GET("/contacts/:id", h.GetContact)
	e.PUT("/contacts/:id", h.UpdateContact)
	e.DELETE("/contacts/:id", h.DeleteContact)

	// health/readiness probes for swarm
	e.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	e.GET("/readyz", func(c echo.Context) error { return c.NoContent(http.StatusOK) })

	// run
	e.Logger.Fatal(e.Start(":8010"))
}
