package frequency

import (
  "errors"
  "fmt"
  "math/rand"
  "time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string,error) {
    if name == "" {
        return "",errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message,nil
}

func init() {
  rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    return formats[rand.Intn(len(formats))]
}

type Edldate struct {
	X time.Time
        Y int
}

func (v *Edldate) Show() {
	fmt.Println(v.Y)
}
