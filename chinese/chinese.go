package chinese

type HeavenlyStem string  // 天干
type EarthlyBranch string // 地支
type Zodiac string        // 生肖
type Season string        // 季节
type Year string          // 年
type Month string         // 月
type Day string           // 日
type Hour string          // 时

var (
	FiveElements = []string{"木", "火", "土", "金", "水"}

	// misc
	YinYang = []string{"阴", "阳"}

	HeavenlyStems   = []HeavenlyStem{"甲", "乙", "丙", "丁", "戊", "已", "庚", "辛", "壬", "癸"}
	EarthlyBranches = []EarthlyBranch{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	Zodiacs = []Zodiac{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

	Seasons    = []Season{"春", "夏", "秋", "冬"}
	SolarTerms = []string{
		"立春", "雨水", "惊蛰", "春分", "清明", "谷雨", // 春
		"立夏", "小满", "芒种", "夏至", "小暑", "大暑", // 夏
		"立秋", "处暑", "白露", "秋分", "寒露", "霜降", // 秋
		"立冬", "小雪", "大雪", "冬至", "小寒", "大寒", // 冬
	}

	// year, month, day
	Years = []Year{
		"甲子", "乙丑", "丙寅", "丁卯", "戊辰", "己巳", "庚午", "辛未", "壬申", "癸酉",
		"甲戌", "乙亥", "丙子", "丁丑", "戊寅", "己卯", "庚辰", "辛巳", "壬午", "癸未",
		"甲申", "乙酉", "丙戌", "丁亥", "戊子", "己丑", "庚寅", "辛卯", "壬辰", "癸巳",
		"甲午", "乙未", "丙申", "丁酉", "戊戌", "己亥", "庚子", "辛丑", "壬寅", "癸卯",
		"甲辰", "乙巳", "丙午", "丁未", "戊申", "己酉", "庚戌", "辛亥", "壬子", "癸丑",
		"甲寅", "乙卯", "丙辰", "丁巳", "戊午", "己未", "庚申", "辛酉", "壬戌", "癸亥",
	}
	Months = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	Days   = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	Hours  = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	// misc
	Moments = []string{"未明", "拂晓", "早晨", "午前", "午后", "傍晚", "薄暮", "深夜"}
)

var (
	SolarMonthDays = []struct {
		Name     string
		Position int
		Month    int
		StartDay int
		EndDay   int
	}{
		{"立春", 315, 2, 3, 5}, {"雨水", 330, 2, 18, 20}, {"惊蛰", 345, 3, 5, 7},
		{"春分", 0, 3, 20, 22}, {"清明", 15, 4, 4, 6}, {"谷雨", 30, 4, 19, 21},
		{"立夏", 45, 5, 5, 7}, {"小满", 60, 5, 20, 22}, {"芒种", 75, 6, 5, 7},
		{"夏至", 90, 6, 20, 22}, {"小暑", 105, 7, 6, 8}, {"大暑", 120, 7, 22, 24},
		{"立秋", 135, 8, 7, 9}, {"处暑", 150, 8, 22, 24}, {"白露", 165, 9, 7, 9},
		{"秋分", 180, 9, 22, 24}, {"寒露", 195, 10, 7, 9}, {"霜降", 210, 10, 23, 24},
		{"立冬", 225, 11, 7, 8}, {"小雪", 240, 11, 21, 23}, {"大雪", 255, 12, 6, 8},
		{"冬至", 270, 12, 21, 23}, {"小寒", 285, 1, 5, 7}, {"大寒", 300, 1, 19, 21},
	}

	SolarTermLyrics = []string{
		"春雨惊春清谷天", "夏满芒夏暑相连", "秋处露秋寒霜降", "冬雪雪冬小大寒",
		"每月两节不变更", "最多相差一两天", "上半年来六廿一", "下半年来八廿三",
	}
	SeasonLyrics = []string{
		"打春阳气转", "雨水沿河边", "惊蛰乌鸦叫", "春分沥皮干", "清明忙种麦", "谷雨种大田", // 春
		"立夏鹅毛住", "小满雀来全", "芒种五月节", "夏至不纳棉", "小暑不算热", "大暑三伏天", // 夏
		"立秋忙打靛", "处暑动刀镰", "白露烟上架", "秋分无生田", "寒露不算冷", "霜降变了天", // 秋
		"立冬交十月", "小雪地封严", "大雪河叉上", "冬至不行船", "小寒进腊月", "大寒又一年", // 冬
	}
	SolarTermLyrics2 = []string{
		"立春雨水节", "惊蛰及春分", "清明并谷雨", // 春
		"立夏小满方", "芒种及夏至", "小暑大暑当", // 夏
		"立秋还处暑", "白露秋分忙", "寒露又霜降", // 秋
		"立冬小雪张", "大雪冬至节", "小寒大寒昌", // 冬
	}
)

func ChineseYear(year int) Year { return Years[(year-4)%60] }
func ChineseMonth(month, day int) (m string, err error) {
	return
}

func ChineseSeason(month, day int) (season string, err error) {
	// (2, 3) - (5, 5) - (8, 7), (11, 7)
	// 203, 505, 807, 1107

	month = month % 12
	day = day % 31
	val := month*100 + day
	switch {
	case month < 1 || month > 12:
		err = ErrorMonth
	case day < 1 || day > 31:
		err = ErrorDay
	case 203 <= val && val < 505:
		season = "春"
	case 505 <= val && val < 807:
		season = "夏"
	case 807 <= val && val < 1107:
		season = "秋"
	case val >= 1107 || val < 203:
		season = "冬"
	default:
		err = ErrorDate
	}
	return
}

func ChineseDay(day int) Day {
	return ""
}

func GetMoment(hour int) string {
	return Moments[hour/3]
}

func GetHour(hour int) string {
	return Hours[(hour+1)%24/2]
}

func GetZodiac(year int) Zodiac {
	return Zodiacs[(year-4)%12]
}
