package model

type Routes []Route

type Route struct {
	Title string `json:"title"`
	Path  string `json:"path"`
}
