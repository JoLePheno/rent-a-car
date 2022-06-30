package car

import (
	"fmt"
	"project/internal/model"
	"project/internal/port"
)

type CarController struct {
	Store port.RentalStore
}

func (u *CarController) GetCarsController() ([]*model.Car, error) {
	cars, err := u.Store.GetCars()
	if err != nil {
		fmt.Println("error during fetching cars, err: %w", err)
		return nil, err
	}

	return cars, nil
}

func (u *CarController) RentCarController(car *model.Car) error {
	err := u.Store.RentCar(car)
	if err != nil {
		fmt.Println("error during free car, err: %w", err)
		return err
	}

	return nil
}

func (u *CarController) FreeCarController(car *model.Car) error {
	err := u.Store.FreeCar(car)
	if err != nil {
		fmt.Println("error during free car, err: %w", err)
		return err
	}

	return nil
}
