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
	CreateMenu(user dto.CreateMenuRequestDTO) (dto.MenuDTO, error)
	GetAllMenu() ([]dto.MenuDTO, error)
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
