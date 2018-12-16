package iracing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUglify(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal("%26%23106%3B%26%23111%3B%26%23104%3B%26%23110%3B%26%2332%3B%26%23119%3B%26%2397%3B%26%23121%3B%26%23110%3B%26%23101%3B", uglifyString("john wayne"))
}
