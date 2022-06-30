package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"project/internal/model"
	"project/internal/utils"

	"github.com/google/uuid"
)

type userReq struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type userResp struct {
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	ID          uuid.UUID `json:"user_id"`
}

func (s *RentalService) CreateUserHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request /user, creating new user")

		user := &userReq{}

		err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct, failed if any error occured
		if err != nil {
			fmt.Println(errors.New("An error occurred while decoding request, err: " + err.Error()))
			utils.Respond(w, utils.Message(false, "Invalid request"), 400)
			return
		}

		userModel := &model.User{
			ID:          uuid.Nil,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}

		err = s.User.CreateUserController(userModel)
		if err != nil {
			fmt.Println("an error append in creatUserController, err : %w", err)
			utils.Respond(w, utils.Message(false, err.Error()), 400)
		}

		resp := make(map[string]interface{})
		utils.Respond(w, resp, 201)
	})
}

func (s *RentalService) DeleteUserHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request /user, deleting user")

		user := &userReq{}

		err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct, failed if any error occured
		if err != nil {
			fmt.Println(errors.New("An error occurred while decoding request, err: " + err.Error()))
			utils.Respond(w, utils.Message(false, "Invalid request"), 400)
			return
		}

		userModel := &model.User{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}

		err = s.User.DeleteUserController(userModel)
		if err != nil {
			fmt.Println("an error append in deleteUserController, err : %w", err)
			utils.Respond(w, utils.Message(false, err.Error()), 400)
		}

		resp := make(map[string]interface{})
		utils.Respond(w, resp, 200)
	})
}

func (s *RentalService) GetUserHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request /user, get user")

		user := &userReq{}

		err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct, failed if any error occured
		if err != nil {
			fmt.Println(errors.New("An error occurred while decoding request, err: " + err.Error()))
			utils.Respond(w, utils.Message(false, "Invalid request"), 400)
			return
		}

		userModel := &model.User{
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}

		u, err := s.User.GetUserController(userModel.Email)
		if err != nil {
			fmt.Println("an error append in GetUserController, err : %w", err)
			utils.Respond(w, utils.Message(false, err.Error()), 400)
		}

		resp := userResp{
			Email:       u.Email,
			Name:        u.Name,
			PhoneNumber: u.PhoneNumber,
			ID:          u.ID,
		}
		utils.Respond(w, map[string]interface{}{
			"user": resp,
		}, 200)
	})
}
