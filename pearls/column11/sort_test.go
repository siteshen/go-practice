package column11

import (
	"math/rand"
	"sort"
	"testing"

	check "gopkg.in/check.v1"
)

// define a test suite
type IsortSuite struct{ input []int }

// registe test suite
func init() { check.Suite(&IsortSuite{}) }

// make `go test` happy
func TestMe(t *testing.T) { check.TestingT(t) }

func (s *IsortSuite) Comment() check.CommentInterface {
	return check.Commentf("%+v", s.input)
}

func (s *IsortSuite) SimpleTest(c *check.C) {
	Sort(s.input)

	c.Check(sort.IsSorted(sort.IntSlice(s.input)), check.Equals, true, s.Comment())
}

// 冒烟测试
func (s *IsortSuite) TestSmoke(c *check.C) {
	s.input = []int{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}
	s.SimpleTest(c)
}

// 边界测试
func (s *IsortSuite) TestBoundary(c *check.C) {
	s.input = []int{}
	s.SimpleTest(c)

	s.input = []int{42}
	s.SimpleTest(c)

	s.input = []int{-1, 2, -3}
	s.SimpleTest(c)

	s.input = []int{1, 3, 2, 4, 5}
	s.SimpleTest(c)
}

// 随机测试
func (s *IsortSuite) xTestRandom(c *check.C) {
	maxSize, testCount := 100, 1000
	// test count N
	for i := 0; i < testCount; i++ {
		// random input (size)
		size := rand.Intn(maxSize)
		s.input = make([]int, size)
		for j := 0; j < size; j++ {
			s.input[j] = int(rand.NormFloat64() * 100000)
		}

		s.SimpleTest(c)
	}
}

// 性能测试
func (s *IsortSuite) BenchmarkMaxSum(c *check.C) {}
