package csv
import "testing"

func DB4Test() *textCsv  {
	return NewDb("../../../testData/FoodDB.csv","../../../testData/Exercise.csv")
}

func TestCSV(t *testing.T) {
	database := DB4Test()
	if(len(database.exercise) != 6){
		t.Fatal("Expected 6 records read from Exercise.csv csv file. Found",len(database.exercise))
	}
	if(len(database.food) != 112){
		t.Fatal("Expected 112 records read from FoodDB csv file. Found",len(database.food))
	}
	if exerciseIndex,ok := database.GetExerciseIndex(3); ok{
		if exerciseIndex != 40 {
			t.Error("expected index 40. Found",exerciseIndex)
		}
	}
	if foodIndex,ok := database.GetFoodIndex(3); ok{
		if foodIndex != 46 {
			t.Error("expected index 46. Found",foodIndex)
		}
	}
}
