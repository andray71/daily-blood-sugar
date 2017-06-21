package main


import (
"../simulator"
 "../config"
 "../database"
	"../input"
	"../utils"
	"flag"
	"os"
	"encoding/csv"
	"time"
	"fmt"
)

func main() {
	csvFood := flag.String("foodcsv", "testData/FoodDB.csv", "Food data csv file path")
	csvExercise := flag.String("exercisecsv", "testData/Exercise.csv", "Exercise data csv file path")
	in := flag.String("in", "testData/input1.csv", "Input data csv file path")
	out := flag.String("out", "", "[optional] Output data csv file path")

	flag.Parse()

	var file *os.File
	var err error
	if *out != "" {
		file, err = os.Create(*out)
		defer file.Close()
		if err != nil {
			panic(err)
		}
	} else {
		file = os.Stdout
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	events := input.ReadCsv(*in)
	_, err = simulator.NewSimulator(config.NewSimulatorConfig(),
		database.ConnectToCsvDatabase(*csvFood, *csvExercise)).Run(events, func(t time.Time, bloodSugar float64, glycation int) {
		w.Write([]string{
			t.Format(utils.DateTimeFormat),
			fmt.Sprintf("%f", bloodSugar),
			fmt.Sprintf("%d", glycation),
		})
	})
	if err != nil {
		panic(err)
	}
}
