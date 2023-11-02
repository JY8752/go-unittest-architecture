// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/JY8752/go-unittest-architecture/controller"
	"github.com/JY8752/go-unittest-architecture/domain"
	"github.com/JY8752/go-unittest-architecture/infrastructure/handler"
	"github.com/JY8752/go-unittest-architecture/infrastructure/repository"
	"github.com/labstack/echo/v4"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeRootHandler(db *sql.DB, e *echo.Echo) *handler.Root {
	gacha := repository.NewGacha(db)
	item := repository.NewItem(db)
	seedGenerator := domain.NewSeedGenerator()
	controllerGacha := controller.NewGacha(gacha, item, seedGenerator)
	handlerGacha := handler.NewGacha(e, controllerGacha)
	root := handler.NewRoot(handlerGacha)
	return root
}
