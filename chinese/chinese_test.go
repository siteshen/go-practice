package chinese

import "testing"

func TestChineseYear(t *testing.T) {
	tests := []struct {
		in  int
		out Year
	}{
		{1861, "辛酉"},
		{1898, "戊戌"},
		{1901, "辛丑"},
	}
	for i, tt := range tests {
		result := ChineseYear(tt.in)
		if result != tt.out {
			t.Errorf("%+v. ChineseYear(%+v) => %+v, want %+v\n", i, tt.in, result, tt.out)
		}
	}
}

func TestGetMoment(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{0, "未明"},
		{1, "未明"},
		{2, "未明"},
		{3, "拂晓"},
		{6, "早晨"},
		{9, "午前"},
		{12, "午后"},
		{15, "傍晚"},
		{18, "薄暮"},
		{21, "深夜"},
		{22, "深夜"},
		{23, "深夜"},
	}

	for i, tt := range tests {
		result := GetMoment(tt.in)
		if result != tt.out {
			t.Errorf("%+v. GetShiKe(%+v) => %+v, want %+v\n", i, tt.in, result, tt.out)
		}
	}
}

func TestGetHour(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{0, "子"},
		{1, "丑"},
		{2, "丑"},
		{3, "寅"},
		{5, "卯"},
		{7, "辰"},
		{9, "巳"},
		{11, "午"},
		{13, "未"},
		{15, "申"},
		{17, "酉"},
		{19, "戌"},
		{21, "亥"},
		{23, "子"},
	}

	for i, tt := range tests {
		result := GetHour(tt.in)
		if result != tt.out {
			t.Errorf("%+v. GetHour(%+v) => %+v, want %+v\n", i, tt.in, result, tt.out)
		}
	}
}

func TestGetZodiac(t *testing.T) {
	tests := []struct {
		in  int
		out Zodiac
	}{
		{1987, "兔"},
		{1988, "龙"},
	}

	for i, tt := range tests {
		result := GetZodiac(tt.in)
		if result != tt.out {
			t.Errorf("%+v. GetZodiac(%+v) => %+v, want %+v\n", i, tt.in, result, tt.out)
		}
	}
}
