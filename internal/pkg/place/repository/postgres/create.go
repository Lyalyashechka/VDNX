package postgres

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	db_models "github.com/Lyalyashechka/VDNX/internal/pkg/postgres/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (p *Postgres) Create(ctx context.Context, place models.Place) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		dbPlace := db_models.Place{
			Id:          place.Id,
			Qr:          place.Qr,
			Order:       place.Order,
			Coordinates: place.Coordinates,
			Facilities:  place.Facilities,
			VideoRoute:  place.VideoRoute,
			TicketsLink: place.TicketsLink,
			TicketsText: place.TicketsText,
			Cat:         place.Cat,
			Visibility:  place.Visibility,
			Color:       place.Color,
			ColorCode:   place.ColorCode,
			PreviewText: place.PreviewText,
			DetailText:  place.DetailText,
			Title:       place.Title,
			Time:        place.Time,
			Type:        place.Type,
			Url:         place.Url,
			Pic:         place.Pic,
			Code:        place.Code,
		}

		res := p.db.Table("places").Create(&dbPlace)
		if err := res.Error; err != nil {
			return errors.Wrapf(err, "failed to create place")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "failed to make transaction")
	}

	return nil
}
