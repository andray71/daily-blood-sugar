package input

import "sort"

type events []Event

func (s events) Len() int {
	return len(s)
}
func (s events) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s events) Less(i, j int) bool {
	return s[i].Time.Before(s[j].Time)
}

func Sort(in []Event)  {
	sort.Sort(events(in))
}