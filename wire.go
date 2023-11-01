//go:build wireinject

package main

import (
	"database/sql"

	"github.com/JY8752/go-unittest-architecture/infrastructure"
	"github.com/JY8752/go-unittest-architecture/presentation/controller"
	"github.com/JY8752/go-unittest-architecture/presentation/handler"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeRootHandler(db *sql.DB, e *echo.Echo) *handler.Root {
	wire.Build(
		infrastructure.NewGachaRepository,
		controller.NewGacha,
		handler.NewGacha,
		handler.NewRoot,
	)
	return &handler.Root{}
}
