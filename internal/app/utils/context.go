package utils

import (
	"encoding/json"
	"strings"

	errors "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/labstack/echo/v4"
)

func ExtractToken(c echo.Context) (string, error) {
	authorizationHeader := c.Request().Header.Get("Authorization")

	if authorizationHeader == "" {
		return "", errors.ErrMissingAuthorization
	}

	if !strings.HasPrefix(strings.ToLower(authorizationHeader), "bearer") {
		return "", errors.ErrWrongAuthorization
	}

	splitted := strings.Split(authorizationHeader, " ")

	if len(splitted) != 2 {
		return "", errors.ErrWrongAuthorization
	}

	token := splitted[1]

	return token, nil
}

func GetUserFromToken(token string, jwt lib.JWT) (dto.UserDTO, error) {
	userMap, err := jwt.Decode(token)
	if err != nil {
		return dto.UserDTO{}, err
	}

	jsonString, _ := json.Marshal(userMap)
	var user dto.UserDTO
	json.Unmarshal(jsonString, &user)
	if user.ID == 0 {
		return dto.UserDTO{}, errors.ErrDecodingJWT
	}

	return user, nil
}
