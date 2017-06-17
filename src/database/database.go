package database

import "./csv"

type Database interface {
	GetFoodIndex(id int) (int,bool)
	GetExerciseIndex(id int) (int,bool)
}

func ConnectToCsvDatabase(foodPath, exercisePath string) Database {
 return csv.NewDb(foodPath,exercisePath)
}