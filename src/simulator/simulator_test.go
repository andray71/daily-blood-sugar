package simulator

import (
	"testing"
	"../database"
	"../config"
)

func DB4Test() database.Database  {
	return database.ConnectToCsvDatabase("../../testData/FoodDB.csv","../../testData/Exercise.csv")
}

func TestSimulator(t *testing.T) {

	simulator:= NewSimulator(config.NewSimulatorConfig(),DB4Test())
	t.Log(simulator.currentGlycation)

}
