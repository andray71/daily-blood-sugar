package csv

import (
	"os"
	"encoding/csv"
	"bufio"
	"io"
	"fmt"
	"strconv"
)

type textCsv struct {
	food []food
	exercise []exercise
}
func (s *textCsv)GetFoodIndex(id int) (index int,ok bool) {
	for _, item := range s.food {
		if item.id == id {
			ok = true
			index = item.index
			return
		}
	}
	return
}
func (s *textCsv)GetExerciseIndex(id int) (index int,ok bool){
	for _, item := range s.exercise {
		if item.id == id {
			ok = true
			index = item.index
			return
		}
	}
	return
}

func readCsvFile(path string, rowHandler func([]string), skipFirstRow bool){
	f, _ := os.Open(path)

	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if skipFirstRow {
			skipFirstRow = false
			continue
		}
		rowHandler(record)
	}}

func toInt(s string) (i int){
 i,err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("error convert %s to int",s))
	}
	return
}
func NewDb(foodPath , exercisePath string) *textCsv {
	foodTable := []food{}
	readCsvFile(foodPath, func(record []string) {
		foodTable = append(foodTable,food{id: toInt(record[0]),description: record[1],index: toInt(record[2])})
	},true)
	exerciseSlice := []exercise{}
	readCsvFile(exercisePath, func(record []string) {
		exerciseSlice = append(exerciseSlice,exercise{id:toInt(record[0]),description:record[1],index:toInt(record[2])})
	},true)
  return &textCsv{food: foodTable,exercise: exerciseSlice}
}
type food struct {
	id int
	description string
	index int
}
type exercise struct {
	id int
	description string
	index int
}
