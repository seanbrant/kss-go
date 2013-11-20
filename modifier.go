package kss

import (
	"strings"
)

type Modifier struct {
	Name        string
	Description string
}

func (m *Modifier) ClassName() string {
	name := strings.Replace(m.Name, ".", " ", -1)
	name = strings.Replace(name, ":", " pseudo-class-", -1)
	return strings.TrimSpace(name)
}

func NewModifier(name string, description string) *Modifier {
	return &Modifier{
		Name:        name,
		Description: description,
	}
}
