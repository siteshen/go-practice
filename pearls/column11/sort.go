package column11

func swap(x []int, i, j int) {
	x[i], x[j] = x[j], x[i]
}

// InsertionSort
//
// invariant: x[0..i-1] is sorted
// goal: sift x[i] down to its proper place in x[0..i]
func Isort1(x []int) {
	length := len(x)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && x[j-1] > x[j]; j-- {
			swap(x, j-1, j)
		}
	}
}

func Isort2(x []int) {
	length := len(x)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && x[j-1] > x[j]; j-- {
			t := x[j]
			x[j] = x[j-1]
			x[j-1] = t
		}
	}
}

func Isort3(x []int) {
	length := len(x)
	for i := 1; i < length; i++ {
		t := x[i]
		var j int
		for j = i; j > 0 && x[j-1] > t; j-- {
			x[j] = x[j-1]
		}
		x[j] = t
	}
}

var Sort = Isort3
