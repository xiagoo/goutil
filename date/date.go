package date

import (
	"log"
	"time"
)

const (
	YMDHMS = "2006-01-02 15:04:05"
	YMDHM  = "2006-01-02 15:04"
	YMDH   = "2006-01-02 15"
	YMD    = "2006-01-02"
	YM     = "2006-01"
	MD     = "01-02"
	Y      = "2006"
	M      = "01"
	D      = "02"
	HMS    = "15:04:05"
	HM     = "15:04"
	MS     = "04:05"
	H      = "15"
	S      = "05"
)

func GetConstellation(month, day int) (star string) {
	if month <= 0 || month >= 13 {
		star = "-1"
	}
	if day <= 0 || day >= 32 {
		star = "-1"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "双鱼座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "双子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "狮子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "处女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天秤座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蝎座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "魔蝎座"
	}
	return star
}

func GetZodiac(year int) (zodiac string) {
	if year <= 0 {
		zodiac = "-1"
	}
	start := 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}

func Stamp2S(layout string, stamp int) string {
	return time.Unix(int64(stamp), 0).Format(layout)
}

func S2Stamp(layout string, s string) int {
	t, err := time.ParseInLocation(layout, s, time.Local)
	if err != nil {
		log.Printf("S2Stamp:%s\n", err.Error())
		return 0
	}
	return int(t.Unix())
}

func S2Time(layout string, s string) *time.Time {
	t, err := time.ParseInLocation(layout, s, time.Local)
	if err != nil {
		log.Printf("S2Time:%s\n", err.Error())
		return nil
	}
	return &t
}

func Time2S(layout string, t time.Time) string {
	return t.Format(layout)
}

func Stamp2Time(stamp int) time.Time {
	return time.Unix(int64(stamp), 0)
}

func Time2Stamp(t time.Time) int {
	return int(t.Unix())
}

func Today0Clock() int {
	return Get0Clock(0)
}

func Get0Clock(d int) int {
	s := time.Now().Format(YMD)
	t, err := time.ParseInLocation(YMD, s, time.Local)
	if err != nil {
		log.Printf("Get0Clock:%s\n", err.Error())
		return 0
	}
	return int(t.AddDate(0, 0, d).Unix())
}

func AddSecond(s int) time.Time {
	return time.Now().Add(time.Second * time.Duration(s))
}

func AddMinute(m int) time.Time {
	return time.Now().Add(time.Minute * time.Duration(m))
}

func AddHour(h int) time.Time {
	return time.Now().Add(time.Hour * time.Duration(h))
}

func AddDay(d int) time.Time {
	return time.Now().AddDate(0, 0, d)
}

func AddMonth(m int) time.Time {
	return time.Now().AddDate(0, m, 0)
}

func AddYear(y int) time.Time {
	return time.Now().AddDate(y, 0, 0)
}