package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	QUOTE = Atom{"quote"}
	ATOM  = Atom{"atom"}
	EQ    = Atom{"eq"}
	CAR   = Atom{"car"}
	CDR   = Atom{"cdr"}
	CONS  = Atom{"cons"}
	COND  = Atom{"cond"}
)

var (
	ERROR_EOF = `End of file during parsing`
	ERROR_DOT = `Invalid read syntax: ". in wrong context"`
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
	ch, _, err := ReadChar(in)
	chstr := string(ch)
	if err == io.EOF {
		panic(ERROR_EOF)
	}

	buffer := make([]rune, 128)
	var i int
	for i = 0; !unicode.IsSpace(ch) && err == nil && chstr != "(" && chstr != ")"; i++ {
		buffer[i] = ch
		ch, _, err = in.ReadRune()
		chstr = string(ch)
	}
	if err != nil && err != io.EOF {
		panic(err)
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
		list := ReadList(in)
		if list == NIL || fn_cdr(list) != NIL {
			panic(ERROR_DOT)
		}
		return fn_car(list)
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
		return fn_cons(Atom{"quote"}, fn_cons(ReadSexp(in), NIL))
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

func fn_cond(ss ...Sexper) Sexper {
	for _, list := range ss {
		if fn_car(list) != NIL {
			return fn_car(fn_cdr(list))
		}
	}
	return NIL
}

type Evaler interface {
	Get(Sexper) Sexper
	Set(Sexper, Sexper) Evaler
	Eval(Sexper) Sexper
}

// List is an Evaler
func (this List) Get(key Sexper) Sexper {
	if key == NIL || key == TEE {
		return key
	}

	for list := Sexper(this); list != NIL; list = fn_cdr(list) {
		scar := fn_car(list)
		if fn_eq(fn_car(scar), key) == TEE {
			return fn_cdr(scar)
		}
	}
	return NIL
}

func (this List) Set(key, val Sexper) List {
	return fn_cons(fn_cons(key, val), this).(List)
}

func (this List) Eval(sexp Sexper) Sexper {
	switch {
	case fn_atom(sexp) == TEE:
		return this.Get(sexp)
	case fn_atom(fn_car(sexp)) == TEE:
		fn := fn_car(sexp)
		switch {
		case fn_eq(fn, QUOTE) == TEE:
			return fn_car(fn_cdr(sexp))
		case fn_eq(fn, ATOM) == TEE:
			return fn_atom(this.Eval(fn_car(fn_cdr(sexp))))
		case fn_eq(fn, EQ) == TEE:
			scaddr := fn_car(fn_cdr(fn_cdr(sexp)))
			return fn_eq(this.Eval(fn_car(fn_cdr(sexp))), this.Eval(scaddr))
		case fn_eq(fn, CAR) == TEE:
			return fn_car(this.Eval(fn_car(fn_cdr(sexp))))
		case fn_eq(fn, CDR) == TEE:
			return fn_cdr(this.Eval(fn_car(fn_cdr(sexp))))
		case fn_eq(fn, CONS) == TEE:
			scaddr := fn_car(fn_cdr(fn_cdr(sexp)))
			return fn_cons(this.Eval(fn_car(fn_cdr(sexp))), this.Eval(scaddr))
		case fn_eq(fn, COND) == TEE:
			for list := fn_cdr(sexp); list != NIL; list = fn_cdr(list) {
				cond := fn_car(list)
				if this.Eval(fn_car(cond)) != NIL {
					return this.Eval(fn_car(fn_cdr(cond)))
				}
			}
			return NIL
		default:
			return NIL
		}
	default:
		return NIL
	}
}

var GlobalEnv = fn_cons(fn_cons(Atom{"os"}, Atom{"mac"}), NIL).(List)

func InitEnv() List {
	GlobalEnv = GlobalEnv.Set(Atom{"os"}, Atom{"mac"})
	GlobalEnv = GlobalEnv.Set(Atom{"who"}, Atom{"siteshen"})
	GlobalEnv = GlobalEnv.Set(Atom{"editor"}, Atom{"emacs"})
	return GlobalEnv
}

func main() {
	env := InitEnv()
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sexp := ReadSexp(in)
		fmt.Println("    read:", sexp)
		eval := env.Eval(sexp)
		fmt.Println("    eval:", eval)
	}
}
