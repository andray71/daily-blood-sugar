package main

import "../simulator"
import "../config"
import "../database"
import (
	"../input"
)

func main() {
 	db := database.ConnectToCsvDatabase("testData/FoodDB.csv","testData/Exercise.csv")

	events := input.ReadCsv("testData/input1.csv")
	sim,err := simulator.NewSimulator(config.NewSimulatorConfig(),db).Run(events)
	if err != nil {
		panic(err)
	}

	println("Chart: Blood Sugar")
	println(sim.GetBloodSugarChart().StringCsv())


	//println("Chart: Glycation")
	//println(sim.GetGlycationChart().StringCsv())
}
