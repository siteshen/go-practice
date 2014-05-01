package htmlhelper

import (
	"fmt"
	"strings"
)

// ul li a
type SimpleSelector []string

// ul li a, tr td a
type Selector []SimpleSelector

// height: 20px;
type Role struct{ Name, Value string }

// .container { max-width: 1200px; }
type CSS struct {
	Selector Selector
	Roles    []Role
}

// Stringer
func (self SimpleSelector) String() string {
	return strings.Join(self, " ")
}

func (self Selector) String() string {
	// Selectors part
	var sels = []string{}
	for _, simple := range self {
		sels = append(sels, simple.String())
	}

	return strings.Join(sels, ", ")
}

func (self Role) String() string {
	return fmt.Sprintf("%s: %s;", self.Name, self.Value)
}

func (self CSS) string(pretty bool) string {
	// pretty format?
	var space, newline string
	if pretty {
		space, newline = " ", "\n"
	}

	var roles []string
	for _, elem := range self.Roles {
		roles = append(roles, strings.Repeat(space, 4)+elem.String())
	}

	return fmt.Sprintf("%s {%s%s%s}",
		self.Selector.String(), newline,
		strings.Join(roles, newline), newline)
}

func (self CSS) String() string      { return self.string(true) }
func (self CSS) ShortString() string { return self.string(false) }

func newSimpleSelector(in string) *SimpleSelector {
	var ss SimpleSelector
	for _, word := range strings.Split(in, " ") {
		trimed := strings.TrimSpace(word)
		if trimed != "" {
			ss = append(ss, word)
		}
	}

	return &ss
}

func NewSelector(in string) *Selector {
	var selector Selector
	for _, line := range strings.Split(in, ",") {
		ss := newSimpleSelector(line)
		if ss != nil {
			selector = append(selector, *ss)
		}
	}

	return &selector
}
