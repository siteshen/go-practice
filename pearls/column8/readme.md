# 第八章 算法设计技术

## The Problem and a Simple Algorithm
    input:  A vector x of n floating-point numbers
    output: maximum sum found in any contiguous subvector of the input

    example: a[1..n] = [31, -41, 59, 26, -53, 58, 97, -93, -23, 84]
    output:  59 + 26 + (-53) + 58 + 97 = 187

算法1 - O(n^3)

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


## Two Quadratic Algorithms
算法2 - O(n^2)

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

算法2b - O(n^2)

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

## A Divide-and-Conquer Algorithm
算法3 - O(n log(n))

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

        // recursively left && right
        maxNow := math.Max(maxSum(v, low, middle), maxSum(v, middle+1, high))
        maxNow = math.Max(maxNow, mc)

        return maxNow
    }

    func MaxSum3(v []float64) (maxSofar float64) {
        return maxSum(v, 0, len(v)-1)
    }


## A Scanning Algorithm
算法4 - O(n)

    func MaxSum4(v []float64) (maxSofar float64) {
        maxHere := 0.0
        for length, i := len(v), 0; i < length; i++ {
            maxHere = math.Max(maxHere+v[i], 0)
            maxSofar = math.Max(maxSofar, maxHere)
        }
        return
    }

## What Does It Matter
## Principles
### Some important algorithm design techniques
1. Save state to avoid recomputation (MaxSum2, MaxSum4)
2. Preprocess information into data structure (MaxSum2b)
3. Divide-and-conquer algorithms (MaxSum3)
4. Scanning algorithms: Problems on arrays can often by solved by asking ``how
   can I extend a solution for x[0..i-1] to a solution for x[0..i] (MaxSum4)''
5. Cumulatives (MaxSum2b)
6. Lower bounds

## Problems

1. 使用第四章的技术，验证算法3和算法4代码的正确；
2. 在你的计算机上测试程序运行时间，生成类似8.5章的图表；
3. 精确计算max函数调用次数，空间复杂度？
4. 如果输入是[-1, 1]之间的随机值，那么最大子串的长度期望值是多少？
