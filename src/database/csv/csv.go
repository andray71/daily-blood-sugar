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

func readCsvFile(path string, rowHandler func([]string)){
	f, _ := os.Open(path)

	r := csv.NewReader(bufio.NewReader(f))
	first := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if first {
			first = false
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
func NewDataBaseFromCsv(foodPath string, exercisePath string) *textCsv {
	foodSlice := []food{}
	readCsvFile(foodPath, func(record []string) {
		foodSlice = append(foodSlice,food{id:toInt(record[0]),description:record[1],index:toInt(record[2])})
	})
	exerciseSlice := []exercise{}
	readCsvFile(exercisePath, func(record []string) {
		exerciseSlice = append(exerciseSlice,exercise{id:toInt(record[0]),description:record[1],index:toInt(record[2])})
	})
  return &textCsv{food: foodSlice,exercise: exerciseSlice}
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
