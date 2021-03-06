package dto

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
)

type (
	CreateUserRequestDTO struct {
		Name        string `json:"name" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		Password    string `json:"password" validate:"required,gte=8"`
		Image       string `json:"image"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
	}

	UserDTO struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		ID          uint64 `json:"id"`
		Image       string `json:"image"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		IsAdmin     bool   `json:"isAdmin"`
	}

	UpdateUserDTO struct {
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		IsAdmin     bool   `json:"is_admin"`
		Image       string `json:"image"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8"`
	}
	ChangePasswordRequest struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password" validate:"required,gte=8"`
	}
	ForgetPasswordRequest struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password" validate:"required,gte=8"`
	}

	LoginResponse struct {
		User  UserDTO `json:"user"`
		Token string  `json:"token"`
	}
)

func (c CreateUserRequestDTO) ToDAO() dao.User {
	return dao.User{
		Name:        c.Name,
		Email:       c.Email,
		Password:    c.Password,
		Image:       c.Image,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
	}
}

func NewUserDTO(user dao.User) UserDTO {
	return UserDTO{
		Name:        user.Name,
		Email:       user.Email,
		ID:          uint64(user.ID),
		Image:       user.Image,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		IsAdmin:     user.IsAdmin,
	}
}
