package main

import "../simulator"
import "../config"
import "../database"
import (
	"../input"
	"flag"
	"strings"
	"os"
	"encoding/csv"
)

func main() {
	csvFood := flag.String("foodcsv", "testData/FoodDB.csv","Food data csv file path")
	csvExercise := flag.String("exercisecsv", "testData/Exercise.csv","Exercise data csv file path")
	inFile := flag.String("inputcsv", "testData/input1.csv","Input data csv file path")
	outBSFile := flag.String("bschart", "testData/bloodShugarChart.csv","Output Blood Sugar Chart csv file path")
	outGlFile := flag.String("glchart", "testData/glycationChart.csv","Output Glycation Chart csv file path")

	flag.Parse()

	events := input.ReadCsv(*inFile)
	sim,err := simulator.NewSimulator(config.NewSimulatorConfig(),database.ConnectToCsvDatabase(*csvFood,*csvExercise)).Run(events)
	if err != nil {
		panic(err)
	}


	if strings.HasSuffix(*outBSFile,".csv") {
		*outBSFile = *outBSFile + ".csv"
	}
	if strings.HasSuffix(*outGlFile,".csv") {
		*outGlFile = *outGlFile + ".csv"
	}

	WriteCsv(*outBSFile,sim.GetBloodSugarChart().WriteCsv)
	WriteCsv(*outGlFile,sim.GetGlycationChart().WriteCsv)
	println("Done")
}

func WriteCsv(path string, handler func(*csv.Writer)) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	handler(w)
}