package main

import (
	// "bufio"
	"fmt"
	"io"
	// "os"
	"strings"
	"unicode"
)

// Atom, List, Sexper
type Atom struct{ name string }
type List struct{ car, cdr Sexper }
type Sexper interface {
	Atomp() bool
	Listp() bool
	String() string
}

// Atom is a Sexper
func (this Atom) Atomp() bool    { return true }
func (this Atom) Listp() bool    { return false }
func (this Atom) String() string { return this.name }

// List is a Sexper
func (this List) Atomp() bool { return false }
func (this List) Listp() bool { return true }
func (this List) String() string {
	scar, scdr := this.car.String(), this.cdr.String()
	switch {
	case scdr[0] == '(':
		return strings.Join([]string{"(", scar, " ", scdr[1 : len(scdr)-1], ")"}, "")
	case scdr == "nil":
		return strings.Join([]string{"(", scar, ")"}, "")
	default:
		return strings.Join([]string{"(", scar, " . ", scdr, ")"}, "")
	}
}

var (
	NIL = Atom{"nil"}
	TEE = Atom{"t"}
)

// like in.ReadRune() but ignore all leading whitespace.
func ReadChar(in io.RuneScanner) (r rune, size int, err error) {
	r, size, err = in.ReadRune()
	for unicode.IsSpace(r) && err == nil {
		r, size, err = in.ReadRune()
	}
	return
}

// sexp reader
func ReadAtom(in io.RuneScanner) Sexper {
	ch, _, _ := ReadChar(in)
	chstr := string(ch)

	buffer := make([]rune, 128)
	var i int
	for i = 0; !unicode.IsSpace(ch) && chstr != "(" && chstr != ")"; i++ {
		buffer[i] = ch
		ch, _, _ = in.ReadRune()
		chstr = string(ch)
	}
	if chstr == "(" || chstr == ")" {
		in.UnreadRune()
	}
	return Atom{string(buffer[0:i])}
}

func ReadList(in io.RuneScanner) Sexper {
	ch, _, _ := ReadChar(in)
	chstr := string(ch)

	switch chstr {
	case ")":
		return NIL
	case ".":
		return ReadAtom(in)
	default:
		in.UnreadRune()
		scar := ReadSexp(in)
		scdr := ReadList(in)
		return List{scar, scdr}
	}
}

func ReadSexp(in io.RuneScanner) Sexper {
	ch, _, _ := ReadChar(in)

	switch string(ch) {
	case "(":
		return ReadList(in)
	case "'":
		return fn_list(Atom{"quote"}, ReadSexp(in))
	case ")":
		panic("unexcepted ')' found")
	default:
		in.UnreadRune()
		return ReadAtom(in)
	}
}

func ReadFrom(str string) Sexper {
	return ReadSexp(strings.NewReader(str))
}

// seven basic functions
func fn_quote(sexp Sexper) Sexper {
	return sexp
}

func fn_atom(sexp Sexper) Sexper {
	if sexp.Atomp() {
		return TEE
	} else {
		return NIL
	}
}

func fn_eq(s1, s2 Sexper) Sexper {
	if s1 == s2 {
		return TEE
	} else {
		return NIL
	}
}

func fn_car(sexp Sexper) Sexper {
	return sexp.(List).car
}

func fn_cdr(sexp Sexper) Sexper {
	return sexp.(List).cdr
}

func fn_cons(s1, s2 Sexper) Sexper {
	return List{s1, s2}
}

func fn_list(s1, s2 Sexper) Sexper {
	return List{s1, List{s2, NIL}}
}

func fn_cond(ss ...Sexper) Sexper {
	for _, list := range ss {
		if fn_car(list) != NIL {
			return fn_car(fn_cdr(list))
		}
	}
	return NIL
}

func main() {
	atomn := NIL
	atom1 := Atom{"atom1"}
	atom2 := Atom{"atom2"}
	atom3 := Atom{"atom3"}
	list1 := List{Sexper(&atom1), Sexper(&atom2)}
	list2 := List{Sexper(&atom3), Sexper(&atomn)}
	list3 := List{Sexper(&list1), Sexper(&list2)}
	list4 := List{Sexper(&list2), Sexper(&list3)}
	fmt.Println(list1, list2, list3, list4)
	hello_world := ReadFrom("'(hello . world)")
	fmt.Println(hello_world, fn_car(hello_world), fn_cdr(hello_world))
	fmt.Println(fn_cons(Atom{"hello"}, fn_cons(Atom{"emacs"}, NIL)))

	// env := List{fn_cons(Atom{"hello"}, Atom{"world"}), NIL}
	// fmt.Println(env.Get(Atom{"hello"}))
	// fmt.Println(env.Set(Atom{"os"}, Atom{"mac"}))
	// fmt.Println(env.Set(Atom{"os"}, Atom{"mac"}).Get(Atom{"os"}))
}
