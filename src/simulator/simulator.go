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
	if s.currentBloodSugar > s.BloodSugarLimitToEntreesGlycation {
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
func (s *Simulator) processNormalizationEvent(t time.Time) {

	if t.Before(s.normalizationLockTime) {
		return
	}
	if s.currentBloodSugar <= s.MinBloodSugar {
		s.currentBloodSugar = s.MinBloodSugar
	} else {
		s.currentBloodSugar--
	}
	s.updateCharts(t)
}

func (s *Simulator) updateNormalizationLock(d time.Duration){
	newLock := s.normalizationLockTime.Add(d)
	if newLock.After(s.normalizationLockTime){
		s.normalizationLockTime = newLock
	}
}

func (s *Simulator) processFoodEvent(e input.Food) (err error) {
	if idx, ok := s.db.GetFoodIndex(e.Id); ok {
		s.currentBloodSugar += idx
		s.updateNormalizationLock(s.FoodLoock)
	} else {
		err = errors.New(fmt.Sprintf("Index not found for Food id %d", e.Id))
	}
	s.updateCharts(e.GetTime())
	return
}

func (s *Simulator) processExerciseEvent(e input.Exercise) (err error) {
	if idx, ok := s.db.GetExerciseIndex(e.Id); ok {
		s.currentBloodSugar -= idx
		s.updateNormalizationLock(s.ExerciseLock)
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
		s.processNormalizationEvent(eType.GetTime())
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

	currentTime := events[0].GetTime()
	nextEventTime := currentTime
	for {
		if currentTime.Equal(nextEventTime){
			err = s.processEvent(events[0])
		} else if currentTime.Before(nextEventTime){
			s.processNormalizationEvent(currentTime)
		}

		if currentTime.Equal(nextEventTime) || currentTime.After(nextEventTime){
			events = events[1:]
			if len(events) > 0 {
				nextEventTime = events[0].GetTime()
			}
		}

		if len(events) == 0 {
			break
		}
		currentTime = currentTime.Add(time.Minute)
	}
	return
}
func (s Simulator)GetGlycationChart() Chart  {
	return s.glycation
}
func (s Simulator)GetBloodSugarChart() Chart  {
	return s.bloodSugar
}