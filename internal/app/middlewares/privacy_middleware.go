package middlewares

import (
	"net/http"
	"strconv"

	errors "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type privacyMiddlewareParams struct {
	fx.In

	JWT lib.JWT
}

type PrivacyMiddleware interface {
	Setup() echo.MiddlewareFunc
}

func NewPrivacyMiddleware(p privacyMiddlewareParams) PrivacyMiddleware {
	return &p
}

func (m privacyMiddlewareParams) Setup() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := utils.ExtractToken(c)
			if err != nil {
				return err
			}

			user, err := utils.GetUserFromToken(token, m.JWT)

			if err != nil {
				return lib.Response{
					Message: err.Error(),
					Status:  http.StatusUnauthorized,
				}.JSON(c)
			}

			idInString := c.Param("id")
			id, _ := strconv.ParseUint(idInString, 10, 32)
			if user.ID == id {
				return next(c)
			}

			return errors.ErrRoleUnauthorized

		}
	}

}
