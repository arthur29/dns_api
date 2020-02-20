package bind

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wpalmer/gozone"
)

func TestParseZoneFileShouldReturnASliceOfGozoneRecords(t *testing.T) {
	zoneInfo := `$ORIGIN myzone.com.
				 $TTL 86400
				 @ IN SOA myzone.com arthur29 ( 2020021600 3600 900 604800 86400 )
				 @      IN NS .
				 @      IN A  192.168.11.79`
	scanner := strings.NewReader(zoneInfo)
	arrayGozoneRecords := parseZoneFile(scanner)

	assert.Equal(t, reflect.Slice, reflect.TypeOf(arrayGozoneRecords).Kind())

	var expected gozone.Record
	assert.Equal(t, reflect.TypeOf(expected), reflect.TypeOf(arrayGozoneRecords[0]))
}
