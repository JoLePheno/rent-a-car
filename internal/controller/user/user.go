package user

import (
	"fmt"
	"project/internal/model"
	"project/internal/port"
)

type UserController struct {
	Store port.RentalStore
}

func (u *UserController) CreateUserController(user *model.User) error {
	err := u.Store.CreateUser(user)
	if err != nil {
		fmt.Println("error during create user, err: %w", err)
		return err
	}

	return nil
}

func (u *UserController) DeleteUserController(user *model.User) error {
	err := u.Store.DeleteUser(user)
	if err != nil {
		fmt.Println("error during deleting user, err: %w", err)
		return err
	}

	return nil
}

func (u *UserController) GetUserController(userEmail string) (*model.User, error) {	
	user, err := u.Store.GetUserByMail(userEmail)
	if err != nil {
		fmt.Println("error during fetching user, err: %w", err)
		return nil, err
	}

	return user, nil
}
