package input

import (
	"testing"
	"time"
)

func TestSort(t *testing.T) {

	events := []Event{
		NewExercise(1,time.Now().Add(time.Hour*3)),
		NewFood(1,time.Now().Add(time.Hour*5)),
		NewFood(1,time.Now().Add(time.Hour*2)),
		NewExercise(1,time.Now().Add(time.Hour)),
	}

	for _,e := range events {
		println(e.GetTime().String())
	}
	println("sorting")
	Sort(events)
	for _,e := range events {
		println(e.GetTime().String())
	}
}
