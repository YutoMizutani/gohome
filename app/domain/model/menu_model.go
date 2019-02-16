package model

type MenuModels []MenuModel

type MenuModel struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	NumberOfContents int    `json:"number_of_contents"`
	ImageURL         string `json:"image_url"`
}
