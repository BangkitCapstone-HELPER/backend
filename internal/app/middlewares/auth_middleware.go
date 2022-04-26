package middlewares

import (
	"fmt"
	"net/http"

	"github.com/BangkitCapstone-HELPER/backend/internal/app/constants"
	errors "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type authMiddlewareParams struct {
	fx.In

	JWT lib.JWT
}

type AuthMiddleware interface {
	Setup(permission constants.Permission) echo.MiddlewareFunc
}

func NewAuthMiddleware(p authMiddlewareParams) AuthMiddleware {
	return &p
}

func (m authMiddlewareParams) Setup(permission constants.Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := utils.ExtractToken(c)
			if err != nil {
				return err
			}

			user, err := utils.GetUserFromToken(token, m.JWT)
			fmt.Println(permission)

			if err != nil {
				return lib.Response{
					Message: err.Error(),
					Status:  http.StatusUnauthorized,
				}.JSON(c)
			}

			if permission == constants.PermissionAll {
				return next(c)
			} else if permission == constants.PermissionAdmin && user.IsAdmin {
				return next(c)
			} else if permission == constants.PermissionNonAdmin && !user.IsAdmin {
				return next(c)
			}

			return errors.ErrRoleUnauthorized

		}
	}

}
