package shortid

import (
	"log"

	"github.com/teris-io/shortid"
)

var sid, err = shortid.New(1, shortid.DefaultABC, 2342)

func init() {
	if err != nil {
		log.Fatalf("error initializing shortid: %v", err)
	}
}

// Generate generates a 9 characters long unique id that is expected
// to have no collision for 34 years (1/1/2016-1/1/2050).
// see: https://github.com/teris-io/shortid
func Generate() (string, error) {
	return sid.Generate()
}
