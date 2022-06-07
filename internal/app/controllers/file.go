package controllers

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/services"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"net/http"
)

type fileControllerParams struct {
	fx.In

	FileService services.FileService
}

type FileController interface {
	//GetFile(ctx echo.Context) error
	UploadFile(ctx echo.Context) error
	PredictImage(ctx echo.Context) error
}

func NewFileController(params fileControllerParams) FileController {
	return &params
}

//// CreateOrder godoc
//// @Summary Get file
//// @Description Get file
//// @Tags file
//// @Param id path string true "file id"
//// @Success 200
//// @Router /api/v1/file/ [get]
//func (p fileControllerParams) GetFile(ctx echo.Context) error {
//	fileCode := ctx.Param("id")
//
//	fileBytes, _, err := p.FileService.GetFile(fileCode)
//	if err != nil {
//		return err
//	}
//
//	mimeType := http.DetectContentType(fileBytes)
//
//	return ctx.Blob(http.StatusOK, mimeType, fileBytes)
//}

// CreateOrder godoc
// @Summary Create file
// @Description Create file
// @Tags file
// @Accept mpfd
// @Produce  json
// @Param file formData file true "this is a test file"
// @Param folder formData string true "file folder"
// @Success 200
// @Router /api/v1/file/ [post]
func (p fileControllerParams) UploadFile(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	folder := ctx.FormValue("folder")
	if err != nil {
		return err
	}
	uploadedFile, err := p.FileService.UploadFile(*file, folder)
	if err != nil {
		return err
	}

	return lib.Response{
		Status:  http.StatusOK,
		Data:    dto.NewFileDTO(uploadedFile),
		Message: "upload file successfull",
	}.JSON(ctx)

}

// CreateOrder godoc
// @Summary Predict image
// @Description Predict image
// @Tags file
// @Accept mpfd
// @Produce  json
// @Param file formData file true "this is a test file"
// @Success 200
// @Router /api/v1/file/predict/ [post]
func (p fileControllerParams) PredictImage(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	uploadedFile, err := p.FileService.PredictImage(*file)
	if err != nil {
		return err
	}

	return lib.Response{
		Status:  http.StatusOK,
		Data:    uploadedFile,
		Message: "upload file successfull",
	}.JSON(ctx)

}
