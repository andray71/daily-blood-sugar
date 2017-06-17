package simulator

import "../config"
import "../database"
import (
	"../input"
	"time"
)

type Simulator struct {
	config.Simulator
	glycation int
	bloodSugar int
	currentTime time.Time
	db database.Database
}

func NewSimulator(conf config.Simulator,db database.Database) Simulator {
	return Simulator{
		Simulator:conf,
		glycation:0,
		bloodSugar: conf.MinBloodSugar,
		db:db,
	}
}
func (s *Simulator) processFood(e input.Food){
	println("got food", e.Id)
	if idx, ok := s.db.GetFoodIndex(e.Id); ok {
		println("\tindex", idx)
	} else {
		println("\tindex not found", idx)
	}
}

func (s *Simulator) processExercise(e input.Exercise) {
	println("got exercise", e.Id)
	if idx, ok := s.db.GetExerciseIndex(e.Id); ok {
		println("\tindex", idx)
	} else {
		println("\tindex not found", idx)
	}
}

func (s Simulator) Run(in []input.Event) Simulator {
	if len(in) == 0 {
		return s
	}

	input.Sort(in)

	for event := range in {
		switch eType := interface{}( event).(type) {
		case input.Food:
			s.processFood(eType)
		case input.Exercise:
			s.processExercise(eType)
		case input.Event:
			println("got simple event")
		default:
			panic("Unknown event type")
		}
	}

	return s
}
func (s Simulator)GetGlycation() int  {
	return s.glycation
}
func (s Simulator)GetBloodSugar() int  {
	return s.bloodSugar
}