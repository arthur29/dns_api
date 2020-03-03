package bind

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/wpalmer/gozone"
)

const zonefile = "/etc/bind/zone/myzone.com.zone"

type Record struct {
	DomainName string
	TimeToLive string
	Class      string
	Type       string
	Data       string
	Comment    string
}

type Bind struct {
	bindBehavior BindBehavior
	ArrayRecords []Record
}

type BindBehavior interface {
	openFileStream() (io.Reader, error)
}

type bindImp struct{}

func InitializeBind() Bind {
	var bind Bind
	bind.bindBehavior = new(bindImp)

	return bind
}

func (bind *Bind) GetZoneRecords() error {
	if bind.ArrayRecords == nil {
		stream, err := bind.bindBehavior.openFileStream()

		if err == nil {
			bind.ArrayRecords = parseZoneFile(stream)

			return nil
		}

		return err
	}

	return nil
}

func (bind *bindImp) openFileStream() (io.Reader, error) {
	stream, err := os.Open(zonefile)

	if err != nil {
		return nil, err
	}

	return stream, nil
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
	record.Data = strings.Join(gozoneRecord.Data, " ")
	record.Comment = gozoneRecord.Comment

	return record
}
