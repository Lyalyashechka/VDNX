package postgres

import (
	"github.com/Lyalyashechka/VDNX/internal/pkg/upload_data"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Postgres struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func New(db *gorm.DB, logger *logrus.Logger) upload_data.Repository {
	return &Postgres{
		logger: logger,
		db:     db,
	}
}
