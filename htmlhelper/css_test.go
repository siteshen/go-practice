package htmlhelper

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	sel1 := SimpleSelector{"ul", "li", "a"}
	sel2 := SimpleSelector{"tr", "td", "a"}
	sel := Selector{sel1, sel2}

	role1 := Role{"width", "100px"}
	role2 := Role{"height", "20px"}
	css := CSS{
		Selector: sel,
		Roles:    []Role{role1, role2},
	}

	fmt.Println(sel1, role1)
	fmt.Println(sel2, role2)
	fmt.Println(css.ShortString())
	fmt.Println(css)
}

func TestNew(t *testing.T) {
	ss1 := newSimpleSelector("ul li a")
	ss2 := newSimpleSelector("tr td a")
	fmt.Println(ss1)
	fmt.Println(ss2)

	selector := NewSelector("ul li a,tr td a")
	fmt.Println(selector)
	fmt.Println(Selector{*ss1, *ss2})
}
