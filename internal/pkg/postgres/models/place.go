package models

import "github.com/lib/pq"

type Place struct {
	Id               int             `gorm:"column:id"`
	Qr               string          `gorm:"column:qr"`
	Order            int             `gorm:"column:order_place"`
	Coordinates      pq.Float64Array `gorm:"column:coordinates"`
	Facilities       pq.StringArray  `gorm:"column:facilities"`
	VideoRoute       string          `gorm:"column:video_route"`
	TicketsLink      string          `gorm:"column:tickets_link"`
	TicketsText      string          `gorm:"column:tickets_text"`
	Cat              string          `gorm:"column:cat"`
	Visibility       string          `gorm:"column:visibility"`
	Color            string          `gorm:"column:color"`
	ColorCode        string          `gorm:"column:color_code"`
	PreviewText      string          `gorm:"column:preview_text"`
	DetailText       string          `gorm:"column:detail_text"`
	Title            string          `gorm:"column:title"`
	Time             string          `gorm:"column:time_work"`
	Type             string          `gorm:"column:type_place"`
	Url              string          `gorm:"column:url"`
	Pic              string          `gorm:"column:pic"`
	Code             string          `gorm:"column:code"`
	ElectrobusFlag   int             `gorm:"column:electrobus_flag"`
	ScooterFlag      int             `gorm:"column:scooter_flag"`
	Free             int             `gorm:"column:free"`
	Rating           int64           `gorm:"column:rating"`
	ChildrenFlag     int             `gorm:"column:children_flag"`
	DateFlag         int             `gorm:"column:date_flag"`
	FamilyFlag       int             `gorm:"column:family_flag"`
	AnimalsFlag      int             `gorm:"column:animals_flag"`
	BikeFlag         int             `gorm:"column:bike_flag"`
	Promenade        int             `gorm:"column:promenade"`
	Activity         int             `gorm:"column:activity"`
	Nature           int             `gorm:"column:nature"`
	Science          int             `gorm:"column:science"`
	National         int             `gorm:"column:national"`
	Workshop         int             `gorm:"column:workshop"`
	Creation         int             `gorm:"column:creation"`
	Kids             int             `gorm:"column:kids"`
	Tech             int             `gorm:"column:tech"`
	AboutRussia      int             `gorm:"column:about_russia"`
	TypeService      string          `gorm:"column:type_service"`
	Duration         int64           `gorm:"column:duration"`
	ProcessedPreview string          `gorm:"column:processed_preview"`
	ProcessedTitle   string          `gorm:"column:processed_title"`
	WorkSchedule     string          `gorm:"column:work_schedule"`
	VectorText       pq.Float64Array `gorm:"column:vector_text"`
	VectorTitle      pq.Float64Array `gorm:"column:vector_title"`

	IsEvent bool `gorm:"column:is_event"`
}

func (Place) TableName() string {
	return "places"
}

type PlaceEventSharing struct {
	Id      int `gorm:"column:id"`
	PlaceId int `gorm:"column:place_id"`
	EventId int `gorm:"column:event_id"`
}

func (PlaceEventSharing) TableName() string {
	return "place_event_sharing"
}
