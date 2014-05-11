package maxsum

import (
	"math/rand"
	"testing"
	"time"

	check "gopkg.in/check.v1"
)

// define a test suite
type MaxSumSuite struct{ input []float64 }

// registe test suite
func init() { check.Suite(&MaxSumSuite{}) }

// make `go test` happy
func TestMe(t *testing.T) { check.TestingT(t) }

func (s *MaxSumSuite) Comment() check.CommentInterface {
	return check.Commentf("%+v", s.input)
}

func (s *MaxSumSuite) SimpleTest(c *check.C, output float64) {
	c.Check(MaxSum(s.input), check.Equals, output, s.Comment())
}

// 冒烟测试
func (s *MaxSumSuite) TestSmoke(c *check.C) {
	s.input = []float64{31, -41, 59, 26, -53, 58, 97, -93, -23, 84}
	s.SimpleTest(c, 187)
}

// 边界测试
func (s *MaxSumSuite) TestBoundary(c *check.C) {
	s.input = []float64{}
	s.SimpleTest(c, 0)

	s.input = []float64{-42}
	s.SimpleTest(c, 0)

	s.input = []float64{42}
	s.SimpleTest(c, 42)

	s.input = []float64{-1, -2, -3}
	s.SimpleTest(c, 0)

	s.input = []float64{1, 2, 3, 4, 5}
	s.SimpleTest(c, 15)
}

// 随机测试
func (s *MaxSumSuite) TestRandom(c *check.C) {
	maxSize, testCount := 100, 1000
	// test count N
	for i := 0; i < testCount; i++ {
		// random input (size)
		size := rand.Intn(maxSize)
		s.input = make([]float64, size)
		for j := 0; j < size; j++ {
			s.input[j] = rand.NormFloat64() * 100000
		}

		// check output
		// float64 equal -> int64 equal
		c.Check(int64(MaxSum(s.input)), check.Equals, int64(MaxSum1(s.input)), s.Comment())
		c.Check(int64(MaxSum(s.input)), check.Equals, int64(MaxSum2(s.input)), s.Comment())
		c.Check(int64(MaxSum(s.input)), check.Equals, int64(MaxSum2b(s.input)), s.Comment())
	}

	rand.Seed(time.Now().UnixNano())
	rand.NormFloat64()
}

// 性能测试
func (s *MaxSumSuite) BenchmarkMaxSum(c *check.C) {}
