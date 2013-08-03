package lisp

import (
	"fmt"
	"testing"
)

// testing helper
func AssertEqual(t *testing.T, sexp Sexper, output string) {
	if sexp_str := fmt.Sprintf("%s", sexp); sexp_str != output {
		t.Errorf("AssertEqual Error: %q, %q", sexp, output)
	}
}

func AssertEval(t *testing.T, env List, input string, output string) {
	sexp := env.Eval(ReadFrom(input))
	if sexp_str := fmt.Sprintf("%s", sexp); sexp_str != output {
		t.Errorf("AssertEval Error: %q -> %q, %q", input, sexp, output)
	}
}

func TestSexp(t *testing.T) {
	atom0 := Atom{"nil"}
	atom1 := Atom{"atom1"}
	atom2 := Atom{"atom2"}
	atom3 := Atom{"atom3"}
	list1 := List{Sexper(&atom1), Sexper(&atom2)}
	list2 := List{Sexper(&atom3), Sexper(&atom0)}
	list3 := List{Sexper(&list1), Sexper(&list2)}
	list4 := List{Sexper(&list2), Sexper(&list3)}

	AssertEqual(t, atom0, "nil")
	AssertEqual(t, atom3, "atom3")
	AssertEqual(t, list1, "(atom1 . atom2)")
	AssertEqual(t, list2, "(atom3)")
	AssertEqual(t, list3, "((atom1 . atom2) atom3)")
	AssertEqual(t, list4, "((atom3) (atom1 . atom2) atom3)")
}

func TestRead(t *testing.T) {
	var atom, list Sexper
	atom = ReadFrom("atom")
	AssertEqual(t, atom, "atom")

	atom = ReadFrom("12345")
	AssertEqual(t, atom, "12345")

	atom = ReadFrom("012345")
	AssertEqual(t, atom, "12345")

	atom = ReadFrom(`"hello world"`)
	AssertEqual(t, atom, `"hello world"`)

	list = ReadFrom("(hello . world)")
	AssertEqual(t, fn_car(list), "hello")
	AssertEqual(t, fn_cdr(list), "world")

	list = ReadFrom("(hello world)")
	AssertEqual(t, fn_car(list), "hello")
	AssertEqual(t, fn_cdr(list), "(world)")

	list = ReadFrom("(defun f (a b) (+ a b))")
	AssertEqual(t, fn_car(list), "defun")
	AssertEqual(t, fn_cdr(list), "(f (a b) (+ a b))")

	list = ReadFrom("((list . a) (+ a b) (it should works))")
	AssertEqual(t, fn_car(list), "(list . a)")
	AssertEqual(t, fn_cdr(list), "((+ a b) (it should works))")
}

func TestFunc(t *testing.T) {
	AssertEqual(t, fn_quote(ReadFrom("(quote a)")), "(quote a)")
	AssertEqual(t, fn_quote(ReadFrom("(cons a b)")), "(cons a b)")
	AssertEqual(t, fn_atom(ReadFrom("nil")), "t")
	AssertEqual(t, fn_atom(ReadFrom("()")), "t")
	AssertEqual(t, fn_atom(ReadFrom("(a)")), "nil")
	AssertEqual(t, fn_atom(ReadFrom("(a . b)")), "nil")
	AssertEqual(t, fn_car(ReadFrom("(a)")), "a")
	AssertEqual(t, fn_cdr(ReadFrom("(a)")), "nil")
	AssertEqual(t, fn_cdr(ReadFrom("(hello world)")), "(world)")
	AssertEqual(t, fn_car(ReadFrom("(hello world)")), "hello")
	AssertEqual(t, fn_cdr(ReadFrom("(hello world)")), "(world)")
	AssertEqual(t, fn_car(ReadFrom("((list a) (list b))")), "(list a)")
	AssertEqual(t, fn_cdr(ReadFrom("((list a) (list b))")), "((list b))")
	AssertEqual(t, fn_cons(ReadFrom("hello"), ReadFrom("world")), "(hello . world)")
	AssertEqual(t, fn_cons(ReadFrom("hello"), ReadFrom("(world)")), "(hello world)")
}

func TestEval(t *testing.T) {
	env := InitEnv()

	// arithmetic
	AssertEval(t, env, "(+ 10 3)", "13")
	AssertEval(t, env, "(- 10 3)", "7")
	AssertEval(t, env, "(* 10 3)", "30")
	AssertEval(t, env, "(/ 10 3)", "3")
	AssertEval(t, env, "(% 10 3)", "1")

	AssertEval(t, env, "(+ 3 10)", "13")
	AssertEval(t, env, "(- 3 10)", "-7")
	AssertEval(t, env, "(* 3 10)", "30")
	AssertEval(t, env, "(/ 3 10)", "0")
	AssertEval(t, env, "(% 3 10)", "3")

	// atom
	AssertEval(t, env, "os", "mac")
	AssertEval(t, env, "nil", "nil")
	AssertEval(t, env, "(quote who)", "who")

	// quote, atom
	AssertEval(t, env, "(atom ())", "t")
	AssertEval(t, env, "(atom '())", "t")
	AssertEval(t, env, "(atom '(a))", "nil")
	AssertEval(t, env, "(atom '(nil))", "nil")

	// car, cdr, cons
	AssertEval(t, env, "(cons 'a nil)", "(a)")
	AssertEval(t, env, "(cons 'a 'b)", "(a . b)")
	AssertEval(t, env, "(cons 'a '(b))", "(a b)")
	AssertEval(t, env, "(car (cons 'a 'b))", "a")
	AssertEval(t, env, "(cdr (cons 'a 'b))", "b")
	AssertEval(t, env, "(car '(cons 'a 'b))", "cons")
	AssertEval(t, env, "(cdr '(cons 'a 'b))", "((quote a) (quote b))")

	// eq
	AssertEval(t, env, "(eq (cons (quote a) 'b) '(a . b))", "t")
	AssertEval(t, env, "(eq 'nil 't)", "nil")
	AssertEval(t, env, "(eq who 'who)", "nil")
	AssertEval(t, env, "(eq editor os)", "nil")
	AssertEval(t, env, "(eq t 't)", "t")
	AssertEval(t, env, "(eq nil ())", "t")
	AssertEval(t, env, "(eq () '())", "t")
	AssertEval(t, env, "(eq os 'mac)", "t")

	// cond
	AssertEqual(t,
		env.Eval(ReadFrom(
			"(cond ((eq 'a 'b) 'error) ((eq 'b 'b) 'works))",
		)),
		"works")
	AssertEqual(t,
		env.Eval(ReadFrom(
			"(cond ((eq (cons 'a 'b) '(a . b)) 'works) ('t 'error))",
		)),
		"works")
	AssertEqual(t,
		env.Eval(ReadFrom(
			"(cond ((eq 'a 'b) 'error_ab) ((eq 'b 'c) 'error_bc))",
		)),
		"nil")
}
