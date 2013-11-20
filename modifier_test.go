package kss

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModifer(t *testing.T) {
	modifier := NewModifier(".callout.extreme:hover", "calls things out")

	// handles pseudo and multiple classes
	assert.Equal(t, modifier.ClassName(), "callout extreme pseudo-class-hover")
}
