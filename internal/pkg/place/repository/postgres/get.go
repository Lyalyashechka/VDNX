package postgres

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
	db_models "github.com/Lyalyashechka/VDNX/internal/pkg/postgres/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (p *Postgres) GetAll(ctx context.Context, param place.GetAllPlacesParam) ([]models.Place, error) {
	var result []models.Place

	err := p.db.Transaction(func(tx *gorm.DB) error {
		var dbPlaces []db_models.Place

		condition := ""
		if param.IsEvent.IsDefinedTrue() {
			condition += "places.is_event = true"
		} else if param.IsEvent.IsDefinedFalse() {
			condition += "places.is_event = false"
		}

		res := p.db.Table("places").Find(&dbPlaces, condition)
		if err := res.Error; err != nil {
			return errors.Wrap(err, "failed to get places")
		}

		result = make([]models.Place, 0, len(dbPlaces))
		for _, dbPlace := range dbPlaces {
			place := models.Place{
				Qr:          dbPlace.Qr,
				Id:          dbPlace.Id,
				Order:       dbPlace.Order,
				Coordinates: dbPlace.Coordinates,
				Facilities:  dbPlace.Facilities,
				VideoRoute:  dbPlace.VideoRoute,
				TicketsLink: dbPlace.TicketsLink,
				TicketsText: dbPlace.TicketsText,
				Cat:         dbPlace.Cat,
				Visibility:  dbPlace.Visibility,
				Color:       dbPlace.Color,
				ColorCode:   dbPlace.ColorCode,
				PreviewText: dbPlace.PreviewText,
				DetailText:  dbPlace.DetailText,
				Title:       dbPlace.Title,
				Time:        dbPlace.Time,
				Type:        dbPlace.Type,
				Url:         dbPlace.Url,
				Pic:         dbPlace.Pic,
				Code:        dbPlace.Code,
				IsEvent:     dbPlace.IsEvent,
			}

			result = append(result, place)
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to make transaction")
	}

	return result, nil
}
