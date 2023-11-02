package controller

import (
	"context"

	"github.com/JY8752/go-unittest-architecture/domain"
	"github.com/JY8752/go-unittest-architecture/infrastructure/repository"
)

type Gacha struct {
	gachaRep      *repository.Gacha
	itemRep       *repository.Item
	seedGenerator domain.SeedGenerator
}

func NewGacha(gachaRep *repository.Gacha, itemRep *repository.Item, seedGenerator domain.SeedGenerator) *Gacha {
	return &Gacha{gachaRep, itemRep, seedGenerator}
}

func (g *Gacha) Draw(ctx context.Context, gachaId domain.GachaId) (*domain.Item, error) {
	// 指定ガチャの重み一覧を取得
	weights, err := g.gachaRep.GetGachaItemWeights(ctx, gachaId)
	if err != nil {
		return nil, err
	}

	gacha := domain.NewGacha(weights)
	seed := g.seedGenerator.New()

	// ビジネスロジック ガチャの抽選
	itemId, err := gacha.Draw(seed)
	if err != nil {
		return nil, err
	}

	// 抽選したアイテム情報を取得
	item, err := g.itemRep.FindById(ctx, itemId)
	if err != nil {
		return nil, err
	}

	// 管理下にないプロセス外依存 決済する

	return item, err
}
