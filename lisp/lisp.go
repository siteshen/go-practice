package lisp

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// Atom, List, Sexper
type String struct{ val string }
type Number struct{ val int }
type Atom struct{ name string }
type List struct{ car, cdr Sexper }
type Sexper interface {
	Numberp() bool
	Stringp() bool
	Atomp() bool
	Listp() bool
	String() string
}

// String is a Sexper
func (this String) Numberp() bool  { return false }
func (this String) Stringp() bool  { return true }
func (this String) Atomp() bool    { return true }
func (this String) Listp() bool    { return false }
func (this String) String() string { return fmt.Sprintf("%q", this.val) }

// Number is a Sexper
func (this Number) Numberp() bool  { return true }
func (this Number) Stringp() bool  { return false }
func (this Number) Atomp() bool    { return true }
func (this Number) Listp() bool    { return false }
func (this Number) String() string { return fmt.Sprintf("%d", this.val) }

// Atom is a Sexper
func (this Atom) Numberp() bool  { return false }
func (this Atom) Stringp() bool  { return false }
func (this Atom) Atomp() bool    { return true }
func (this Atom) Listp() bool    { return false }
func (this Atom) String() string { return this.name }

// List is a Sexper
func (this List) Numberp() bool { return false }
func (this List) Stringp() bool { return false }
func (this List) Atomp() bool   { return false }
func (this List) Listp() bool   { return true }
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

var (
	ADD = Atom{"+"}
	SUB = Atom{"-"}
	MUL = Atom{"*"}
	MOD = Atom{"%"}
	DIV = Atom{"/"}
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

func ReadNumber(in io.RuneScanner) Sexper {
	ch, _, err := ReadChar(in)
	chstr := string(ch)
	if err != nil {
		panic(err)
	}

	result := 0
	for unicode.IsDigit(ch) && err == nil && chstr != "(" && chstr != ")" {
		number, _ := strconv.Atoi(string(ch))
		result = result*10 + number
		ch, _, err = in.ReadRune()
		chstr = string(ch)
	}
	if err != nil && err != io.EOF {
		panic(err)
	}
	if chstr == "(" || chstr == ")" {
		in.UnreadRune()
	}
	return Number{result}
}

func ReadString(in io.RuneScanner) Sexper {
	ch, _, err := ReadChar(in)
	chstr := string(ch)
	if err != nil {
		panic(err)
	}

	buffer := make([]rune, 128)
	var i int
	for i = 0; err == nil && chstr != "\"" && chstr != "(" && chstr != ")"; i++ {
		buffer[i] = ch
		ch, _, err = in.ReadRune()
		chstr = string(ch)
	}
	if err != nil && err != io.EOF {
		panic(err)
	}

	if err == io.EOF && chstr != "\"" {
		panic("End of file during parsing")
	}
	return String{string(buffer[0:i])}
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
	case "\"":
		return ReadString(in)
	case "(":
		return ReadList(in)
	case "'":
		return fn_cons(Atom{"quote"}, fn_cons(ReadSexp(in), NIL))
	case ")":
		panic("unexcepted ')' found")
	default:
		in.UnreadRune()
		if unicode.IsDigit(ch) {
			return ReadNumber(in)
		} else {
			return ReadAtom(in)
		}
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

// arithmetic
func fn_add(x, y Sexper) Sexper {
	return Number{x.(Number).val + y.(Number).val}
}

func fn_sub(x, y Sexper) Sexper {
	return Number{x.(Number).val - y.(Number).val}
}

func fn_mul(x, y Sexper) Sexper {
	return Number{x.(Number).val * y.(Number).val}
}

func fn_div(x, y Sexper) Sexper {
	return Number{x.(Number).val / y.(Number).val}
}

func fn_mod(x, y Sexper) Sexper {
	return Number{x.(Number).val % y.(Number).val}
}

// Evaler
type Evaler interface {
	Get(Sexper) Sexper
	Set(Sexper, Sexper) Evaler
	Eval(Sexper) Sexper
}

// List is an Evaler
func (this List) Get(key Sexper) Sexper {
	if key == NIL || key == TEE || key.Stringp() || key.Numberp() {
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
		// add, sub, mul, div, mod
		case fn_eq(fn, ADD) == TEE:
			scaddr := fn_car(fn_cdr(fn_cdr(sexp)))
			return fn_add(this.Eval(fn_car(fn_cdr(sexp))), this.Eval(scaddr))
		case fn_eq(fn, SUB) == TEE:
			scaddr := fn_car(fn_cdr(fn_cdr(sexp)))
			return fn_sub(this.Eval(fn_car(fn_cdr(sexp))), this.Eval(scaddr))
		case fn_eq(fn, MUL) == TEE:
			scaddr := fn_car(fn_cdr(fn_cdr(sexp)))
			return fn_mul(this.Eval(fn_car(fn_cdr(sexp))), this.Eval(scaddr))
		case fn_eq(fn, DIV) == TEE:
			scaddr := fn_car(fn_cdr(fn_cdr(sexp)))
			return fn_div(this.Eval(fn_car(fn_cdr(sexp))), this.Eval(scaddr))
		case fn_eq(fn, MOD) == TEE:
			scaddr := fn_car(fn_cdr(fn_cdr(sexp)))
			return fn_mod(this.Eval(fn_car(fn_cdr(sexp))), this.Eval(scaddr))

		// quote, atom, eq, car, cdr, cons, cond
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
