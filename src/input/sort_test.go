package input

import (
	"testing"
	"time"
	"reflect"
)

func TestSort(t *testing.T) {
	events := []Event{
		NewExercise(1,time.Now().Add(time.Hour*3)),
		NewFood(1,time.Now().Add(time.Hour*5)),
		NewFood(1,time.Now().Add(time.Hour*2)),
		NewExercise(1,time.Now().Add(time.Hour)),
	}
	newEvents := make([]Event,len(events))
	for i,e := range events {
		println(e.GetTime().String(),reflect.TypeOf(e).String())
		newEvents[i] = events[i]
	}
	println("sorting")

	Sort(newEvents)
	println("non sorted")
	for _,e := range events {
		println(e.GetTime().String(),reflect.TypeOf(e).String())
	}
	println("sorted")
	for _,e := range newEvents {
		println(e.GetTime().String(),reflect.TypeOf(e).String())
	}
	if(!events[0].GetTime().Equal(newEvents[2].GetTime())){
		t.Fatal("not match")
	}
	if(!events[1].GetTime().Equal(newEvents[3].GetTime())){
		t.Fatal("not match")
	}
	if(!events[2].GetTime().Equal(newEvents[1].GetTime())){
		t.Fatal("not match")
	}
	if(!events[3].GetTime().Equal(newEvents[0].GetTime())){
		t.Fatal("not match")
	}
}
