package models

type Place struct {
	Qr          string          `json:"qr,omitempty"`
	Id          int             `json:"id"`
	Order       int             `json:"order,omitempty"`
	Coordinates []float64       `json:"coordinates,omitempty"`
	Schedule    []SchedulePlace `json:"schedule,omitempty"`
	Contacts    []Contacts      `json:"contacts,omitempty"`
	Facilities  []string        `json:"facilities,omitempty"`
	VideoRoute  string          `json:"video_route,omitempty"`
	Tickets     []Ticket        `json:"tickets,omitempty"`
	TicketsLink string          `json:"tickets_link,omitempty"`
	TicketsText string          `json:"tickets_text,omitempty"`
	Cat         string          `json:"cat,omitempty"`
	Visibility  string          `json:"visibility,omitempty"`
	Color       string          `json:"color,omitempty"`
	ColorCode   string          `json:"color_code,omitempty"`
	PreviewText string          `json:"preview_text,omitempty"`
	DefaultText string          `json:"default_text,omitempty"`
	Title       string          `json:"title,omitempty"`
	Time        string          `json:"time,omitempty"`
	Type        string          `json:"type,omitempty"`
	Url         string          `json:"url,omitempty"`
	Pic         string          `json:"pic,omitempty"`
	Code        string          `json:"code,omitempty"`
}

type SchedulePlace struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Contacts struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type Ticket struct {
	Title       string `json:"title,omitempty"`
	Value       string `json:"value,omitempty"`
	Description string `json:"description"`
}
