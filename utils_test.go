package contentscraper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceFormatter(t *testing.T) {
	// replacing additional spaces
	must := `Should return normalized text string`
	test := []string{
		"     Should",
		"    return",
		"   normalized   ",
		" text",
		" string",
	}

	result := sliceFormatter(test)
	assert.Equal(t, must, result, "Myst be equal")

}

func TestReplTxt(t *testing.T) {
	must := `Nice string to process. And continue...`
	test := `Nice\u00a0string to ""process.""And continue...`
	result := replTxt(test)
	assert.Equal(t, must, result, "Must be equal")
}
