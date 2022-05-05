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
		Image    string       `json:"image"`
		Day      string       `json:"day"`
		Contents []ContentDTO `json:"contents"`
	}

	ContentDTO struct {
		Kind  string    `json:"kind"`
		Items []ItemDTO `json:"items"`
	}

	ItemDTO struct {
		Name string `json:"name"`
	}
)

func (c CreateMenuRequestDTO) ToDAO() dao.Menu {
	dayMenus := []dao.DayMenu{}

	for _, dayMenu := range c.DayMenus {
		newDayMenu := dao.DayMenu{
			Day:   dayMenu.Day,
			Image: dayMenu.Image,
		}
		contents := []dao.Content{}

		for _, content := range dayMenu.Contents {
			newContent := dao.Content{
				Kind: content.Kind,
			}

			items := []dao.Item{}

			for _, item := range content.Items {
				newItem := dao.Item{
					Name: item.Name,
				}
				items = append(items, newItem)
			}
			newContent.Items = items
			contents = append(contents, newContent)
		}
		newDayMenu.Contents = contents
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
		contents := []ContentDTO{}

		for _, content := range dayMenu.Contents {
			newContent := ContentDTO{
				Kind: content.Kind,
			}

			items := []ItemDTO{}

			for _, item := range content.Items {
				newItem := ItemDTO{
					Name: item.Name,
				}
				items = append(items, newItem)
			}
			newContent.Items = items
			contents = append(contents, newContent)
		}
		newDayMenu.Contents = contents
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
