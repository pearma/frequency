package frequency

import (
	"fmt"
	"time"
)

type Freq int

const (
	Day Freq = iota
	DayB
	Week
	Tenday
	Month
	Quarter
	Halfyear
	Year
)

func (f Freq) String() string {
	switch f {
	case Day:
		return "Day"
	case DayB:
		return "DayB"
	case Week:
		return "Week"
	case Tenday:
		return "Tenday"
	case Month:
		return "Month"
	case Quarter:
		return "Quarter"
	case Halfyear:
		return "Halfyear"
	case Year:
		return "Year"
	}
	return "N/A"
}

type Edldate struct {
	freq   Freq
	offset int
}

func (v *Edldate) Show() {
	fmt.Printf("%s %d\n", v.freq, v.freq)
	fmt.Println(v.offset + 15)
}

func getOrdinal(freq Freq, base time.Time) int {

	year, month, day := base.Date()
	switch freq {
	case Week:
		return int(base.Weekday())
	case Month:
		fmt.Printf("%v %v %d\n", year, month, day)
		return base.Day()
	}
	return 0
}
func getStartOfQuarter(year int, month int) time.Time {
	switch {
	case month < 4:
		return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	case month < 7:
		return time.Date(year, 4, 1, 0, 0, 0, 0, time.UTC)
	case month < 10:
		return time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)
	case month <= 12:
		return time.Date(year, 10, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Now()
}

//获取起始日期
func getRangeOfTerm(freq Freq, base time.Time) (time.Time, time.Time) {
	var start time.Time
	var end time.Time
	year, month, day := base.Date()

	switch freq {
	case Week:
		var wkd int
		if base.Weekday() == time.Sunday {
			wkd = 7
		} else {
			wkd = int(base.Weekday())
		}
		start = time.Date(year, month, day-wkd+1, 0, 0, 0, 0, time.UTC)
		end = start.AddDate(0, 0, 6)
		return start, end
	case Month:
		start = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		end = start.AddDate(0, 1, -1)
		return start, end
	case Quarter:
		start = getStartOfQuarter(year, int(month))
		end = start.AddDate(0, 3, -1)
		return start, end

	}
	return time.Now(), time.Now()
}

func (v *Edldate) NextDate(base time.Time) {

	fmt.Printf("it is %d of %s\n", getOrdinal(v.freq, base), v.freq)
	s, e := getRangeOfTerm(v.freq, base)
	fmt.Printf("start  is %v, end is %v\n", s, e)

	if v.freq == Day {
		daysLater := base.AddDate(0, 0, v.offset)
		fmt.Printf("next valid date of %s is %s after  %d days \n", base.Format("2006-01-02"), daysLater.Format("2006-01-02"), v.offset)

	} else {
		fmt.Printf("next vald date called with %s\n", base.Format("2006-01-02"))
	}
}
