//go:build integration

package controller_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/JY8752/go-unittest-architecture/controller"
	"github.com/JY8752/go-unittest-architecture/domain"
	"github.com/JY8752/go-unittest-architecture/infrastructure/repository"
	mock_domain "github.com/JY8752/go-unittest-architecture/mocks/domain"
	"github.com/JY8752/go-unittest-architecture/test"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var db *sql.DB

func TestMain(m *testing.M) {
	container := test.RunMySQLContainer()
	db = container.DB

	if err := test.Migrate(db); err != nil {
		container.Close()
		log.Fatal(err)
	}

	code := m.Run()

	container.Close()
	os.Exit(code)
}

// seed 1000 total weight 71 -> random 67 なのでitem9が抽選されるはず
func TestDraw(t *testing.T) {
	// Arange
	gachaRep := repository.NewGacha(db)
	itemRep := repository.NewItem(db)

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	sg := mock_domain.NewMockSeedGenerator(ctrl)
	sg.EXPECT().New().Return(int64(1000))

	controller := controller.NewGacha(gachaRep, itemRep, sg)

	gachaId := domain.NewGachaId(1)
	expected := domain.NewItem(
		domain.NewItemId(9),
		"item9",
		domain.R,
	)

	// Act
	act, err := controller.Draw(context.Background(), gachaId)
	if err != nil {
		t.Fatal(err)
	}

	// Assertion
	assert.Equal(t, expected, act)
}
