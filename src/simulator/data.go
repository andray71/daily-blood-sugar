package simulator

import (
	"time"
	"fmt"
	"bytes"
)
type Chart []data

type data struct {
	time time.Time
	value int
}

func (s data) String() string {
	return fmt.Sprintf("{%s, %d}",s.time.String(),s.value)
}

func (s data) StringCsv() string {
	return fmt.Sprintf("%s, %d",s.time.String(),s.value)
}

func (s Chart) StringCsv() string {
	buffer := bytes.Buffer{}
	for _, d := range s {
		buffer.WriteString(fmt.Sprintln(d.StringCsv()))
	}
	return buffer.String()
}