package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	var match bool
	match = Match("x", "x")
	assert.Equal(t, match, true, "they should be equal")
	match = Match("x*", "xyz")
	assert.Equal(t, match, true, "they should be equal")
	match = Match("*x", "zyx")
	assert.Equal(t, match, true, "they should be equal")
	match = Match("*y*", "xyz")
	assert.Equal(t, match, true, "they should be equal")
}

func TestNotMatch(t *testing.T) {
	var match bool
	match = Match("", "y")
	assert.Equal(t, match, false, "they should be equal")
	match = Match("x", "y")
	assert.Equal(t, match, false, "they should be equal")
	match = Match("x", "xyz")
	assert.Equal(t, match, false, "they should be equal")
	match = Match("x", "zyx")
	assert.Equal(t, match, false, "they should be equal")
	match = Match("y", "xyz")
	assert.Equal(t, match, false, "they should be equal")
}
