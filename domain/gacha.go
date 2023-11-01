package domain

type GachaId int
type ItemId int64

type GachaItemWeights []struct {
	ItemId ItemId
	Weight int
}

type Gacha struct {
	Weights GachaItemWeights
}

type Item struct {
	Id     ItemId
	Name   string
	Rarity string
}

func (g *Gacha) Draw() (ItemId, error) {
	// TODO 抽選ロジック
	return ItemId(0), nil
}
