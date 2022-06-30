package port

import "project/internal/model"

type RentalStore interface {
	CreateUser(user *model.User) error
	DeleteUser(user *model.User) error
	GetUserByMail(userMail string) (*model.User, error)

	GetCars() ([]*model.Car, error)
	RentCar(car *model.Car) error
	FreeCar(car *model.Car) error
}
