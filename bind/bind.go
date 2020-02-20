package bind

import (
	"io"
	"os"

	"github.com/wpalmer/gozone"
)

var zonefile = "/etc/bind/zone/myzone.com.zone"

type Bind interface {
	ReadZoneFile() ([]gozone.Record, error)
}

func ReadZoneFile() (arrayGozoneRecords []gozone.Record, err error) {
	stream, err := os.Open(zonefile)

	if err != nil {
		return nil, err
	}

	return parseZoneFile(stream), nil
}

func parseZoneFile(reader io.Reader) (arrayGozoneRecords []gozone.Record) {
	scanner := gozone.NewScanner(reader)
	var record gozone.Record

	for {
		err := scanner.Next(&record)

		if err != nil {
			break
		}
		arrayGozoneRecords = append(arrayGozoneRecords, record)
	}
	return arrayGozoneRecords
}
