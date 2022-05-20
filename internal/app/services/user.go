package services

import (
	"encoding/json"
	errors "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/repo"
	"go.uber.org/fx"
)

type userServiceParams struct {
	fx.In

	UserRepo repo.UserRepo
	JWT      lib.JWT
	Hash     lib.Hash
}

type UserService interface {
	// GetUser(id uint64) (dao.User, error)
	// GetUser(token string) (dao.User, error)
	GetUserById(id uint64) (dao.User, error)
	GetUserByEmail(email string) (dao.User, error)
	CreateUser(user dto.CreateUserRequestDTO) (dto.UserDTO, error)
	UpdateUser(userId uint, userDTO dto.UpdateUserDTO) (dto.UserDTO, error)
	Login(loginRequest dto.LoginRequest) (dto.LoginResponse, error)
}

func NewUserService(params userServiceParams) UserService {
	return &params
}

func (u *userServiceParams) GetUserById(id uint64) (dao.User, error) {
	return u.UserRepo.GetUser(id)
}

// func (u *userServiceParams) GetUser(token string) (dao.User, error) {
// 	user, err := utils.GetUserFromToken(token, u.JWT)
// 	return u.UserRepo.GetUser(id)
// }

func (u *userServiceParams) GetUserByEmail(email string) (dao.User, error) {
	return u.UserRepo.GetUserByEmail(email)
}

func (u *userServiceParams) CreateUser(user dto.CreateUserRequestDTO) (dto.UserDTO, error) {

	if _, err := u.UserRepo.GetUserByEmail(user.Email); err == nil {
		return dto.UserDTO{}, errors.ErrEmailAlreadyExist
	}

	if hashed, hasherr := u.Hash.Hash(user.Password); hasherr != nil {
		return dto.UserDTO{}, hasherr
	} else {
		user.Password = hashed
	}
	newUser, err := u.UserRepo.CreateUser(user.ToDAO())
	return dto.NewUserDTO(newUser), err
}

func (u *userServiceParams) UpdateUser(userId uint, userDTO dto.UpdateUserDTO) (dto.UserDTO, error) {
	var updateMap map[string]interface{}
	data, _ := json.Marshal(userDTO)
	json.Unmarshal(data, &updateMap)
	for k, v := range updateMap {
		if v == nil {
			delete(updateMap, k)
		}
	}

	user, err := u.UserRepo.UpdateUser(uint64(userId), updateMap)
	return dto.NewUserDTO(user), err
}

func (u *userServiceParams) Login(loginRequest dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := u.UserRepo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	passwordValid := u.Hash.Compare(loginRequest.Password, user.Password)
	if !passwordValid {
		return dto.LoginResponse{}, errors.ErrWrongPassword
	}

	data, _ := json.Marshal(dto.NewUserDTO(user))
	var claims map[string]interface{}
	json.Unmarshal(data, &claims)
	token, err := u.JWT.Encode(claims)

	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		User: dto.UserDTO{
			Name:        user.Name,
			Email:       user.Email,
			ID:          uint64(user.ID),
			Address:     user.Address,
			PhoneNumber: user.PhoneNumber,
			IsAdmin:     user.IsAdmin,
		},
		Token: token,
	}, nil
}
