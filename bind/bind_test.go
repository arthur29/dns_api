package bind

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BindMocked struct {
	returnStream string
	err          error
}

func (bindMocked *BindMocked) openFileStream() (io.Reader, error) {
	return strings.NewReader(bindMocked.returnStream), bindMocked.err
}

func initializeBindMocked(returnValue string, err error) Bind {
	var bind Bind
	var bindMocked = BindMocked{returnStream: returnValue, err: err}

	bind.bindBehavior = &bindMocked

	return bind
}

func TestWhenArrayRecordsIsNilShouldPopulateItWithDataFromFile(t *testing.T) {
	zoneInfo := `$ORIGIN test.
				 $TTL 1
				 @      IN A  192.168.11.79`

	bind := initializeBindMocked(zoneInfo, nil)

	var arrayRecords = []Record{
		Record{
			DomainName: "test.",
			TimeToLive: "1",
			Type:       "A",
			Class:      "IN",
			Data:       []string{"192.168.11.79"},
			Comment:    "",
		},
	}

	bind.GetZoneRecords()

	assert.Equal(t, arrayRecords, bind.ArrayRecords)
}

func TestWhenArrayRecordsIsNilShouldParseZoneFileAndReturnASliceOfRecords(t *testing.T) {
	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )
				 @      IN NS .
				 @      IN A  192.168.11.79`
	bind := initializeBindMocked(zoneInfo, nil)

	bind.GetZoneRecords()
	array := bind.ArrayRecords

	assert.Equal(t, reflect.Slice, reflect.TypeOf(array).Kind())

	var expected Record
	assert.Equal(t, reflect.TypeOf(expected), reflect.TypeOf(array[0]))
}

func TestWhenArrayRecordsIsNilShouldParseZoneFileAndReturnClassTypeAndTTLAsStringValues(t *testing.T) {
	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )`
	bind := initializeBindMocked(zoneInfo, nil)

	bind.GetZoneRecords()
	array := bind.ArrayRecords

	record := array[0]

	assert.Equal(t, record.Class, "IN")
	assert.Equal(t, record.Type, "SOA")
	assert.Equal(t, record.TimeToLive, "86400")
}

func TestWhenArrayRecordsIsNotNilShouldIgnoreStreamAndReturnThatValue(t *testing.T) {
	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )
				 @      IN NS .
				 @      IN A  192.168.11.79`

	var arrayRecords = []Record{
		Record{
			DomainName: "test.",
			TimeToLive: "1",
			Type:       "A",
			Class:      "IN",
			Data:       []string{"192.168.11.79"},
			Comment:    "",
		},
	}

	bind := initializeBindMocked(zoneInfo, nil)

	bind.ArrayRecords = arrayRecords

	bind.GetZoneRecords()
	array := bind.ArrayRecords

	assert.Equal(t, array, arrayRecords)
}
