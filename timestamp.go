package timestamp

import (
	"fmt"
	"time"
)

type Source struct {
	Time int
}

func (s *Source)Date() (string, error) {
	t := time.Unix(1557277201, 0)
	// TODO
	fmt.Println(t)
}
