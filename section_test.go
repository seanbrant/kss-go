package kss

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSection(t *testing.T) {
	comment := strings.TrimSpace(`
# Form Button

Your standard form button.

:hover    - Highlights when hovering.
:disabled - Dims the button when disabled.
.primary  - Indicates button is the primary action.
.smaller  - A smaller button

Styleguide 2.1.1.
    `)

	section := NewSection(comment, "example.css")
	assert.Equal(t, section.Reference, "2.1.1")
}

func TestSectionWithStringReference(t *testing.T) {
	comment := strings.TrimSpace(`
Styleguide Buttons.1.1.
    `)

	section := NewSection(comment, "example.css")
	assert.Equal(t, section.Reference, "Buttons.1.1")
}
