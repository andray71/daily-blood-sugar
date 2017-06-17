package main

import "../simulator"
import "../config"
import "../database"
import (
	"../input"
)

func main() {
 	dataBase := database.ConnectToCsvDatabase("testData/FoodDB.csv","testData/Exercise.csv")

	sim := simulator.NewSimulator(config.NewSimulatorConfig(),dataBase).Run([]input.Event{})
	println(sim.GetGlycation())
}
