package entity

type Menus []Menu

type Menu struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Routes      Routes `json:"routes"`
}
