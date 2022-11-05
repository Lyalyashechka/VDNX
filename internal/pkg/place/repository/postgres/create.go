package postgres

import (
	"context"
	"encoding/json"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	db_models "github.com/Lyalyashechka/VDNX/internal/pkg/postgres/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sort"
	"strconv"
)

func (p *Postgres) Create(ctx context.Context, place models.Place) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		jsonWorkShedule, err := json.Marshal(place.WorkSchedule)
		if err != nil {
			p.logger.WithError(err).Errorf("[Create] error marshal")
			return err
		}
		sort.Sort(sort.Reverse(sort.Float64Slice(place.Coordinates)))
		dbPlace := db_models.Place{
			Id:               place.Id,
			Qr:               place.Qr,
			Order:            place.Order,
			Coordinates:      place.Coordinates,
			Facilities:       place.Facilities,
			VideoRoute:       place.VideoRoute,
			TicketsLink:      place.TicketsLink,
			TicketsText:      place.TicketsText,
			Cat:              place.Cat,
			Visibility:       place.Visibility,
			Color:            place.Color,
			ColorCode:        place.ColorCode,
			PreviewText:      place.PreviewText,
			DetailText:       place.DetailText,
			Title:            place.Title,
			Time:             place.Time,
			Type:             place.Type,
			Url:              place.Url,
			Pic:              place.Pic,
			Code:             place.Code,
			IsEvent:          place.IsEvent,
			ElectrobusFlag:   place.ElectrobusFlag,
			ScooterFlag:      place.ScooterFlag,
			Free:             place.Free,
			Rating:           place.Rating,
			ChildrenFlag:     place.ChildrenFlag,
			DateFlag:         place.DateFlag,
			FamilyFlag:       place.FamilyFlag,
			AnimalsFlag:      place.AnimalsFlag,
			BikeFlag:         place.BikeFlag,
			Promenade:        place.Promenade,
			Activity:         place.Activity,
			Nature:           place.Nature,
			Science:          place.Science,
			National:         place.National,
			Workshop:         place.Workshop,
			Creation:         place.Creation,
			Kids:             place.Kids,
			Tech:             place.Tech,
			AboutRussia:      place.AboutRussia,
			TypeService:      place.TypeService,
			Duration:         place.Duration,
			ProcessedPreview: place.ProcessedPreview,
			ProcessedTitle:   place.ProcessedTitle,
			VectorText:       place.VectorText,
			VectorTitle:      place.VectorTitle,
			WorkSchedule:     string(jsonWorkShedule),
		}

		res := p.db.Table("places").Create(&dbPlace)
		if err := res.Error; err != nil {
			return errors.Wrapf(err, "failed to create place")
		}

		if dbPlace.IsEvent {
			for _, v := range place.Places {
				placeId, err := strconv.Atoi(v)
				if err != nil {
					return errors.Wrapf(err, "failed convert place id to int")
				}

				dbPlaceEventSharing := db_models.PlaceEventSharing{
					EventId: dbPlace.Id,
					PlaceId: placeId,
				}
				res = p.db.Create(&dbPlaceEventSharing)
				if err := res.Error; err != nil {
					return errors.Wrap(err, "failed to create place event sharing")
				}
			}

		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "failed to make transaction")
	}

	return nil
}
