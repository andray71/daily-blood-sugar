package input

import (
	"testing"
	"fmt"
)

func TestReadCsv(t *testing.T){
	events := ReadCsv("../../testData/input1.csv")
	if len(events) != 5 {
		t.Fatal(fmt.Sprint("Expecting 5 events got", len(events)))
	}
}