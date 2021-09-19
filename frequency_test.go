package frequency

import (
	"testing"
	"time"
)

func TestWithGreeting(t *testing.T) {
	thisTime := time.Now()
	nowString := thisTime.Format("2006-01-02 15:04:05 Monday")
        result,err := Hello("mac")
        if err == nil {
		t.Log("Starting unit test at " + nowString+"wth"+result)
	}
}
