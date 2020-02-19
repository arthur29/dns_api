package bind

import (
	"fmt"
	"os"

	"github.com/wpalmer/gozone"
)

var zonefile = "/etc/bind/zone/myzone.com.zone"

func ReadZoneFile() (array []gozone.Record) {
	stream, _ := os.Open(zonefile)

	scanner := gozone.NewScanner(stream)
	var record gozone.Record

	for {
		err := scanner.Next(&record)
		if err != nil {
			break
		}
		array = append(array, record)
	}
	for _, entry := range array {
		fmt.Printf("%v\n", entry)
	}

	return array
}
