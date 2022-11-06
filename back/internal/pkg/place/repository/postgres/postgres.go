package postgres

import (
	"github.com/Lyalyashechka/VDNX/internal/pkg/place"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Postgres struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func New(db *gorm.DB, logger *logrus.Logger) place.Repository {
	return &Postgres{
		db:     db,
		logger: logger,
	}
}
