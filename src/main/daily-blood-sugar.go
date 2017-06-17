package main

import "../simulator"
import "../config"
import "../database"
func main() {
 	dataBase := database.ConnectToCsvDatabase("testData/FoodDB.csv","testData/Exercise.csv")
	if index, ok := dataBase.GetExerciseIndex(2); ok {
		println("index",index)
	}
	sim := simulator.NewSimulator(config.NewSimulatorConfig()).Run(dataBase)
	println(sim.GetGlycation())
}
