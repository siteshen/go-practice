package maxsum

import "math"

func MaxSum1(v []float64) (maxSofar float64) {
	length := len(v)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			sum := 0.0
			// sum of [i..j]
			for k := i; k <= j; k++ {
				sum += v[k]
			}
			maxSofar = math.Max(maxSofar, sum)
		}
	}
	return
}

func MaxSum2(v []float64) (maxSofar float64) {
	length := len(v)
	for i := 0; i < length; i++ {
		sum := 0.0
		for j := i; j < length; j++ {
			// sum of [i..j]
			sum += v[j]
			maxSofar = math.Max(maxSofar, sum)
		}
	}
	return
}

func MaxSum2b(v []float64) (maxSofar float64) {
	length := len(v)
	// length+1 avoid sumArray[-1]
	sumArray := make([]float64, length+1)
	for i := 0; i < length; i++ {
		sumArray[i+1] = sumArray[i] + v[i]
	}

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			maxSofar = math.Max(maxSofar, sumArray[j+1]-sumArray[i])
		}
	}

	return
}

// ma, mb, mc
func maxSum(v []float64, low, high int) float64 {
	if low > high { // zero elements
		return 0
	}
	if low == high { // one element
		return math.Max(0, v[low])
	}

	middle := (low + high) / 2
	// find max crossing to left
	lmax, sum := 0.0, 0.0
	for i := middle; i >= low; i-- {
		sum += v[i]
		lmax = math.Max(lmax, sum)
	}

	// find max crossing to right
	rmax, sum := 0.0, 0.0
	for i := middle + 1; i <= high; i++ {
		sum += v[i]
		rmax = math.Max(rmax, sum)
	}

	mc := lmax + rmax

	// recusively left && right
	maxNow := math.Max(maxSum(v, low, middle), maxSum(v, middle+1, high))
	maxNow = math.Max(maxNow, mc)

	return maxNow
}

func MaxSum3(v []float64) (maxSofar float64) {
	return maxSum(v, 0, len(v)-1)
}

// MaxSumSuite.BenchmarkMaxSum	      20	  96768717 ns/op
func MaxSum4(v []float64) (maxSofar float64) {
	maxHere := 0.0
	for length, i := len(v), 0; i < length; i++ {
		maxHere = math.Max(maxHere+v[i], 0)
		maxSofar = math.Max(maxSofar, maxHere)
	}
	return
}

var MaxSum = MaxSum4
