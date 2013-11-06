package kss

import (
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

	NewSection(comment, "example.css")
}
