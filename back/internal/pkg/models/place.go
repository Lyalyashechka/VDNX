package models

type WorkSchedule struct {
	Monday    string `json:"Пн" mapstructure:"Пн"`
	Tuesday   string `json:"Вт" mapstructure:"Вт"`
	Wednesday string `json:"Ср" mapstructure:"Ср"`
	Thursday  string `json:"Чт" mapstructure:"Чт"`
	Friday    string `json:"Пт" mapstructure:"Пт"`
	Saturday  string `json:"Сб" mapstructure:"Сб"`
	Sunday    string `json:"Вс" mapstructure:"Вс"`
}

type Place struct {
	Qr               string          `json:"qr,omitempty"  mapstructure:"qr"`
	Id               int             `json:"id" mapstructure:"id"`
	Order            int             `json:"order,omitempty" mapstructure:"order"`
	Coordinates      []float64       `json:"coordinates,omitempty" mapstructure:"coordinates"`
	Schedule         []SchedulePlace `json:"schedule,omitempty" mapstructure:"schedule"`
	Contacts         []Contacts      `json:"contacts,omitempty" mapstructure:"contacts"`
	Facilities       []string        `json:"facilities,omitempty" mapstructure:"facilities"`
	VideoRoute       string          `json:"video_route,omitempty" mapstructure:"video_route"`
	Tickets          []Ticket        `json:"tickets,omitempty" mapstructure:"tickets"`
	TicketsLink      string          `json:"tickets_link,omitempty" mapstructure:"tickets_link"`
	TicketsText      string          `json:"tickets_text,omitempty" mapstructure:"tickets_text"`
	Cat              string          `json:"cat,omitempty" mapstructure:"cat"`
	Visibility       string          `json:"visibility,omitempty" mapstructure:"visibility"`
	Color            string          `json:"color,omitempty" mapstructure:"color"`
	ColorCode        string          `json:"color_code,omitempty" mapstructure:"color_code"`
	PreviewText      string          `json:"preview_text,omitempty" mapstructure:"preview_text"`
	DetailText       string          `json:"default_text,omitempty" mapstructure:"detail_text"`
	Title            string          `json:"title,omitempty" mapstructure:"title"`
	Time             string          `json:"time,omitempty" mapstructure:"time"`
	Type             string          `json:"type,omitempty" mapstructure:"type"`
	Url              string          `json:"url,omitempty" mapstructure:"url"`
	Pic              string          `json:"pic,omitempty" mapstructure:"pic"`
	Code             string          `json:"code,omitempty" mapstructure:"code"`
	Places           []string        `json:"places,omitempty" mapstructure:"places"`
	WorkSchedule     WorkSchedule    `json:"work_schedule" mapstructure:"work_schedule"`
	ElectrobusFlag   int             `json:"electrobus_flag" mapstructure:"electrobus_flag"`
	ScooterFlag      int             `json:"scooter_flag" mapstructure:"scooter_flag"`
	Free             int             `json:"free" mapstructure:"free"`
	Rating           int64           `json:"rating" mapstructure:"rating"`
	ChildrenFlag     int             `json:"children_flag" mapstructure:"children_flag"`
	DateFlag         int             `json:"date_flag" mapstructure:"date_flag"`
	FamilyFlag       int             `json:"family_flag" mapstructure:"family_flag"`
	AnimalsFlag      int             `json:"animals_flag" mapstructure:"animals_flag"`
	BikeFlag         int             `json:"bike_flag" mapstructure:"bike_flag"`
	Promenade        int             `json:"promenade" mapstructure:"promenade"`
	Activity         int             `json:"activity" mapstructure:"activity"`
	Nature           int             `json:"nature" mapstructure:"nature"`
	Science          int             `json:"science" mapstructure:"science"`
	National         int             `json:"national" mapstructure:"national"`
	Workshop         int             `json:"workshop" mapstructure:"workshop"`
	Creation         int             `json:"creation" mapstructure:"creation"`
	Kids             int             `json:"kids" mapstructure:"kids"`
	Tech             int             `json:"tech" mapstructure:"tech"`
	AboutRussia      int             `json:"about_russia" mapstructure:"about_russia"`
	TypeService      string          `json:"type_service" mapstructure:"type_service"`
	Duration         int64           `json:"duration" mapstructure:"duration"`
	ProcessedPreview string          `json:"processed_preview" mapstructure:"processed_preview"`
	ProcessedTitle   string          `json:"processed_title" mapstructure:"processed_title"`
	VectorText       []float64       `json:"vector_text" mapstructure:"vector_text"`
	VectorTitle      []float64       `json:"vector_title" mapstructure:"vector_title"`
	IsEvent          bool            `json:"is_event"`
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
