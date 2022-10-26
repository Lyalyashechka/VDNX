package place

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, place models.Place) error
}
