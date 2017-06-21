package csv

import (
"../../utils"
)

type textCsv struct {
	food map[int]food
	exercise map[int]exercise
}
func (s *textCsv)GetFoodIndex(id int) (index int,ok bool) {
	var row food
	if row,ok = s.food[id]; ok {
		index = row.index
	}
	return
}
func (s *textCsv)GetExerciseIndex(id int) (index int,ok bool){
	var row exercise
	if row,ok = s.exercise[id]; ok {
		index = row.index
	}
	return
}

func NewDb(foodPath , exercisePath string) *textCsv {
	foodTable := map[int]food{}
	utils.ReadCsvFile(foodPath, func(record []string,_ int) {
		id := utils.ToIntOrPanic(record[0])
		foodTable[id] = food{
			id: id,
			description: record[1],
			index: utils.ToIntOrPanic(record[2])}
	},true)
	exerciseTable := map[int]exercise{}
	utils.ReadCsvFile(exercisePath, func(record []string,_ int) {
		id := utils.ToIntOrPanic(record[0])
		exerciseTable[id] = exercise{
			id: id,
			description: record[1],
			index: utils.ToIntOrPanic(record[2])}
	},true)
  return &textCsv{food: foodTable,exercise: exerciseTable}
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
