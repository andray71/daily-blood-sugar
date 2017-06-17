package input

import "time"

type Event struct {
	Time time.Time
}
type Food struct {
	Event
	Id int
}
type Exercise struct {
	Event
	Id int
}
