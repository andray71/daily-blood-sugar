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
	currentBloodSugar     float64
	currentTime           time.Time

	bloodSugarAffectTable []data

	db                    database.Database
}

func NewSimulator(conf config.Simulator,db database.Database) Simulator {
	return Simulator{
		Simulator:         conf,
		currentGlycation:  0,
		currentBloodSugar: 0,
		db:                db,
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
}

func (s *Simulator) updateGlycation(t time.Time){
	if s.currentBloodSugar > s.BloodSugarLimitToEncrimentGlycation {
		s.currentGlycation++
	}
}

func (s *Simulator) addBloodSugarAffectingItem(value float64, until time.Time){
	s.bloodSugarAffectTable = append(s.bloodSugarAffectTable,data{time:until,value:value})
}

func (s *Simulator) processFoodEvent(e input.Food) (err error) {
	if idx, ok := s.db.GetFoodIndex(e.Id); ok {
		s.addBloodSugarAffectingItem(float64(idx)/100+1,e.Time.Add(s.FoodLock))
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

func setTimeToZero(t time.Time) time.Time {
 return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func (s Simulator) Run(events []input.Event, receiver func(time.Time,float64,int)) (ret Simulator,err error) {

	ref := &s
	if len(events) == 0 {
		ret = *ref
		return
	}

	input.Sort(events)

	if ref.currentTime.Equal(time.Time{}) {
		begin := events[0].GetTime()
		begin = begin.Add(time.Minute* -1)
		events = append([]input.Event{input.NewEvent(begin)},events...)
	}

	end := setTimeToZero( events[len(events)-1].GetTime()).Add(time.Hour*24 - time.Minute)
	events = append(events, input.NewEvent(end))

	currentTime := events[0].GetTime()
	nextEventTime := currentTime

	for {

		lastBloodSugar :=ref.currentBloodSugar
		lastGlycation := ref.currentGlycation
		if !nextEventTime.After(currentTime) {

			newEvents := []input.Event{}
			for _, e := range events {
				if !currentTime.Before(e.GetTime()) {
					err = ref.processEvent(e)
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
		ref.updateGlycation(currentTime)
		if(lastBloodSugar != ref.currentBloodSugar || lastGlycation!=ref.currentGlycation) {
			receiver(currentTime, ref.currentBloodSugar, ref.currentGlycation)
		}
		currentTime = currentTime.Add(time.Minute)
	}
	ret = *ref
	return
}
