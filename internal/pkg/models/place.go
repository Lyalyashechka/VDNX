package models

type Place struct {
	Qr          string          `json:"qr,omitempty"  mapstructure:"qr"`
	Id          int             `json:"id" mapstructure:"id"`
	Order       int             `json:"order,omitempty" mapstructure:"order"`
	Coordinates []float64       `json:"coordinates,omitempty" mapstructure:"coordinates"`
	Schedule    []SchedulePlace `json:"schedule,omitempty" mapstructure:"schedule"`
	Contacts    []Contacts      `json:"contacts,omitempty" mapstructure:"contacts"`
	Facilities  []string        `json:"facilities,omitempty" mapstructure:"facilities"`
	VideoRoute  string          `json:"video_route,omitempty" mapstructure:"video_route"`
	Tickets     []Ticket        `json:"tickets,omitempty" mapstructure:"tickets"`
	TicketsLink string          `json:"tickets_link,omitempty" mapstructure:"tickets_link"`
	TicketsText string          `json:"tickets_text,omitempty" mapstructure:"tickets_text"`
	Cat         string          `json:"cat,omitempty" mapstructure:"cat"`
	Visibility  string          `json:"visibility,omitempty" mapstructure:"visibility"`
	Color       string          `json:"color,omitempty" mapstructure:"color"`
	ColorCode   string          `json:"color_code" mapstructure:"color_code"`
	PreviewText string          `json:"preview_text,omitempty" mapstructure:"preview_text"`
	DetailText  string          `json:"default_text,omitempty" mapstructure:"detail_text"`
	Title       string          `json:"title,omitempty" mapstructure:"title"`
	Time        string          `json:"time,omitempty" mapstructure:"time"`
	Type        string          `json:"type,omitempty" mapstructure:"type"`
	Url         string          `json:"url,omitempty" mapstructure:"url"`
	Pic         string          `json:"pic,omitempty" mapstructure:"pic"`
	Code        string          `json:"code,omitempty" mapstructure:"code"`
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
