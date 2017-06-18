package simulator

import "../config"
import "../database"
import (
	"../input"
	"time"
)

type Simulator struct {
	config.Simulator
	currentGlycation      int
	glycation []data
	bloodSugar []data
	currentBloodSugar     int
	currentTime           time.Time
	normalizationLockTime time.Time
	db                    database.Database
}
func (s *Simulator) updateGlycation(t time.Time){
	oldValue := s.currentGlycation
	if s.currentBloodSugar >= s.BloodSugarLimitToEntreesGlycation {
		oldValue++
	}

	if len(s.glycation) == 0 || oldValue != s.currentGlycation {
		s.glycation = append(s.glycation,data{time:t,value:s.currentGlycation})
	}
}

func (s *Simulator) processNormalization(e input.Event) {

	if e.GetTime().Before(s.normalizationLockTime) {
		return
	}

	if s.currentBloodSugar <= s.MinBloodSugar {
		s.currentBloodSugar = s.MinBloodSugar
	} else {
		s.currentBloodSugar--
	}
}

func NewSimulator(conf config.Simulator,db database.Database) Simulator {
	return Simulator{
		Simulator:         conf,
		currentGlycation:  0,
		currentBloodSugar: conf.MinBloodSugar,
		db:                db,
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

func (s Simulator) Run(events []input.Event) Simulator {

	if len(events) == 0 {
		return s
	}

	input.Sort(events)

	if s.currentTime.Equal(time.Time{}){
		begin := events[0].GetTime()
		begin = time.Date(begin.Year(),begin.Month(),begin.Day(),0,0,0,0,begin.Location())
	}



	for event := range events {
		switch eType := interface{}(event).(type) {
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
	return s.currentGlycation
}
func (s Simulator)GetBloodSugar() int  {
	return s.currentBloodSugar
}