package domain

type ItemId struct {
	ValueObject[int64]
}

func NewItemId(itemId int64) ItemId {
	return ItemId{ValueObject[int64]{itemId}}
}

type Rarity string

const (
	N  = Rarity("N")
	R  = Rarity("R")
	SR = Rarity("SR")
)

type Item struct {
	Id     ItemId `json:"item_id"`
	Name   string `json:"name"`
	Rarity Rarity `json:"rarity"`
}

func NewItem(id ItemId, name string, rarity Rarity) *Item {
	return &Item{id, name, rarity}
}
