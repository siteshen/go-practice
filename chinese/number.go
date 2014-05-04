package chinese

import (
	"bytes"
	"fmt"
	"math"
)

var Number1 = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var Number2 = []string{"", "十", "百", "千"}
var Number3 = []string{
	"",                 // 10^0
	"万", "亿", "兆", "京", // 10^16
	"垓", "秭", "穰", "沟", // 10^32
	"涧", "正", "载", "极", // 10^48
	// "恒河沙", "阿僧祇", "那由他", "不可思议", "无量大数", // 10^68
}

func Reverse(in []string) (out []string) {
	length := len(in)

	for i := length - 1; i >= 0; i-- {
		out = append(out, in[i])
	}
	return
}

// convert number < 10k to chinese
//
// maybeZero: has higher number, can be zero
// lastZero:  last one contains zero, ignore here
func kiloToChinese(number int, maybeZero, lastZero bool) (string, bool, bool) {
	var buffer bytes.Buffer

	for i, str := range Reverse(Number2) {
		n := int(math.Pow10(3 - i))
		switch {
		case number >= n:
			buffer.WriteString(Number1[number/n] + str)
			number = number % n
			maybeZero = true
			lastZero = false
		case maybeZero:
			if !lastZero && number != 0 {
				buffer.WriteString("零")
				lastZero = true
			}
		}
	}

	return buffer.String(), maybeZero, lastZero
}

func ToChinese(number int64) string {
	var buffer bytes.Buffer
	length := len(fmt.Sprintf("%d", int(number))) + 3
	exp := length / 4

	if number < 10 {
		return Number1[number]
	}

	maybeZero, lastZero := false, false
	for i, str := range Reverse(Number3[:exp]) {
		i = i

		var result string
		n := int64(math.Pow10(4 * (exp - i - 1)))
		kilo := number / n
		result, maybeZero, lastZero = kiloToChinese(int(kilo), maybeZero, lastZero)
		buffer.WriteString(result + str)
		number = number % n
	}

	return buffer.String()
}

func main() {
	var tests = []struct {
		in  int64
		out string
	}{
		{1, "一"},
		{12, "十二"},
		{123, "一百二十三"},
		{1234, "一千二百三十四"},
		{12345, "一万二千三百四十五"},
		{123456, "一十二万三千四百五十六"},
		{1234567, "一百二十三万四千五百六十七"},
		{12345678, "一千二百三十四万五千六百七十八"},
		{123456789, "一亿二千三百四十五万六千七百八十九"},
		{1234567890, "一十二亿三千四百五十六万七千八百九十"},
		{12345678900, "一百二十三亿四千五百六十七万八千九百"},
		{123456789000, "一千二百三十四亿五千六百七十八万九千"},

		{0, "零"},
		{5, "五"},
		{123, "一百二十三"},
		{1234123, "一百二十三万四千一百二十三"},

		{101, "一百零一"},

		{10, "十"},
		{213, "二百一十三"},
		{1010, "一千零一十"},

		{1001, "一千零一"},
		{1000001, "一百万零一"},
		{320000032, "三亿二千万零三十二"},

		{10000, "一万"},
		{100200, "十万零两百"},
		{100010, "十万零一十"},
		{1000000000, "十亿"},
		{10001000, "一千万零一千"},
		{1000001000, "十亿零一千"},
	}

	for _, test := range tests {
		if ToChinese(test.in) != test.out {
			fmt.Println(test.in, ToChinese(test.in), test.out)
		}
	}
}
