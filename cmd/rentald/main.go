package main

import (
	"context"
	"log"
	"net/http"
	"project/internal/adapter/postgres"
	"project/internal/controller/car"
	"project/internal/controller/user"
	"project/internal/service"

	"github.com/go-pg/pg"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	store := postgres.NewRentalStore(pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "stoik",
	}))
	rentalService := service.RentalService{
		User: &user.UserController{
			Store: store,
		},
		Car: &car.CarController{
			Store: store,
		},
	}

	log.Fatal(http.ListenAndServe(":3000", rentalService.Router()))
}
