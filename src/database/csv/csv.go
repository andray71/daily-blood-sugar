package csv

import (
"../../utils"
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

func NewDb(foodPath , exercisePath string) *textCsv {
	foodTable := []food{}
	utils.ReadCsvFile(foodPath, func(record []string,_ int) {
		foodTable = append(foodTable,food{
			id: utils.ToIntOrPanic(record[0]),
			description: record[1],
			index: utils.ToIntOrPanic(record[2])})
	},true)
	exerciseSlice := []exercise{}
	utils.ReadCsvFile(exercisePath, func(record []string,_ int) {
		exerciseSlice = append(exerciseSlice,exercise{
			id: utils.ToIntOrPanic(record[0]),
			description: record[1],
			index: utils.ToIntOrPanic(record[2])})
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
