package simulator

import "../config"
import "../database"
import (
	"../input"
	"time"
	"errors"
	"fmt"
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

func NewSimulator(conf config.Simulator,db database.Database) Simulator {
	return Simulator{
		Simulator:         conf,
		currentGlycation:  0,
		currentBloodSugar: conf.MinBloodSugar,
		db:                db,
	}
}
func (s *Simulator) updateGlycation(t time.Time){
	oldValue := s.currentGlycation
	if s.currentBloodSugar >= s.BloodSugarLimitToEntreesGlycation {
		s.currentGlycation++
	}

	if len(s.glycation) == 0 || oldValue != s.currentGlycation {
		s.glycation = append(s.glycation,data{time:t,value:s.currentGlycation})
	}
}

func (s *Simulator) updateCharts(t time.Time) {

	if len(s.bloodSugar) == 0 || s.currentBloodSugar != s.bloodSugar[len(s.bloodSugar)-1].value {
		s.bloodSugar = append(s.bloodSugar, data{time: t, value: s.currentBloodSugar})
	}
	s.updateGlycation(t)
}
func (s *Simulator) processNormalizationEvent(e input.Event) {

	if e.GetTime().Before(s.normalizationLockTime) {
		return
	}
	if s.currentBloodSugar <= s.MinBloodSugar {
		s.currentBloodSugar = s.MinBloodSugar
	} else {
		s.currentBloodSugar--
	}
	s.updateCharts(e.GetTime())
}
func (s *Simulator) processFoodEvent(e input.Food) (err error) {
	if idx, ok := s.db.GetFoodIndex(e.Id); ok {
		s.currentBloodSugar += idx
	} else {
		err = errors.New(fmt.Sprintf("Index not found for Food id %d", e.Id))
	}
	s.updateCharts(e.GetTime())
	return
}

func (s *Simulator) processExerciseEvent(e input.Exercise) (err error) {
	if idx, ok := s.db.GetExerciseIndex(e.Id); ok {
		s.currentBloodSugar -= idx
		if s.currentBloodSugar < s.MinBloodSugar {
			s.currentBloodSugar = s.MinBloodSugar
		}
	} else {
		err = errors.New(fmt.Sprintf("Index not found for Exercise id %d", e.Id))
	}
	s.updateCharts(e.GetTime())
	return
}
func (s *Simulator) processEvent(e input.Event) (err error){
	switch eType := interface{}(e).(type) {
	case input.Food:
		err = s.processFoodEvent(eType)
	case input.Exercise:
		err = s.processExerciseEvent(eType)
	case input.Event:
		s.processNormalizationEvent(eType)
	default:
		err = errors.New("Unknown event type")
	}
	return
}
func (s Simulator) Run(events []input.Event) (sim Simulator,err error) {

	sim = s
	if len(events) == 0 {
		return
	}

	input.Sort(events)

	if s.currentTime.Equal(time.Time{}) {
		begin := events[0].GetTime()
		begin = time.Date(begin.Year(), begin.Month(), begin.Day(), 0, 0, 0, 0, begin.Location())
		events = append([]input.Event{input.NewEvent(begin)},events...)
	}

	tLine := events[0].GetTime()
	for {
		nextEventTime := events[0].GetTime()
		if tLine.Equal(nextEventTime){
			err = s.processEvent(events[0])
		}
		if tLine.Equal(nextEventTime) || tLine.After(nextEventTime){
			events = events[1:]
		}

		if len(events) == 0 {
			break
		}
		tLine = tLine.Add(time.Minute)
	}
	return
}
func (s Simulator)GetGlycation() Chart  {
	return s.glycation
}
func (s Simulator)GetBloodSugar() Chart  {
	return s.bloodSugar
}