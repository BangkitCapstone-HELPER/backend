package dto

import "github.com/BangkitCapstone-HELPER/backend/internal/app/model/dao"

type (
	CreateMenuRequestDTO struct {
		Title       string       `json:"title"`
		Price       int          `json:"price"`
		Description string       `json:"description"`
		DayMenus    []DayMenuDTO `json:"day_menus"`
	}
	MenuDTO struct {
		ID          uint         `json:"id"`
		Title       string       `json:"title"`
		Price       int          `json:"price"`
		Description string       `json:"description"`
		DayMenus    []DayMenuDTO `json:"day_menus"`
	}

	DayMenuDTO struct {
		Image string   `json:"image"`
		Day   string   `json:"day"`
		Items []string `json:"items"`
	}

	//ItemDTO struct {
	//	Name string `json:"name"`
	//}
)

func (c CreateMenuRequestDTO) ToDAO() dao.Menu {
	dayMenus := []dao.DayMenu{}

	for _, dayMenu := range c.DayMenus {
		newDayMenu := dao.DayMenu{
			Day:   dayMenu.Day,
			Image: dayMenu.Image,
		}
		//items := []dao.Item{}
		//
		//for _, item := range dayMenu.Items {
		//	newItem := dao.Item{
		//		Name: item.Name,
		//	}
		//	items = append(items, newItem)
		//}

		newDayMenu.Items = dayMenu.Items
		dayMenus = append(dayMenus, newDayMenu)
	}

	return dao.Menu{
		Title:       c.Title,
		Price:       c.Price,
		Description: c.Description,
		DayMenus:    dayMenus,
	}
}

func NewMenuDTO(menu dao.Menu) MenuDTO {
	dayMenus := []DayMenuDTO{}

	for _, dayMenu := range menu.DayMenus {
		newDayMenu := DayMenuDTO{
			Day:   dayMenu.Day,
			Image: dayMenu.Image,
		}

		//items := []ItemDTO{}
		//
		//for _, item := range dayMenu.Items {
		//	newItem := ItemDTO{
		//		Name: item.Name,
		//	}
		//	items = append(items, newItem)
		//}
		newDayMenu.Items = dayMenu.Items
		dayMenus = append(dayMenus, newDayMenu)
	}
	return MenuDTO{
		ID:          menu.ID,
		Title:       menu.Title,
		Price:       menu.Price,
		Description: menu.Description,
		DayMenus:    dayMenus,
	}
}
