package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/JY8752/go-unittest-architecture/controller"
	"github.com/JY8752/go-unittest-architecture/domain"
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
	g.e.POST("/gacha/:gachaId/draw", func(c echo.Context) error {
		gachaId, err := strconv.Atoi(c.Param("gachaId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, NewErrorRes(
				WithMessage("gachaId parameter is invalid."),
				WithErr(err),
			))
		}

		item, err := g.gc.Draw(context.Background(), domain.NewGachaId(gachaId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, NewErrorRes(
				WithMessage("fail to draw gacha."),
				WithErr(err),
			))
		}

		return c.JSON(http.StatusOK, item)
	})
}
