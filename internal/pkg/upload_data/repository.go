package upload_data

import (
	"context"
	"github.com/Lyalyashechka/VDNX/internal/pkg/models"
)

type Repository interface {
	UploadPlaces(ctx context.Context, places []models.Place) error
}
