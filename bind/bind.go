package bind

import (
	"io"
	"os"
	"strconv"

	"github.com/wpalmer/gozone"
)

var zonefile = "/etc/bind/zone/myzone.com.zone"

type Record struct {
	DomainName string
	TimeToLive string
	Class      string
	Type       string
	Data       []string
	Comment    string
}

type Bind interface {
	ReadZoneFile() ([]Record, error)
}

func readZoneFile() ([]Record, error) {
	stream, err := os.Open(zonefile)

	if err != nil {
		return nil, err
	}

	return parseZoneFile(stream), nil
}

func parseZoneFile(reader io.Reader) (array []Record) {
	scanner := gozone.NewScanner(reader)
	var record gozone.Record

	for {
		err := scanner.Next(&record)

		if err != nil {
			break
		}
		array = append(array, castRecord(record))
	}
	return array
}

func castRecord(gozoneRecord gozone.Record) Record {
	var record Record

	record.DomainName = gozoneRecord.DomainName
	record.TimeToLive = strconv.FormatInt(gozoneRecord.TimeToLive, 10)
	record.Class = gozoneRecord.Class.String()
	record.Type = gozoneRecord.Type.String()
	record.Data = gozoneRecord.Data
	record.Comment = gozoneRecord.Comment

	return record
}
