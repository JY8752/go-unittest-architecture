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
	mock_api "github.com/JY8752/go-unittest-architecture/mocks/api"
	mock_domain "github.com/JY8752/go-unittest-architecture/mocks/domain"
	"github.com/JY8752/go-unittest-architecture/test"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const migrationsPath = "../migrations"

var db *sql.DB

func TestMain(m *testing.M) {
	container, err := test.RunMySQLContainer()
	if err != nil {
		container.Close()
		log.Fatal(err)
	}

	db = container.DB

	if err := test.Migrate(db, migrationsPath); err != nil {
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

	p := mock_api.NewMockPayment(ctrl)
	p.EXPECT().Buy(100).Return(nil).Times(1)

	controller := controller.NewGacha(gachaRep, itemRep, sg, p)

	expected := domain.NewItem(
		domain.NewItemId(9),
		"item9",
		domain.R,
	)

	// Act
	act, err := controller.Draw(context.Background(), domain.NewGachaId(1))
	if err != nil {
		t.Fatal(err)
	}

	// Assertion
	assert.Equal(t, expected, act)
}
