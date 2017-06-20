package utils

import (
	"strconv"
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
	"io"
)

var DateTimeFormat = "2006-01-02 15:04:05"
func ToIntOrPanic(s string) (i int){
	i,err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("error convert %s to int",s))
	}
	return
}

func ReadCsvFile(path string, rowHandler func([]string, int), skipFirstRow bool) {
	f, _ := os.Open(path)

	i:=0
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		i++
		if skipFirstRow {
			skipFirstRow = false
			continue
		}
		rowHandler(record,i)
	}
}