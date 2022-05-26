package dto

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
)

type (
	FileDTO struct {
		FileCode         string `json:"file_code"`
		OriginalFileName string `json:"file_name"`
		Extension        string `json:"file_extension"`
		PublicUrl        string `json:"url"`
	}
)

func NewFileDTO(file dao.File) FileDTO {

	return FileDTO{
		FileCode:         file.FileCode,
		OriginalFileName: file.OriginalFileName,
		Extension:        file.Extension,
		PublicUrl:        "https://storage.googleapis.com/serantau/" + file.FileCode + file.Extension,
	}
}
