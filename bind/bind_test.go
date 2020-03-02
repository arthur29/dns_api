package bind

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseZoneFileShouldReturnASliceOfRecords(t *testing.T) {
	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )
				 @      IN NS .
				 @      IN A  192.168.11.79`
	scanner := strings.NewReader(zoneInfo)
	array := parseZoneFile(scanner)

	assert.Equal(t, reflect.Slice, reflect.TypeOf(array).Kind())

	var expected Record
	assert.Equal(t, reflect.TypeOf(expected), reflect.TypeOf(array[0]))
}

func TestParseZoneFileShouldReturnClassTypeAndTTLAsStringValues(t *testing.T) {

	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )`
	scanner := strings.NewReader(zoneInfo)
	array := parseZoneFile(scanner)
	record := array[0]

	assert.Equal(t, record.Class, "IN")
	assert.Equal(t, record.Type, "SOA")
	assert.Equal(t, record.TimeToLive, "86400")
}
