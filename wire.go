//go:build wireinject

package main

import (
	"database/sql"

	"github.com/JY8752/go-unittest-architecture/controller"
	"github.com/JY8752/go-unittest-architecture/domain"
	"github.com/JY8752/go-unittest-architecture/infrastructure/api"
	"github.com/JY8752/go-unittest-architecture/infrastructure/handler"
	"github.com/JY8752/go-unittest-architecture/infrastructure/repository"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeRootHandler(db *sql.DB, e *echo.Echo) *handler.Root {
	wire.Build(
		repository.NewGacha,
		repository.NewItem,
		domain.NewSeedGenerator,
		api.NewPayment,
		controller.NewGacha,
		handler.NewGacha,
		handler.NewRoot,
	)
	return &handler.Root{}
}
