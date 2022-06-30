package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"project/internal/model"
	"project/internal/utils"
	"time"

	"github.com/google/uuid"
)

type rentCarResp struct {
	ID           uuid.UUID `json:"car_id"`
	Model        string    `json:"model"`
	Brand        string    `json:"brand"`
	Availability bool      `json:"availibility"`
	RentAt       time.Time `json:"rent_at"`
}

type rentCarReq struct {
	ID       uuid.UUID `json:"car_id"`
	RenterID uuid.UUID `json:"renter_id"`
}

func (s *RentalService) GetCarsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request /cars, fetching cars")

		cars, err := s.Car.GetCarsController()
		if err != nil {
			fmt.Println("an error append in creatUserController, err : %w", err)
			utils.Respond(w, utils.Message(false, err.Error()), 400)
		}

		carsResp := make([]*rentCarResp, len(cars))
		for i := range cars {
			carsResp[i] = convertCarModelToRespModel(cars[i])
		}
		utils.Respond(w, map[string]interface{}{
			"cars": carsResp,
		}, 200)
	})
}

func convertCarModelToRespModel(car *model.Car) *rentCarResp {
	return &rentCarResp{
		ID:           car.ID,
		Model:        car.Model,
		Brand:        car.Brand,
		Availability: car.Availability,
		RentAt:       car.RentAt,
	}
}

func (s *RentalService) RentCarHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request /car/rent, renting car")

		car := &rentCarReq{}

		err := json.NewDecoder(r.Body).Decode(car) //decode the request body into struct, failed if any error occured
		if err != nil {
			fmt.Println(errors.New("An error occurred while decoding request, err: " + err.Error()))
			utils.Respond(w, utils.Message(false, "Invalid request"), 400)
			return
		}

		err = s.Car.RentCarController(&model.Car{
			ID:         car.ID,
			RenterUUID: car.RenterID,
		})
		if err != nil {
			fmt.Println("an error append in creatUserController, err : %w", err)
			utils.Respond(w, utils.Message(false, err.Error()), 400)
		}

		resp := make(map[string]interface{})
		utils.Respond(w, resp, 200)
	})
}

func (s *RentalService) FreeCarHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request /car/free, renting car")

		car := &rentCarReq{}

		err := json.NewDecoder(r.Body).Decode(car) //decode the request body into struct, failed if any error occured
		if err != nil {
			fmt.Println(errors.New("An error occurred while decoding request, err: " + err.Error()))
			utils.Respond(w, utils.Message(false, "Invalid request"), 400)
			return
		}

		err = s.Car.FreeCarController(&model.Car{
			ID:         car.ID,
			RenterUUID: car.RenterID,
		})
		if err != nil {
			fmt.Println("an error append in freeCarController, err : %w", err)
			utils.Respond(w, utils.Message(false, err.Error()), 400)
		}

		resp := make(map[string]interface{})
		utils.Respond(w, resp, 200)
	})
}