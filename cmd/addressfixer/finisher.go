package addressfixer

import (
	"bytes"
	"log"
	"strings"

	"github.com/salsalabs/godig"
)

//Finish accepts supporter records at the end of the processing chain.
//The records are written back to the server.
func Finish(t *godig.Table, c chan []Supporter, live bool) {
	var count int32
	count = 0
	for a := range c {
		b := bytes.NewBufferString("")
		for _, s := range a {
			// log.Printf("Finish: %+v\n", s)
			p := []string{
				"",
				"object=supporter",
				"key=" + s.Key,
				"Street=" + s.Street,
				"Street_2=" + s.Street2,
				"City=" + s.City,
				"State=" + s.State,
				"Zip=" + s.Zip,
				"Country=" + s.Country}
			x := strings.Join(p, "&")
			x = strings.Replace(x, " ", "%20", -1)
			_, err := b.WriteString(x)
			if err != nil {
				panic(err)
			}
		}
		count = count + int32(len(a))
		if live {
			body, err := t.SaveBulk(b.String())
			if err != nil {
				panic(err)
			}
			log.Printf("Finish:  saved %v\n", len(a))
			log.Printf("Finish:  /save returned %v\n", string(body))
		}
	}
	log.Printf("Finish:  done, count %v", count)
}