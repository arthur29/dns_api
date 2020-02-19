package bind_controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetListStatusOk(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	rec := httptest.NewRecorder()
	con := e.NewContext(req, rec)

	if assert.NoError(ListIndex(con)) {
		assert.Equal(t, 200, rec.Code)
	}
}
