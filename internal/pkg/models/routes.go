package models

type Route struct {
	Places []Place
}

type Routes struct {
	Routes []Route `json:"routes"`
}
