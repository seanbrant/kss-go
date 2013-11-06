package kss

import (
	"fmt"
	"strings"
)

type Modifier struct {
	Name        string
	Description string
	ClassName   string
	Example     string
}

func (m *Modifier) AddExample(example string) {
	replacement := fmt.Sprintf(" %s", m.ClassName)
	m.Example = strings.Replace(example, "$modifier_class", replacement, 1)
}

func getClassNameFromName(name string) string {
	name = strings.Replace(name, ".", " ", -1)
	name = strings.Replace(name, ":", " pseudo-class-", -1)
	return strings.TrimSpace(name)
}

func NewModifier(name string, description string) Modifier {
	return Modifier{
		Name:        name,
		Description: description,
		ClassName:   getClassNameFromName(name),
	}
}
