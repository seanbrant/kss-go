package kss

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	scss := Parser("fixtures/scss")
	less := Parser("fixtures/less")
	sass := Parser("fixtures/sass")
	css := Parser("fixtures/css")
	multiple := Parser("fixtures/scss", "fixtures/less")

	// parses kss comments in scss
	assert.Equal(t, scss["2.1.1"].Description, "Your standard form button.")

	// parses kss comments in less
	assert.Equal(t, less["2.1.1"].Description, "Your standard form button.")

	// parses kss multi line comments in sass
	assert.Equal(t, sass["2.1.1"].Description, "Your standard form button.")

	// parses kss single line comments in sass
	assert.Equal(t, sass["2.2.1"].Description, "A button suitable for giving stars to someone.")

	// parses kss comments in css
	assert.Equal(t, css["2.1.1"].Description, "Your standard form button.")

	// parses nested scss documents
	assert.Equal(t, scss["3.0.0"].Description, "Your standard form element.")
	assert.Equal(t, scss["3.0.1"].Description, "Your standard text input box.")

	// parses nested less documents
	assert.Equal(t, less["3.0.0"].Description, "Your standard form element.")
	assert.Equal(t, less["3.0.1"].Description, "Your standard text input box.")

	// parses nested sass documents
	assert.Equal(t, sass["3.0.0"].Description, "Your standard form element.")
	assert.Equal(t, sass["3.0.1"].Description, "Your standard text input box.")

	// parse returns map of sections
	assert.Equal(t, len(css), 2)

	// parse multiple paths
	assert.Equal(t, len(multiple), 6)
}
