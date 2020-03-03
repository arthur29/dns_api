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

func TestPopulateArrayRecordWithDataFromFile(t *testing.T) {
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
			Data:       "192.168.11.79",
			Comment:    "",
		},
	}

	bind.LoadZoneRecords()

	assert.Equal(t, arrayRecords, bind.ArrayRecords)
}

func TestParseZoneFileAndReturnASliceOfRecords(t *testing.T) {
	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )
				 @      IN NS .
				 @      IN A  192.168.11.79`
	bind := initializeBindMocked(zoneInfo, nil)

	bind.LoadZoneRecords()
	array := bind.ArrayRecords

	assert.Equal(t, reflect.Slice, reflect.TypeOf(array).Kind())

	var expected Record
	assert.Equal(t, reflect.TypeOf(expected), reflect.TypeOf(array[0]))
}

func TestParseZoneFileAndReturnClassTypeTTLAndDataAsStringValues(t *testing.T) {
	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )`
	bind := initializeBindMocked(zoneInfo, nil)

	bind.LoadZoneRecords()
	array := bind.ArrayRecords

	record := array[0]

	assert.Equal(t, "IN", record.Class)
	assert.Equal(t, "SOA", record.Type)
	assert.Equal(t, "86400", record.TimeToLive)
	assert.Equal(t, "myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )", record.Data)
}
