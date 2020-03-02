package bind_controller

import (
	"dns_api/bind"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockedBindController struct {
	returnValue []bind.Record
	err         error
}

func (bindController *MockedBindController) searchRecords() ([]bind.Record, error) {
	return bindController.returnValue, bindController.err
}

func initializeMockedBindController(returnValueArgument []bind.Record, errArgument error) BindController {
	var bindController BindController
	bindController.bindControllerBehavior = &MockedBindController{returnValue: returnValueArgument, err: errArgument}
	return bindController
}

func TestGetIndextReturnsJSONWithStatusOkWhenBindReturnsNoError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	resp := httptest.NewRecorder()
	con := e.NewContext(req, resp)

	bindController := initializeMockedBindController(nil, nil)

	if assert.NoError(t, bindController.Index(con)) {
		assert.Equal(t, 200, resp.Code)
		assert.Contains(t, resp.Header()["Content-Type"][0], "application/json")
	}
}

func TestGetIdexReturnsAListOfBindRecordsWhenBindReturnsNoError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	resp := httptest.NewRecorder()
	con := e.NewContext(req, resp)

	records := []bind.Record{bind.Record{DomainName: "myzone.com.", TimeToLive: "86400", Class: "IN", Type: "SOA", Data: []string{"test"}, Comment: "Comment"}}

	bindController := initializeMockedBindController(records, nil)

	if assert.NoError(t, bindController.Index(con)) {
		var recordArray []bind.Record
		json.Unmarshal(resp.Body.Bytes(), &recordArray)
		assert.Equal(t, records, recordArray)
	}
}

func TestGetReturns500AndMessageWhenBindReturnsInError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	resp := httptest.NewRecorder()
	con := e.NewContext(req, resp)

	bindController := initializeMockedBindController(nil, errors.New("Error on test"))

	if assert.NoError(t, bindController.Index(con)) {
		assert.Equal(t, 500, resp.Code)
		assert.Equal(t, "Error on read dns zone file", resp.Body.String())
	}

}
