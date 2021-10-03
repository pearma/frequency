package frequency

import (
	"testing"
	"time"

)

func TestWithGreeting(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
	result, err := Hello("mac")
	if err == nil {
		t.Log("Starting unit test at " + nowString + "wth" + result)
	}
}
func TestWithEdl(t *testing.T) {
	x := Edldate{Quarter, 10}
	//x.Show()
	x.NextDate(time.Date(2021, 10, 3, 0, 0, 0, 0, time.UTC))
}
