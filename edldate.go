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
	X time.Time
	Y int
}

func (v *Edldate) Show() {
    fmt.Printf("%s %d\n", Week,Year)
	fmt.Println(v.Y + 15)
}
