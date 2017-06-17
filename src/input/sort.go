package input

import "sort"

func Sort(in []Event) {
	sort.Slice(in, func(i, j int) bool { return in[i].GetTime().Before(in[j].GetTime()) })
}