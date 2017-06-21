package simulator

import (
	"testing"
	"../database"
	"../config"
	"../input"
	"time"
	"fmt"
	"../utils"
)

func DB4Test() database.Database  {
	return database.ConnectToCsvDatabase("../../testData/FoodDB.csv","../../testData/Exercise.csv")
}

var events = input.ReadCsv("../../testData/input1.csv")

func TestSimulator(t *testing.T) {

	simulator:= NewSimulator(config.NewSimulatorConfig(),DB4Test())

	err := simulator.Run(events, func(time time.Time,bloodSugar float64,glycation int){
	 fmt.Println(time.Format(utils.DateTimeFormat),fmt.Sprintf("%f" ,bloodSugar),glycation)
	})

	if err != nil{
		t.Fatal(err)
	}
}
