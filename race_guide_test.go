package iracing

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIRString(t *testing.T) {
	asrt := assert.New(t)
	var is struct{ Foo String }
	asrt.NoError(json.Unmarshal([]byte(`{"foo": "N%C3%BCrburgring+Nordschleife"}`), &is))
	asrt.Equal("Nürburgring Nordschleife", string(is.Foo))
}
