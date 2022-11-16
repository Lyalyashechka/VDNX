package upload_data

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

type UseCase interface {
	CreatePlaces(ctx context.Context, places []models.Place) error
}
