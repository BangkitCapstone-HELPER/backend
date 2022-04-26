package lib

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"

	e "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HTTPServer interface {
	Engine() *echo.Echo
	RouterV1() *echo.Group
}

type HTTPHandler struct {
	engine   *echo.Echo
	routerV1 *echo.Group
}

type Response struct {
	Status  int         `json:"status"`
	Pretty  bool        `json:"-"`
	Data    interface{} `json:"data,omitempty"`
	Message interface{} `json:"message"`
}

type Validator interface {
	Validate(i interface{}) error
}

// ValidatorImpl is a custom validator
type ValidatorImpl struct {
	validate *validator.Validate
}

// Validate validates the request body
func (a *ValidatorImpl) Validate(i interface{}) error {
	return a.validate.Struct(i)
}

func newValidator(v *validator.Validate) Validator {
	return &ValidatorImpl{
		validate: v,
	}
}

type BinderWithValidation struct{}

func (BinderWithValidation) Bind(i interface{}, ctx echo.Context) error {
	binder := &echo.DefaultBinder{}
	var response Response = Response{
		Status: http.StatusUnprocessableEntity,
	}
	if err := binder.Bind(i, ctx); err != nil {
		response.Message = err.(*echo.HTTPError).Message.(string)
		return response.JSON(ctx)
	}

	if err := ctx.Validate(i); err != nil {
		// Validate only provides verification function for struct.
		// When the requested data type is not struct,
		// the variable should be considered legal after the bind succeeds.
		if reflect.ValueOf(i).Elem().Kind() != reflect.Struct {
			return nil
		}
		if ferrs, ok := err.(validator.ValidationErrors); ok {
			var buf bytes.Buffer

			for _, ferr := range ferrs {
				buf.WriteString("Validation failed on ")
				buf.WriteString(ferr.Tag())
				buf.WriteString(" for ")
				buf.WriteString(ferr.StructField())
				buf.WriteString("\n")
			}
			response.Message = buf.String()
			return response.JSON(ctx)
		}
		response.Message = err.Error()
		return response.JSON(ctx)
	}

	return nil
}

func (a Response) JSON(ctx echo.Context) error {
	if a.Message == "" || a.Message == nil {
		a.Message = http.StatusText(a.Status)
	}

	if err, ok := a.Message.(error); ok {
		a.Message = err.Error()
	}

	if a.Pretty {
		return ctx.JSONPretty(a.Status, a, "\t")
	}

	return ctx.JSON(a.Status, a)
}

func errorResponse(err error, errorMappers e.ErrorMappers) Response {
	var (
		code    = http.StatusInternalServerError
		message interface{}
	)

	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		message = he.Message

		if he.Internal != nil {
			message = fmt.Errorf("%v - %v", message, he.Internal)
		}
	}
	for _, e := range errorMappers {
		if status, ok := e[err]; ok {
			code = status
			message = err.Error()
			return Response{
				Status:  code,
				Message: message,
			}
		}
	}
	return Response{
		Status:  code,
		Message: err.Error(),
	}
}

func (s *HTTPHandler) RouterV1() *echo.Group {
	return s.routerV1
}

func (s *HTTPHandler) Engine() *echo.Echo {
	return s.engine
}

func HTTPErrorHandler(errorMapper e.ErrorMappers) func(err error, ctx echo.Context) {
	return func(err error, ctx echo.Context) {
		errRes := errorResponse(err, errorMapper)

		// Send response
		if !ctx.Response().Committed {

			// https://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html
			if ctx.Request().Method == http.MethodHead {
				err = ctx.NoContent(errRes.Status)
			} else {
				err = errRes.JSON(ctx)
			}

			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

// newEchoValidator returns a new validator
func newEchoValidator() echo.Validator {
	v := validator.New()

	return newValidator(v)
}

// notFoundHandler is the http handler for 404 http error
func notFoundHandler(ctx echo.Context) error {
	return Response{Status: http.StatusNotFound}.JSON(ctx)
}

// methodNotAllowedHandler is the http handler for 405 http error
func methodNotAllowedHandler(ctx echo.Context) error {
	return Response{Status: http.StatusMethodNotAllowed}.JSON(ctx)
}

func NewHTTPHandler(errorMappers e.ErrorMappers) HTTPServer {
	// Error handlers
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = methodNotAllowedHandler

	// new engine
	engine := echo.New()
	engine.HidePort = true
	// engine.HideBanner = true
	engine.Binder = &BinderWithValidation{}

	// custom the error handler
	engine.HTTPErrorHandler = HTTPErrorHandler(errorMappers)

	// override the default validator
	engine.Validator = newEchoValidator()

	// set http handler
	httpHandler := &HTTPHandler{
		engine:   engine,
		routerV1: engine.Group("/api/v1"),
	}

	return httpHandler
}
