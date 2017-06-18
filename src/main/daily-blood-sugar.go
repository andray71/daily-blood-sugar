package main

import "../simulator"
import "../config"
import "../database"
import (
	"../input"
)

func main() {
 	db := database.ConnectToCsvDatabase("testData/FoodDB.csv","testData/Exercise.csv")

	sim,err := simulator.NewSimulator(config.NewSimulatorConfig(),db).Run([]input.Event{})
	if err != nil {
		panic(err)
	}

	println("Chart: Blood Sugar")
	println(sim.GetBloodSugar().StringCsv())


	println("Chart: Glycation")
	println(sim.GetGlycation().StringCsv())
}
