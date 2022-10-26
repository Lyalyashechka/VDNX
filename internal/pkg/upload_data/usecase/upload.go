package usecase

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

func (uc *UseCase) CreatePlaces(ctx context.Context, places []models.Place) error {
	for _, place := range places {
		err := uc.placeRepository.Create(ctx, place)
		if err != nil {
			uc.logger.WithError(err).Errorf("[CreatePlaces] error createPlace")
			return err
		}
	}

	return nil
}
