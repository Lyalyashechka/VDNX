package models

type PersonInfoRoute struct {
	With      string   `json:"with"`
	Kids      int      `json:"kids"`
	Animals   int      `json:"animals"`
	Interests []string `json:"interests"`
	Transport string   `json:"transport"`
}
