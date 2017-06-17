package input

import "time"

type Event interface {
	GetTime() time.Time
}

type event struct {
	Time time.Time
}

func (s event) GetTime() time.Time {
	return s.Time
}

type Food struct {
	event
	Id int
}

type Exercise struct {
	event
	Id int
}

func NewEvent(t time.Time) Event {
	return event{t}
}

func NewFood(id int, t time.Time) Event {
	return Event(
		Food{
			event: event{t},
			Id:id,
		},
	)
}
func NewExercise(id int, t time.Time) Event {
	return Event(
		Exercise{
			event:event{t},
			Id:id,
		},
	)
}
