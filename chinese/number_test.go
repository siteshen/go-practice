package chinese

import "testing"

func TestToChinese(t *testing.T) {
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
			t.Errorf("fail %d, %s, %s", test.in, ToChinese(test.in), test.out)
		}
	}
}
