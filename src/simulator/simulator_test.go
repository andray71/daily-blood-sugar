package simulator

import (
	"testing"
	"../database"
	"../config"
	"../input"
)

func DB4Test() database.Database  {
	return database.ConnectToCsvDatabase("../../testData/FoodDB.csv","../../testData/Exercise.csv")
}

var events = input.ReadCsv("../../testData/input1.csv")

func TestSimulator(t *testing.T) {

	simulator:= NewSimulator(config.NewSimulatorConfig(),DB4Test())
	simulator,err := simulator.Run(events)

	if err != nil{
		t.Fatal(err)
	}
	t.Log("Printing charts")
	t.Log(simulator.GetBloodSugar().StringCsv())
//	t.Log(simulator.GetGlycation().StringCsv())
}
