package postgres

import (
	"fmt"
	"project/internal/model"
	"time"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Cars struct {
	ID int64 `sql:"id"`

	RentAt time.Time `sql:"rent_at"`

	CarUUID    uuid.UUID `sql:"car_id,notnull"`
	RenterUUID uuid.UUID `sql:"renter_id"`

	Model string `sql:"model,notnull"`
	Brand string `sql:"brand,notnull"`

	Availability bool `sql:"availability"`
}

func (s *RentalStore) GetCars() ([]*model.Car, error) {
	var cars []Cars
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		return tx.Model(&cars).Select()
	},
	)
	if err != nil {
		fmt.Println("error when fetching cars, err %w", err)
		return nil, err
	}

	carsModel := make([]*model.Car, len(cars))
	for i := range cars {
		carsModel[i] = convertCarModelDBOToModel(cars[i])
	}
	return carsModel, nil
}

func convertCarModelDBOToModel(car Cars) *model.Car {
	return &model.Car{
		ID:           car.CarUUID,
		RenterUUID:   car.RenterUUID,
		Model:        car.Model,
		Brand:        car.Brand,
		Availability: car.Availability,
		RentAt:       car.RentAt,
	}
}

func (s *RentalStore) RentCar(car *model.Car) error {
	var cars Cars
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		err := tx.Model(&cars).Where("car_id = ?", car.ID).Select()
		if err != nil {
			return err
		}
		if !cars.Availability {
			return fmt.Errorf("car already rented")
		}

		_, err = tx.Model(&cars).
			Where("car_id = ?", car.ID).
			Set("renter_id = ?", car.RenterUUID).
			Set("availability = ?", false).
			Set("rent_at = ?", time.Now()).
			Update()
		if err != nil {
			return err
		}
		return nil
	},
	)
	if err != nil {
		fmt.Println("error when renting car, err %w", err)
		return err
	}

	return nil
}

func (s *RentalStore) FreeCar(car *model.Car) error {
	var cars Cars
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		err := tx.Model(&cars).Where("car_id = ?", car.ID).Select()
		if err != nil {
			return err
		}
		if cars.Availability {
			return fmt.Errorf("car already free")
		}

		_, err = tx.Model(&cars).
			Where("car_id = ?", car.ID).
			Set("renter_id = ?", nil).
			Set("availability = ?", true).
			Set("rent_at = ?", nil).
			Update()
		if err != nil {
			return err
		}
		return nil
	},
	)
	if err != nil {
		fmt.Println("error when free car, err %w", err)
		return err
	}

	return nil
}
