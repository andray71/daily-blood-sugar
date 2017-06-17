package database

type Database interface {
	GetFoodIndex(id int) (int,bool)
	GetExerciseIndex(id int) (int,bool)
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
