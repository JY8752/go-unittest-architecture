//go:build e2e

package handler_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/JY8752/go-unittest-architecture/controller"
	"github.com/JY8752/go-unittest-architecture/infrastructure/handler"
	"github.com/JY8752/go-unittest-architecture/infrastructure/repository"
	mock_api "github.com/JY8752/go-unittest-architecture/mocks/api"
	mock_domain "github.com/JY8752/go-unittest-architecture/mocks/domain"
	"github.com/JY8752/go-unittest-architecture/test"
	"github.com/labstack/echo/v4"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const (
	goldenDir = "../../testdata/golden/"
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

func TestDraw(t *testing.T) {
	// Arrange
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

	e := echo.New()

	handler.NewGacha(e, controller).Register()

	gachaId := 1
	w := httptest.NewRecorder()
	r := httptest.NewRequest(echo.POST, fmt.Sprintf("/gacha/%d/draw", gachaId), nil)

	g := goldie.New(t, goldie.WithFixtureDir(goldenDir))

	// Act
	response := handleAndIndentResponse(t, e, w, r)

	// Assertion
	assert.Equal(t, http.StatusOK, w.Code)
	g.Assert(t, t.Name(), response)
}

func handleAndIndentResponse(t *testing.T, e *echo.Echo, w *httptest.ResponseRecorder, r *http.Request) []byte {
	t.Helper()
	e.ServeHTTP(w, r)
	var buf bytes.Buffer
	err := json.Indent(&buf, w.Body.Bytes(), "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}
