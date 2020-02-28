package bind_controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/wpalmer/gozone"
)

type MockedBindController struct {
	returnValue []gozone.Record
	err         error
}

func (bindController *MockedBindController) searchRecords() ([]gozone.Record, error) {
	return bindController.returnValue, bindController.err
}

func initializeMockedBindController(returnValueArgument []gozone.Record, errArgument error) BindControllerDI {
	var bindControllerDI BindControllerDI
	bindControllerDI.bindControllerBehavior = &MockedBindController{returnValue: returnValueArgument, err: errArgument}
	return bindControllerDI
}

func TestGetListReturnsJSONWithStatusOkWhenBindReturnsNoError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	resp := httptest.NewRecorder()
	con := e.NewContext(req, resp)

	bindController := initializeMockedBindController(nil, nil)

	if assert.NoError(t, bindController.ListIndex(con)) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, resp.Header()["Content-Type"][0], "application/json")
	}
}

func TestGetListReturnsAListOfZoneRecordsWhenBindReturnsNoError(t *testing.T) {

}

func TestGetListReturnsMessageWhenBindReturnsInError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	resp := httptest.NewRecorder()
	con := e.NewContext(req, resp)

	bindController := initializeMockedBindController(nil, errors.New("Error on test"))

	if assert.NoError(t, bindController.ListIndex(con)) {
		assert.Equal(t, 500, resp.Code)
		assert.Equal(t, "Error on read dns zone file", resp.Body.String())
	}

}

func TestGetListReturnsInternalServerErrorWhenBindReturnsInError(t *testing.T) {

}
