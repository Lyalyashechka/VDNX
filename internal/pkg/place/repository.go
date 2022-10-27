package place

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, place models.Place) error
	GetAll(ctx context.Context, param GetAllPlacesParam) ([]models.Place, error)
}
