package service

import (
	"net/http"
	"project/internal/controller/car"
	"project/internal/controller/user"
	"project/internal/utils"

	"github.com/gorilla/mux"
)

type RentalService struct {
	Car  *car.CarController
	User *user.UserController
}

func (s *RentalService) Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.Methods("POST").Name("CreateUser").Handler(s.CreateUserHandler()).Path("/user")
	r.Methods("GET").Name("GetUser").Handler(s.GetUserHandler()).Path("/user")
	r.Methods("DELETE").Name("DeleteUser").Handler(s.DeleteUserHandler()).Path("/user")

	r.Methods("GET").Name("GetCars").Handler(s.GetCarsHandler()).Path("/cars")
	r.Methods("POST").Name("RentCar").Handler(s.RentCarHandler()).Path("/car/rent")
	r.Methods("POST").Name("FreeCar").Handler(s.FreeCarHandler()).Path("/car/free")

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	message := utils.Message(true, "Stoïk looks cool")
	utils.Respond(w, message, 200)
}
