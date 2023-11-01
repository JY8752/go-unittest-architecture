package infrastructure

import (
	"context"
	"database/sql"

	"github.com/JY8752/go-unittest-architecture/db"
	"github.com/JY8752/go-unittest-architecture/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GachaRepostory struct {
	exec boil.ContextExecutor
}

func NewGachaRepository(db *sql.DB) *GachaRepostory {
	return &GachaRepostory{db}
}

func (g *GachaRepostory) GetGachaItemWeights(ctx context.Context, gid int) (domain.GachaItemWeights, error) {
	gachaItems, err := db.GachaItems(
		qm.Select(db.GachaItemColumns.ItemID, db.GachaItemColumns.Weight),
		db.GachaItemWhere.GachaID.EQ(gid),
	).All(ctx, g.exec)

	if err != nil {
		return nil, err
	}

	gachaItemWeights := make(domain.GachaItemWeights, len(gachaItems))
	for i, item := range gachaItems {
		gachaItemWeights[i] = struct {
			ItemId domain.ItemId
			Weight int
		}{
			ItemId: domain.ItemId(item.ItemID),
			Weight: item.Weight,
		}
	}

	return gachaItemWeights, nil
}

// type GachaItem struct {
// 	Id       int    `boil:"gachas.id"`
// 	Name     string `boil:"gachas.name"`
// 	Weight   int    `boil:"gacha_items.weight"`
// 	ItemId   int64  `boil:"gacha_items.item_id"`
// 	ItemName string `boil:"items.name"`
// }

// func (g *GachaRepostory) GetGachaItems(ctx context.Context) ([]domain.GachaItem, error) {
// 	q := db.GachaItems(
// 		qm.Select(
// 			db.GachaTableColumns.ID,
// 			db.GachaTableColumns.Name,
// 			db.GachaItemTableColumns.Weight,
// 			db.GachaItemTableColumns.ItemID,
// 			db.ItemTableColumns.Name,
// 		),
// 		qm.InnerJoin("gachas on "+db.GachaItemTableColumns.GachaID+" = gachas.id"),
// 		qm.InnerJoin("items on "+db.GachaItemTableColumns.ItemID+" = items.id"),
// 	)

// 	var gachaItems []domain.GachaItem

// 	if err := q.Bind(ctx, g.exec, &gachaItems); err != nil {
// 		return nil, err
// 	}

// 	return gachaItems, nil
// }
