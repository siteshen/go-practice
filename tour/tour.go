package tour

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"net"
	"os"
	"reflect"
	"runtime"
	"time"
)

func run(tour func()) {
	funcptr := reflect.ValueOf(tour).Pointer()
	funcname := runtime.FuncForPC(funcptr).Name()
	fmt.Println("\n==================== RUN", funcname, "====================")
	tour()
	fmt.Println("==================== END", funcname, "====================")
}

// Hello, 世界
func tour1() {
	fmt.Println("Hello, 世界")
}

// Go local
func tour2() {
}

// The Go Playground
func tour3() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("And if you tyr to open a file:")
	fmt.Println(os.Open("filename"))

	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp", "www.google.com"))
}

// Packages
func tour4() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

// Imports
func tour5() {
	fmt.Println("Now you have %g problems",
		math.Nextafter(2, 3))
}

// Exported names
func tour6() {
	fmt.Println(math.Pi)
}

// Functions
func tour7() {
	var add = func(x int, y int) int {
		return x + y
	}

	fmt.Println(add(42, 13))
}

// Functions continued
func tour8() {
	var add = func(x, y int) int {
		return x + y
	}

	fmt.Println(add(42, 13))
}

// Multiple results
func tour9() {
	var swap = func(x, y string) (string, string) {
		return y, x
	}

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// Named results
func tour10() {
	var split = func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}

	fmt.Println(split(17))
}

// Variables
func tour11() {
	var i int
	var c, python, java bool

	fmt.Println(i, c, python, java)
}

// Variables with initializers
func tour12() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"

	fmt.Println(i, j, c, python, java)
}

// Short variable declarations
func tour13() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// Basic types
func tour14() {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

// Type conversions
func tour15() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(3*3 + 4*4))
	var z int = int(f)
	fmt.Println(x, y, z)
}

// Constants
func tour16() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", math.Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// Numberic Constatns
func tour17() {
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	var needInt = func(x int) int { return x*10 + 1 }
	var needFloat = func(x float64) float64 { return x * 0.1 }

	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// For
func tour18() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// For continued
func tour19() {
	sum := 1
	for sum = sum; sum < 1000; sum = sum {
		sum += sum
	}
	fmt.Println(sum)
}

// For is Go's "while"
func tour20() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// Forever
func tour21() {
	for {
		fmt.Println(`HACK: should break as it's "forever"`)
		break
	}
}

// If
func tour22() {
	var sqrt func(x float64) string
	sqrt = func(x float64) string {
		if x < 0 {
			return sqrt(-x) + "i"
		}
		return fmt.Sprint(math.Sqrt(x))
	}

	fmt.Println(sqrt(2), sqrt(-4))
}

// If with a short statement
func tour23() {
	var pow = func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		return lim
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// If and else
func tour24() {
	var pow = func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		} else {
			fmt.Printf("%g >= %g\n", v, lim)
		}
		// can't use v here, though
		return lim
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*
Exercise: Loops and Functions

As a simple way to play with functions and loops, implement the square root
function using Newton's method.

In this case, Newton's method is to approximate Sqrt(x) by picking a starting
point z and then repeating:

z = z - (z^2 - x) / (2*z)

To begin with, just repeat that calculation 10 times and see how close you get
to the answer for various values (1, 2, 3, ...).

Next, change the loop condition to stop once the value has stopped changing (or
only changes by a very small delta). See if that's more or fewer iterations. How
close are you to the math.Sqrt?

Hint: to declare and initialize a floating point value, give it floating point
syntax or use a conversion:

z := float64(1)
z := 1.0
*/
func tour25() {
	var Sqrt = func(x float64) float64 {
		return x
	}

	fmt.Println(Sqrt(2))
}

// TODO:
func answer25() {}

// Structs
func tour26() {
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})
}

// Structs Fields
func tour27() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

// Pointers
func tour28() {
	type Vertex struct {
		X int
		Y int
	}

	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}

func RunPart1() {
	funcList := []func(){
		tour1, tour2, tour3, tour4, tour5,
		tour6, tour7, tour8, tour9, tour10,
		tour11, tour12, tour13, tour14, tour15,
		tour16, tour17, tour18, tour19, tour20,
		tour21, tour22, tour23, tour24, tour25,
	}

	for _, tour := range funcList {
		run(tour)
	}
}

func RunPart2() {
	funcList := []func(){tour26, tour27, tour28}

	for _, tour := range funcList {
		run(tour)
	}
}
