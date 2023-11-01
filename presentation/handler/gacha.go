package handler

import (
	"context"

	"github.com/JY8752/go-unittest-architecture/presentation/controller"
	"github.com/labstack/echo/v4"
)

type Gacha struct {
	e  *echo.Echo
	gc *controller.Gacha
}

func NewGacha(e *echo.Echo, gc *controller.Gacha) *Gacha {
	return &Gacha{e, gc}
}

func (g *Gacha) Register() {
	g.draw()
}

func (g *Gacha) draw() {
	g.e.POST("/gacha/draw", func(c echo.Context) error {
		_, err := g.gc.Draw(context.Background(), 1)
		if err != nil {
			return err
		}
		return nil
	})
}
