package search_test

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	. "../search"

	check "gopkg.in/check.v1"
)

// define a test suite
type BinarySearchSuite struct{ input []int }

// registe test suite
func init() { check.Suite(&BinarySearchSuite{}) }

// make `go test` happy
func TestMe(t *testing.T) { check.TestingT(t) }

// slow version search from front
func (s *BinarySearchSuite) getFirstLocation(v int) int {
	for i, length := 0, len(s.input); i < length; i++ {
		if s.input[i] == v {
			return i
		}
	}
	return -1
}

// slow version search from back
func (s *BinarySearchSuite) getLastLocation(v int) int {
	for i := len(s.input) - 1; i >= 0; i-- {
		if s.input[i] == v {
			return i
		}
	}
	return -1
}

func (s *BinarySearchSuite) Comment() check.CommentInterface { return check.Commentf("%v", s.input) }

func (s *BinarySearchSuite) Random(maxSize, maxValue int) {
	rand.Seed(time.Now().UnixNano())
	size := 1 + rand.Intn(maxSize)
	s.input = rand.Perm(maxValue)[:size]
	sort.Sort(sort.IntSlice(s.input))
}

// 冒烟测试
func (s *BinarySearchSuite) TestSmoke(c *check.C) {
	s.input = []int{1, 4, 42, 55, 67, 87, 100}
	c.Check(BinarySearch(s.input, 42), check.Equals, 2, s.Comment())
	c.Check(BinarySearch(s.input, 43), check.Equals, -1, s.Comment())
}

// 边界测试
func (s *BinarySearchSuite) TestBoundary(c *check.C) {
	// empty
	s.input = []int{}
	c.Check(BinarySearch(s.input, 42), check.Equals, -1, s.Comment())

	// one element
	s.input = []int{42}
	c.Check(BinarySearch(s.input, 42), check.Equals, 0, s.Comment())
	c.Check(BinarySearch(s.input, 43), check.Equals, -1, s.Comment())

	// first | middle | last
	s.input = []int{-324, -3, -1, 0, 42, 99, 101}
	c.Check(BinarySearch(s.input, -324), check.Equals, 0, s.Comment())
	c.Check(BinarySearch(s.input, 0), check.Equals, 3, s.Comment())
	c.Check(BinarySearch(s.input, 101), check.Equals, 6, s.Comment())

	// with minInt64 maxInt64
	s.input = []int{math.MinInt64, -324, -3, -1, 0, 42, 99, 101, math.MaxInt64}
	c.Check(BinarySearch(s.input, math.MinInt64), check.Equals, 0, s.Comment())
	c.Check(BinarySearch(s.input, math.MaxInt64), check.Equals, 8, s.Comment())
}

// 随机测试
func (s *BinarySearchSuite) TestRandom(c *check.C) {
	for i := 0; i < c.N; i++ {
		maxSize, maxValue := 100, 1000
		s.Random(maxSize, maxValue)
		v := maxValue / 2
		for j := 0; j < 10; j++ {
			v -= 1
			c.Check(BinarySearch(s.input, v), check.Equals, s.getFirstLocation(v))
			c.Check(BinarySearch(s.input, v), check.Equals, s.getLastLocation(v))
		}
	}
}

// 性能测试
func (s *BinarySearchSuite) BenchmarkSearch(c *check.C) {
	for i := 0; i < c.N; i++ {
		maxSize, maxValue := 1000, 1000
		s.Random(maxSize, maxValue)
		_, count := BinarySearchWithCounting1(s.input, maxValue)
		maxCount := int(math.Ceil(math.Log2(float64(maxSize))))

		c.Assert(count < maxCount, check.Equals, true, s.Comment())
	}
}
