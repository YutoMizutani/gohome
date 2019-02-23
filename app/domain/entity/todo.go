package entity

import "github.com/jinzhu/gorm"

type TodoList []Todo

type Todo struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Tags        []*TodoTag `json:"tags" gorm:"default:[]"`
	IsDone      bool       `json:"is_done"`
}

type TodoTag struct {
	gorm.Model
	Tag string `json:"tag"`
}
