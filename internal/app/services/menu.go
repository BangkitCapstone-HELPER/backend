package services

import (
	"github.com/BangkitCapstone-HELPER/backend/internal/app/model/dto"
	"github.com/BangkitCapstone-HELPER/backend/internal/app/repo"
	"go.uber.org/fx"
)

type menuServiceParams struct {
	fx.In
	MenuRepo repo.MenuRepo
}

type MenuService interface {
	CreateMenu(menu dto.CreateMenuRequestDTO) (dto.MenuDTO, error)
	GetAllMenu() ([]dto.MenuDTO, error)
	GetMenu(id uint64) (dto.MenuDTO, error)
	DeleteMenu(id uint64) (dto.MenuDTO, error)
}

func NewMenuService(params menuServiceParams) MenuService {
	return &params
}

func (u *menuServiceParams) CreateMenu(menu dto.CreateMenuRequestDTO) (dto.MenuDTO, error) {
	newMenu, err := u.MenuRepo.CreateMenu(menu.ToDAO())

	return dto.NewMenuDTO(newMenu), err
}

func (u *menuServiceParams) GetAllMenu() ([]dto.MenuDTO, error) {
	menus, err := u.MenuRepo.GetAllMenu()

	var result []dto.MenuDTO
	for _, menu := range menus {
		result = append(result, dto.NewMenuDTO(menu))
	}

	return result, err
}

func (u *menuServiceParams) GetMenu(id uint64) (dto.MenuDTO, error) {
	menu, err := u.MenuRepo.GetMenu(id)

	return dto.NewMenuDTO(menu), err
}

func (u *menuServiceParams) DeleteMenu(id uint64) (dto.MenuDTO, error) {
	menu, err := u.MenuRepo.DeleteMenu(id)

	return dto.NewMenuDTO(menu), err
}