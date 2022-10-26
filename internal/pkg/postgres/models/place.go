package models

import "github.com/lib/pq"

type Place struct {
	Id          int             `gorm:"column:id"`
	Qr          string          `gorm:"column:qr"`
	Order       int             `gorm:"column:order_place"`
	Coordinates pq.Float64Array `gorm:"column:coordinates"`
	Facilities  pq.StringArray  `gorm:"column:facilities"`
	VideoRoute  string          `gorm:"column:video_route"`
	TicketsLink string          `gorm:"column:tickets_link"`
	TicketsText string          `gorm:"column:tickets_text"`
	Cat         string          `gorm:"column:cat"`
	Visibility  string          `gorm:"column:visibility"`
	Color       string          `gorm:"column:color"`
	ColorCode   string          `gorm:"column:color_code"`
	PreviewText string          `gorm:"column:preview_text"`
	DefaultText string          `gorm:"column:default_text"`
	Title       string          `gorm:"column:title"`
	Time        string          `gorm:"column:time_work"`
	Type        string          `gorm:"column:type_place"`
	Url         string          `gorm:"column:url"`
	Pic         string          `gorm:"column:pic"`
	Code        string          `gorm:"column:code"`
}

func (Place) TableName() string {
	return "places"
}
