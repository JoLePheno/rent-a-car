package postgres

import (
	"fmt"
	"project/internal/model"
	"time"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
)

type Users struct {
	ID int64 `sql:"id"`

	CreatedAt time.Time `sql:"created_at,notnull"`
	UpdatedAt time.Time `sql:"updated_at,notnull"`
	DeletedAt time.Time `pg:"deleted_at,soft_delete"`

	UserUUID    uuid.UUID `sql:"user_id,notnull"`
	Name        string    `sql:"name"`
	Email       string    `sql:"email,notnull"`
	PhoneNumber string    `sql:"phone_number"`
}

func (s *RentalStore) CreateUser(user *model.User) error {
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		resp, _ := s.GetUserByMail(user.Email)
		if resp != nil && resp.Email != "" {
			return fmt.Errorf("error email already used")
		}
		currentTime := time.Now()
		return tx.Insert(&Users{
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			UserUUID:    uuid.New(),
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		})
	},
	)
	if err != nil {
		fmt.Println("error when creating user, err %w", err)
		return err
	}
	fmt.Println("user created")
	return nil
}

func (s *RentalStore) DeleteUser(user *model.User) error {
	var u Users
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		err := tx.Model(&u).Where("email = ?", user.Email).Select()
		if err != nil {
			fmt.Println("error when fetching user, err: ", err)
			return err
		}
		fmt.Println("voici user_id: ", u.UserUUID)
		return s.db.Delete(&u)
	},
	)
	if err != nil {
		fmt.Println("error when deleting user, err %w", err)
		return err
	}
	fmt.Println("user deleted")
	return nil
}

func (s *RentalStore) GetUserByMail(userEmail string) (*model.User, error) {
	var userModel Users
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		return tx.Model(&userModel).Where("email = ?", userEmail).Select()
	},
	)
	if err != nil {
		fmt.Println("error when fetching user, err %w", err)
		return nil, err
	}
	return convertUserModelDBOToModel(userModel), nil
}

func (s *RentalStore) UpdateUser(user *model.User) (*model.User, error) {
	var userModel Users
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		userModel.Name = user.Name
		userModel.PhoneNumber = user.PhoneNumber
		userModel.Email = user.Email

		_, err := tx.Model(&userModel).Where("email = ? AND phone_number = ?", user.Email, user.PhoneNumber).Update()
		return err
	},
	)
	if err != nil {
		fmt.Println("error when updating user, err %w", err)
		return nil, err
	}
	fmt.Println("user updated")

	return &model.User{
		ID:          userModel.UserUUID,
		Name:        userModel.Name,
		Email:       userModel.Email,
		PhoneNumber: userModel.PhoneNumber,
		CreatedAt:   userModel.CreatedAt,
		UpdateAt:    userModel.UpdatedAt,
	}, nil
}

func convertUserModelDBOToModel(userModelDBO Users) *model.User {
	return &model.User{
		ID:          userModelDBO.UserUUID,
		Name:        userModelDBO.Name,
		Email:       userModelDBO.Email,
		PhoneNumber: userModelDBO.PhoneNumber,
		CreatedAt:   userModelDBO.CreatedAt,
	}
}
