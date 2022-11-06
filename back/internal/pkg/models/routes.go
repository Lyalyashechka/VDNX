package models

type Route struct {
	Places []Place `json:"places"`
}

type Routes struct {
	Routes []Route `json:"routes"`
}
