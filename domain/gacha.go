package domain

import (
	"errors"
	"math/rand"
	"time"
)

// value object
type GachaId struct {
	ValueObject[int]
}

func NewGachaId(v int) GachaId {
	return GachaId{ValueObject[int]{v}}
}

// collaborator
type SeedGenerator interface {
	New() int64
}

type seedGenerator struct{}

func NewSeedGenerator() SeedGenerator {
	return &seedGenerator{}
}

func (s *seedGenerator) New() int64 {
	return time.Now().UnixNano()
}

// domain model
type GachaItemWeights []struct {
	ItemId ItemId
	Weight int
}

type Gacha struct {
	Weights GachaItemWeights
}

func NewGacha(weights GachaItemWeights) *Gacha {
	return &Gacha{weights}
}

func (g *Gacha) Draw(seed int64) (*ItemId, error) {
	weights := make([]int, len(g.Weights))
	for i, w := range g.Weights {
		weights[i] = w.Weight
	}

	index, err := linearSearchLottery(weights, seed)
	if err != nil {
		return nil, err
	}

	return &g.Weights[index].ItemId, nil
}

/*
線形探索で重み付抽選する
@return 当選した要素のインデックス
*/
func linearSearchLottery(weights []int, seed int64) (int, error) {
	//  重みの総和を取得する
	var total int
	for _, weight := range weights {
		total += weight
	}

	// 乱数取得
	r := rand.New(rand.NewSource(seed))
	rnd := r.Intn(total)

	var currentWeight int
	for i, w := range weights {
		// 現在要素までの重みの総和
		currentWeight += w

		if rnd <= currentWeight {
			return i, nil
		}
	}

	// たぶんありえない
	return 0, errors.New("the lottery failed")
}
