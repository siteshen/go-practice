package underscore

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// func TestInt(t *testing.T) {
// 	input := IntSlice{1, 2, 3, 4, 5}
// 	Trible := func(num int) int { return num * 3 }
// 	IsOdd := func(num int) bool { return num%2 == 1 }
// 	Sum := func(x, y int) int { return x + y }

// 	fmt.Println(input.Map(Trible))
// 	fmt.Println(input.Reduce(Sum, 0))
// 	fmt.Println(input.Filter(IsOdd))
// }

func SliceEqual(t *testing.T, slice1, slice2 IntSlice) {
	pc, file, line, _ := runtime.Caller(1)
	splited := strings.Split(file, "/")
	file = runtime.FuncForPC(pc).Name()
	file = splited[len(splited)-1]

	if !slice1.Equal(slice2) {
		t.Errorf("%s:%d %+v != %+v", file, line, slice1, slice2)
	}
}

func IntEqual(t *testing.T, int1, int2 int) {
	pc, file, line, _ := runtime.Caller(1)
	splited := strings.Split(file, "/")
	file = runtime.FuncForPC(pc).Name()
	file = splited[len(splited)-1]

	if int1 != int2 {
		t.Errorf("%s:%d %+v != %+v", file, line, int1, int2)
	}
}

func Assert(t *testing.T, stmt bool) {
	pc, file, line, _ := runtime.Caller(1)
	splited := strings.Split(file, "/")
	file = runtime.FuncForPC(pc).Name()
	file = splited[len(splited)-1]

	if !stmt {
		t.Errorf("%s:%d %t", file, line, stmt)
	}
}

func TestRange(t *testing.T) {

	SliceEqual(t, Range(10), IntSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	SliceEqual(t, Range(5, 10), IntSlice{5, 6, 7, 8, 9})
	SliceEqual(t, Range(0, 10, 2), IntSlice{0, 2, 4, 6, 8})
	SliceEqual(t, Range(0, 11, 2), IntSlice{0, 2, 4, 6, 8, 10})
	SliceEqual(t, Range(10, 0, -3), IntSlice{10, 7, 4, 1})

	SliceEqual(t, Range(10).Take(3), Range(3))
	SliceEqual(t, Range(10).Initial(3), Range(7))
	SliceEqual(t, Range(10).Last(3), Range(7, 10))
	SliceEqual(t, Range(10).Rest(3), Range(3, 10))
}

func TestMap(t *testing.T) {
	Range(5).Each(func(x int) { fmt.Print(x) })

	SliceEqual(t, Range(5).Map(func(x int) int { return x * 2 }), IntSlice{0, 2, 4, 6, 8})
	IntEqual(t, Range(5).Reduce(func(x, y int) int { return x + y }, 0), 10)

	output, _ := IntSlice{1, 2, 3}.Find(func(x int) bool { return x%2 == 0 })
	Assert(t, output == 2)

	output, _ = IntSlice{1, 2, 3}.Find(func(x int) bool { return x%4 == 0 })
	Assert(t, output == 0)

	SliceEqual(t, Range(10).Filter(func(x int) bool { return x%2 == 0 }), IntSlice{0, 2, 4, 6, 8})
	SliceEqual(t, Range(10).Reject(func(x int) bool { return x%2 == 0 }), IntSlice{1, 3, 5, 7, 9})

	Assert(t, IntSlice{2, 4, 6, 8}.Every(func(x int) bool { return x%2 == 0 }))
	Assert(t, !IntSlice{2, 4, 6, 7}.Every(func(x int) bool { return x%2 == 0 }))
	Assert(t, IntSlice{2, 4, 6, 7}.Some(func(x int) bool { return x%2 == 0 }))
	Assert(t, !IntSlice{1, 3, 5, 7}.Some(func(x int) bool { return x%2 == 0 }))
	Assert(t, IntSlice{1, 3, 5, 7}.Contains(3))
	Assert(t, !IntSlice{1, 3, 5, 7}.Contains(4))

	IntEqual(t, IntSlice{5, 7, 3, 1, 9}.Max(), 9)
	IntEqual(t, IntSlice{5, 7, 3, 1, 9}.Min(), 1)
	IntEqual(t, IntSlice{1, 3, 5}.Size(), 3)

	SliceEqual(t, IntSlice{0, 1, 3, 5, 0}.Compact(), IntSlice{1, 3, 5})

	IntEqual(t, IntSlice{0, 1, 3, 5, 1}.IndexOf(1), 1)
	IntEqual(t, IntSlice{0, 1, 3, 5, 1}.IndexOf(9), -1)
	IntEqual(t, IntSlice{0, 1, 3, 5, 1}.LastIndexOf(1), 4)
	IntEqual(t, IntSlice{0, 1, 3, 5, 1}.LastIndexOf(9), -1)

	Assert(t, Range(3).Equal(IntSlice{0, 1, 2}))
}
