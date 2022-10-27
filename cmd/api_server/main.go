package main

import (
	"github.com/Lyalyashechka/VDNX/config"
	"github.com/Lyalyashechka/VDNX/config/config_middleware"
	"github.com/Lyalyashechka/VDNX/config/config_routing"
	place_hanlder "github.com/Lyalyashechka/VDNX/internal/pkg/place/delivery/http"
	place_repository "github.com/Lyalyashechka/VDNX/internal/pkg/place/repository/postgres"
	place_usecase "github.com/Lyalyashechka/VDNX/internal/pkg/place/usecase"
	"github.com/Lyalyashechka/VDNX/internal/pkg/postgres"
	tools_logger "github.com/Lyalyashechka/VDNX/internal/pkg/tools/logger"
	upload_handler "github.com/Lyalyashechka/VDNX/internal/pkg/upload_data/delivery/http"
	upload_usecase "github.com/Lyalyashechka/VDNX/internal/pkg/upload_data/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	router := echo.New()
	logger := tools_logger.New(os.Stdout, &logrus.JSONFormatter{}, logrus.InfoLevel)
	configApp := config.Config{}
	err := config.LoadConfig(&configApp, "config")
	if err != nil {
		panic(err)
	}

	db, err := postgres.InitDB(configApp.Postgres)
	if err != nil {
		logger.Fatal(err)
	}

	placeRepository := place_repository.New(db, logger)
	placeUseCase := place_usecase.New(logger, placeRepository)
	placeHandler := place_hanlder.New(logger, placeUseCase)

	uploadUseCase := upload_usecase.New(logger, placeRepository)
	uploadHandler := upload_handler.New(logger, uploadUseCase)

	configRouting := config_routing.ServerConfigRouting{
		UploadHandler: uploadHandler,
		PlaceHandler:  placeHandler,
	}

	config_middleware.ConfigMiddleware(router)
	configRouting.ConfigRouting(router)

	router.GET("hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	router.Logger.Fatal(router.Start("localhost:8091"))
}
