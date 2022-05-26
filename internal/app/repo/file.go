package repo

import (
	"errors"
	e "github.com/BangkitCapstone-HELPER/backend/internal/app/error"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/lib"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type FileRepo interface {
	GetFile(code string) (dao.File, error)
	AddFile(file dao.File) (dao.File, error)
}
type fileRepoParams struct {
	fx.In

	lib.Database
}

func NewFileRepo(params fileRepoParams) FileRepo {
	return &params
}

func (p *fileRepoParams) GetFile(code string) (dao.File, error) {
	file := dao.File{}

	if err := p.Db.Where("file_code = ?", code).First(&file).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dao.File{}, e.ErrFileNotFound
		}

		return dao.File{}, err
	}
	return file, nil
}

func (p *fileRepoParams) AddFile(file dao.File) (dao.File, error) {
	if err := p.Db.Create(&file).Error; err != nil {
		return dao.File{}, err
	}

	return file, nil
}
