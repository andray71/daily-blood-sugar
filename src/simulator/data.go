package simulator

import (
	"time"
	"fmt"
	"bytes"
	"encoding/csv"
	"../utils"
)
type Chart []data

type data struct {
	time time.Time
	value float64
}

func (s data) String() string {
	return fmt.Sprintf("{%s, %d}",s.time.String(),s.value)
}

func (s data) StringCsv() string {
	return fmt.Sprintf("%s, %f",s.time.String(),s.value)
}

func (s Chart) StringCsv() string {
	buffer := bytes.Buffer{}
	for _, d := range s {
		buffer.WriteString(fmt.Sprintln(d.StringCsv()))
	}
	return buffer.String()
}

func (s Chart) WriteCsv(writer *csv.Writer){
	for _,d:= range s {
		writer.Write([]string{d.time.Format(utils.DateTimeFormat),fmt.Sprintf("%f",d.value)})
	}
}