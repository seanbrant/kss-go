package kss

import (
	"regexp"
	"strings"
)

var classModifier = "."
var pseudoClassModifier = ":"
var modifierDescriptionSeparator = " - "
var referenceStart = "Styleguide"

var referenceRegexp, _ = regexp.Compile(`Styleguide ([\d\.]+)`)
var optionalRegexp, _ = regexp.Compile(`\[(.*)\]\?`)

type Section struct {
	Reference   string
	Description string
	Modifiers   []Modifier
}

func NewSection(comment string, filename string) Section {
	var reference string
	var description string

	modifiers := []Modifier{}
	descriptionLines := []string{}

	lines := strings.Split(comment, "\n")

	for i := range lines {
		line := lines[i]

		if strings.HasPrefix(line, classModifier) || strings.HasPrefix(line, pseudoClassModifier) {
			bits := strings.Split(line, modifierDescriptionSeparator)

			if len(bits) > 1 {
				modifier := strings.TrimSpace(bits[0])
				description := strings.TrimSpace(bits[1])
				modifiers = append(modifiers, NewModifier(modifier, description))
			}
		} else if strings.HasPrefix(line, referenceStart) {
			reference = strings.TrimRight(referenceRegexp.FindStringSubmatch(line)[1], ".")
		} else {
			descriptionLines = append(descriptionLines, line)
		}
	}

	description = strings.TrimSpace(strings.Join(descriptionLines, "\n"))

	return Section{
		Reference:   reference,
		Description: description,
		Modifiers:   modifiers,
	}
}
