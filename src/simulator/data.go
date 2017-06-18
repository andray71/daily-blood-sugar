package simulator

import (
	"time"
	"fmt"
)

type data struct {
	time time.Time
	value int
}

func (s data) String() string {
	return fmt.Sprintf("%s, %d",s.time.String(),s.value)
}