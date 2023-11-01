package controller

import (
	"context"

	"github.com/JY8752/go-unittest-architecture/domain"
	"github.com/JY8752/go-unittest-architecture/infrastructure"
)

type Gacha struct {
	gachaRep *infrastructure.GachaRepostory
}

func NewGacha(gachaRep *infrastructure.GachaRepostory) *Gacha {
	return &Gacha{gachaRep}
}

func (g *Gacha) Draw(ctx context.Context, gachaId int) (*domain.Item, error) {
	weights, err := g.gachaRep.GetGachaItemWeights(ctx, gachaId)
	if err != nil {
		return nil, err
	}

	gacha := domain.Gacha{Weights: weights}

	_, err = gacha.Draw()
	if err != nil {
		return nil, err
	}

	// 決済する

	// 抽選したアイテムを記録する

	return nil, err
}
