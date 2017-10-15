package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlagMatch(t *testing.T) {
	input := "someinput --debug"
	output, debug := Flag(input, "--debug")
	assert.Equal(t, output, "someinput", "they should be equal")
	assert.Equal(t, debug, true, "they should be equal")
}

func TestFlagNoMatch(t *testing.T) {
	input := "someinput --debug"
	output, debug := Flag(input, "--debunk")
	assert.Equal(t, output, input, "they should be equal")
	assert.Equal(t, debug, false, "they should be equal")
}

func TestFlagSpacing(t *testing.T) {
	input := "someinput--debug"
	output, debug := Flag(input, "--debug")
	assert.Equal(t, output, input, "they should be equal")
	assert.Equal(t, debug, false, "they should be equal")
}

func TestFlagSpacePadding(t *testing.T) {
	input := "   someinput   --debug   "
	output, debug := Flag(input, "--debug")
	assert.Equal(t, output, "someinput", "they should be equal")
	assert.Equal(t, debug, true, "they should be equal")
}
