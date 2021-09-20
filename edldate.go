package frequency

import (
	"fmt"
	"time"
)

type Edldate struct {
	X time.Time
	Y int
}

func (v *Edldate) Show() {
	fmt.Println(v.Y + 10)
}
