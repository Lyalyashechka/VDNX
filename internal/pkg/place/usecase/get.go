package usecase

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
)

func (us *UseCase) GetAllPlaces(ctx context.Context, param place.GetAllPlacesParam) ([]models.Place, error) {
	places, err := us.repository.GetAll(ctx, param)
	if err != nil {
		us.logger.WithError(err).Errorf("[GetAllPlaces] error in usecase")
		return []models.Place{}, nil
	}

	if places == nil {
		places = []models.Place{}
	}

	return places, nil
}

func (us *UseCase) GetPlaceById(ctx context.Context, id int) (models.Place, error) {
	place, err := us.repository.GetPlaceById(ctx, id)
	if err != nil {
		us.logger.WithError(err).Errorf("[GetPlaceById] Faile parse event id")
		return place, err
	}

	return place, nil
}
