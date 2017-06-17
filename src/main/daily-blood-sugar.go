package main

import "../simulator"
import "../config"
import "../database"
import (
	"../input"
)

func main() {
 	db := database.ConnectToCsvDatabase("testData/FoodDB.csv","testData/Exercise.csv")

	sim := simulator.NewSimulator(config.NewSimulatorConfig(),db).Run([]input.Event{})
	println(sim.GetGlycation())
}
