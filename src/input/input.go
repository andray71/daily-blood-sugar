package input

import (
	"time"
	"../utils"
	"fmt"
)

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

func NewEvent(t time.Time) event {
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

func ReadCsv(path string) (events []Event) {
	layout := "2006-01-02 15:04:05"
	utils.ReadCsvFile(path, func(rec []string, i int) {
		ts, err := time.Parse(layout, rec[2])
		if err != nil {
			panic(fmt.Sprintf("Invalid time %s on row %d. %s", rec[2], i,err.Error()))
		}
		switch rec[0] {
		case "food":
			events = append(events, Food{
				event: NewEvent(ts),
				Id:    utils.ToIntOrPanic(rec[1]),
			})
		case "exercise":
			events = append(events, Exercise{
				event: NewEvent(ts),
				Id:    utils.ToIntOrPanic(rec[1]),
			})
		default:
			panic(fmt.Sprintf("Unknown Event %s",rec[0]))
		}
	}, true)

	return
}