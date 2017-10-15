package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemberMatch(t *testing.T) {
	var match bool
	match = Member("x", "x")
	assert.Equal(t, match, true, "they should be equal")
	match = Member("x,y,z", "x")
	assert.Equal(t, match, true, "they should be equal")
	match = Member("x,y,z", "y")
	assert.Equal(t, match, true, "they should be equal")
	match = Member("x,y,z", "z")
	assert.Equal(t, match, true, "they should be equal")
}

func TestMemberWildcard(t *testing.T) {
	var match bool
	match = Member("*", "x")
	assert.Equal(t, match, true, "they should be equal")
	match = Member("*,y,z", "x")
	assert.Equal(t, match, true, "they should be equal")
	match = Member("x,*,z", "y")
	assert.Equal(t, match, true, "they should be equal")
	match = Member("x,y,*", "z")
	assert.Equal(t, match, true, "they should be equal")
}

func TestMemberNoMatch(t *testing.T) {
	var match bool
	match = Member("", "n")
	assert.Equal(t, match, false, "they should be equal")
	match = Member("x", "n")
	assert.Equal(t, match, false, "they should be equal")
	match = Member("x,y,z", "n")
	assert.Equal(t, match, false, "they should be equal")
	match = Member("x,y,z", "n")
	assert.Equal(t, match, false, "they should be equal")
	match = Member("x,y,z", "n")
	assert.Equal(t, match, false, "they should be equal")
}
