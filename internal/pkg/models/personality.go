package models

type PersonInfoRoute struct {
	With      string   `json:"with"`
	Children  bool     `json:"children"`
	Interests []string `json:"interests"`
	Transport string   `json:"transport"`
	Time      string   `json:"time"`
}
