package model

type MenuModels []MenuModel

type MenuModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Routes      Routes `json:"routes"`
}
