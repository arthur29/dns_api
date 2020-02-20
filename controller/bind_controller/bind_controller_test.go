package bind_controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetListStatusOkWhenBindReturnsNoError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	rec := httptest.NewRecorder()
	con := e.NewContext(req, rec)

	if assert.NoError(t, ListIndex(con)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestGetListReturnsJSONWhenBindReturnsNoError(t *testing.T) {

}

func TestGetListReturnsAListOfZoneRecordsWhenBindReturnsNoError(t *testing.T) {

}

func TestGetListReturnsStringMessageWhenBindReturnsInError(t *testing.T) {

}

func TestGetListReturnsInternalServerErrorWhenBindReturnsInError(t *testing.T) {

}
