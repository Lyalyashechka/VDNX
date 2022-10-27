package place

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

type GetAllPlacesParam struct {
	IsEvent models.Bool
}

type UseCase interface {
	GetAllPlaces(ctx context.Context, param GetAllPlacesParam) ([]models.Place, error)
	GetPlaceById(ctx context.Context, id int) (models.Place, error)
}
