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
	currentGlycation      float64
	glycation []data
	bloodSugar []data
	currentBloodSugar     float64
	currentTime           time.Time

	bloodSugarAffectTable []data

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
func (s *Simulator) updateBloodSugar(t time.Time){
	newBloodSugarAffectTable := []data{}
	for _,d := range s.bloodSugarAffectTable {
		if  t.Before(d.time){
			s.currentBloodSugar += d.value
			newBloodSugarAffectTable = append(newBloodSugarAffectTable,d)
		}
	}
	s.bloodSugarAffectTable = newBloodSugarAffectTable
    if len(s.bloodSugarAffectTable) == 0 {
		s.currentBloodSugar--
	}
	if s.currentBloodSugar < s.MinBloodSugar {
		s.currentBloodSugar = s.MinBloodSugar
	}
    s.updateCharts(t)
}
func (s *Simulator) updateCharts(t time.Time) {

	if len(s.bloodSugar) == 0 || s.currentBloodSugar != s.bloodSugar[len(s.bloodSugar)-1].value {
		s.bloodSugar = append(s.bloodSugar, data{time: t, value: s.currentBloodSugar})
	}

	s.updateGlycation(t)
}

func (s *Simulator) addBloodSugarAffectingItem(value float64, until time.Time){
	s.bloodSugarAffectTable = append(s.bloodSugarAffectTable,data{time:until,value:value})
}

func (s *Simulator) processFoodEvent(e input.Food) (err error) {
	if idx, ok := s.db.GetFoodIndex(e.Id); ok {
		s.addBloodSugarAffectingItem(float64(idx)/100+1,e.Time.Add(s.FoodLoock))
	} else {
		err = errors.New(fmt.Sprintf("Index not found for Food id %d", e.Id))
	}
	return
}
func (s *Simulator) processExerciseEvent(e input.Exercise) (err error) {
	if idx, ok := s.db.GetExerciseIndex(e.Id); ok {
		s.addBloodSugarAffectingItem((float64(idx)/100+1)*-1,e.Time.Add(s.ExerciseLock))
	} else {
		err = errors.New(fmt.Sprintf("Index not found for Exercise id %d", e.Id))
	}
	return
}
func (s *Simulator) processEvent(e input.Event) (err error){
	switch eType := interface{}(e).(type) {
	case input.Food:
		err = s.processFoodEvent(eType)
	case input.Exercise:
		err = s.processExerciseEvent(eType)
	}
	return
}
func (s Simulator) Run(events []input.Event) (ret Simulator,err error) {

	ref := &s
	if len(events) == 0 {
		ret = *ref
		return
	}

	input.Sort(events)

	if ref.currentTime.Equal(time.Time{}) {
		begin := events[0].GetTime()
		begin = time.Date(begin.Year(), begin.Month(), begin.Day(), 0, 0, 0, 0, begin.Location())
		events = append([]input.Event{input.NewEvent(begin)},events...)
	}

	currentTime := events[0].GetTime()
	nextEventTime := currentTime
	for {
		if !nextEventTime.After(currentTime) {

			newEvents := []input.Event{}
			for _, e := range events {
				if currentTime.Equal(e.GetTime()) {
					err = ref.processEvent(events[0])
				} else {
					newEvents = append(newEvents, e)
				}
			}
			events = newEvents

			if len(events) == 0 {
				break
			} else {
				nextEventTime = events[0].GetTime()
			}
		}
		ref.updateBloodSugar(currentTime)
		currentTime = currentTime.Add(time.Minute)
	}
	ret = *ref
	return
}
func (s Simulator)GetGlycationChart() Chart  {
	return s.glycation
}
func (s Simulator)GetBloodSugarChart() Chart  {
	return s.bloodSugar
}