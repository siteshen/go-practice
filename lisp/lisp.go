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
type Func struct{ val Callable }
type Atom struct{ name string }
type List struct{ car, cdr Sexper }
type Sexper interface {
	Numberp() bool
	Stringp() bool
	Funcp() bool
	Atomp() bool
	Listp() bool
	String() string
}

type Callable (func(Sexper) Sexper)

// String is a Sexper
func (this String) Numberp() bool  { return false }
func (this String) Stringp() bool  { return true }
func (this String) Funcp() bool    { return false }
func (this String) Atomp() bool    { return true }
func (this String) Listp() bool    { return false }
func (this String) String() string { return fmt.Sprintf("%q", this.val) }

// Number is a Sexper
func (this Number) Numberp() bool  { return true }
func (this Number) Stringp() bool  { return false }
func (this Number) Funcp() bool    { return false }
func (this Number) Atomp() bool    { return true }
func (this Number) Listp() bool    { return false }
func (this Number) String() string { return fmt.Sprintf("%d", this.val) }

// Func is a Sexper
func (this Func) Numberp() bool  { return false }
func (this Func) Stringp() bool  { return false }
func (this Func) Funcp() bool    { return true }
func (this Func) Atomp() bool    { return false }
func (this Func) Listp() bool    { return false }
func (this Func) String() string { return fmt.Sprintf("#Callable", this.val) }

// Atom is a Sexper
func (this Atom) Numberp() bool  { return false }
func (this Atom) Stringp() bool  { return false }
func (this Atom) Funcp() bool    { return false }
func (this Atom) Atomp() bool    { return true }
func (this Atom) Listp() bool    { return false }
func (this Atom) String() string { return this.name }

// List is a Sexper
func (this List) Numberp() bool { return false }
func (this List) Stringp() bool { return false }
func (this List) Funcp() bool   { return false }
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
	for i = 0; err == nil && chstr != "\""; i++ {
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

// Callable
func func_quote(args Sexper) Sexper {
	return fn_car(args)
}

func func_atom(args Sexper) Sexper {
	return fn_atom(GlobalEnv.Eval(fn_car(args)))
}

func func_eq(args Sexper) Sexper {
	if GlobalEnv.Eval(fn_car(args)) == GlobalEnv.Eval(fn_car(fn_cdr(args))) {
		return TEE
	} else {
		return NIL
	}
}
func func_car(args Sexper) Sexper {
	return fn_car(GlobalEnv.Eval(fn_car(args)))
}

func func_cdr(args Sexper) Sexper {
	return fn_cdr(GlobalEnv.Eval(fn_car(args)))
}

func func_cons(args Sexper) Sexper {
	return fn_cons(GlobalEnv.Eval(fn_car(args)),
		GlobalEnv.Eval(fn_car(fn_cdr(args))))
}

func func_cond(args Sexper) Sexper {
	for list := args; list != NIL; list = fn_cdr(list) {
		if GlobalEnv.Eval(fn_car(fn_car(list))) != NIL {
			return GlobalEnv.Eval(fn_car(fn_cdr(fn_car(list))))
		}
	}
	return NIL
}

func func_label(args Sexper) Sexper {
	key := fn_car(args)
	val := GlobalEnv.Eval(fn_car(fn_cdr(args)))
	GlobalEnv = GlobalEnv.Set(key, val)
	return key
}

func func_eval(args Sexper) Sexper {
	return GlobalEnv.Eval(GlobalEnv.Eval(fn_car(args)))
}

func func_read(args Sexper) Sexper {
	if !fn_car(args).Stringp() {
		panic("read require a string args")
	}
	return ReadFrom(fn_car(args).(String).val)
}

// Callable arithmetic
func func_add(args Sexper) Sexper {
	result := 0
	for list := args; list != NIL; list = fn_cdr(list) {
		result += fn_car(list).(Number).val
	}
	return Number{result}
}

func func_sub(args Sexper) Sexper {
	var result int
	if args != NIL {
		result = fn_car(args).(Number).val
		if fn_cdr(args) == NIL {
			return Number{-result}
		}
	} else {
		result = 0
	}
	for list := fn_cdr(args); list != NIL; list = fn_cdr(list) {
		result -= fn_car(list).(Number).val
	}
	return Number{result}
}

func func_mul(args Sexper) Sexper {
	result := 1
	for list := args; list != NIL; list = fn_cdr(list) {
		result *= fn_car(list).(Number).val
	}
	return Number{result}
}

func func_div(args Sexper) Sexper {
	result := fn_car(args).(Number).val
	for list := fn_cdr(args); list != NIL; list = fn_cdr(list) {
		result /= fn_car(list).(Number).val
	}
	return Number{result}
}

func func_mod(args Sexper) Sexper {
	result := fn_car(args).(Number).val % fn_car(fn_cdr(args)).(Number).val
	return Number{result}
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
		fn := this.Eval(fn_car(sexp))
		if fn.Funcp() {
			return fn.(Func).val(fn_cdr(sexp))
		} else {
			panic(fmt.Sprintf("Invalid function: %s", fn))
		}
	default:
		return NIL
	}
}

var GlobalEnv = fn_cons(fn_cons(Atom{"os"}, Atom{"mac"}), NIL).(List)

func InitEnv() {
	GlobalEnv = GlobalEnv.Set(Atom{"os"}, Atom{"mac"})
	GlobalEnv = GlobalEnv.Set(Atom{"who"}, Atom{"siteshen"})
	GlobalEnv = GlobalEnv.Set(Atom{"editor"}, Atom{"emacs"})

	GlobalEnv = GlobalEnv.Set(Atom{"quote"}, Func{func_quote})
	GlobalEnv = GlobalEnv.Set(Atom{"atom"}, Func{func_atom})
	GlobalEnv = GlobalEnv.Set(Atom{"eq"}, Func{func_eq})
	GlobalEnv = GlobalEnv.Set(Atom{"car"}, Func{func_car})
	GlobalEnv = GlobalEnv.Set(Atom{"cdr"}, Func{func_cdr})
	GlobalEnv = GlobalEnv.Set(Atom{"cons"}, Func{func_cons})
	GlobalEnv = GlobalEnv.Set(Atom{"cond"}, Func{func_cond})

	GlobalEnv = GlobalEnv.Set(Atom{"label"}, Func{func_label})
	GlobalEnv = GlobalEnv.Set(Atom{"eval"}, Func{func_eval})
	GlobalEnv = GlobalEnv.Set(Atom{"read"}, Func{func_read})

	GlobalEnv = GlobalEnv.Set(Atom{"+"}, Func{func_add})
	GlobalEnv = GlobalEnv.Set(Atom{"-"}, Func{func_sub})
	GlobalEnv = GlobalEnv.Set(Atom{"*"}, Func{func_mul})
	GlobalEnv = GlobalEnv.Set(Atom{"/"}, Func{func_div})
	GlobalEnv = GlobalEnv.Set(Atom{"%"}, Func{func_mod})
}
