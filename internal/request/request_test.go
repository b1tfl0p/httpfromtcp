package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestLineParse(t *testing.T) {
	assert.Equal(t, "bit flip fl0p", "bit flip fl0p")
}
