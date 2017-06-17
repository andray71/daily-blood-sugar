package simulator

import "../config"
import "../database"
type Simulator struct {
	config.Simulator
	glycation int
	bloodSugar int
}

func NewSimulator(conf config.Simulator) Simulator {
	return Simulator{
		Simulator:conf,
		glycation:0,
		bloodSugar: conf.MinBloodSugar,
	}
}

func (s Simulator) Run(database database.Database) Simulator {

	if index,ok := database.GetFoodIndex(3); ok {
		println("in Run:", index)
	}
	return s
}
func (s Simulator)GetGlycation() int  {
	return s.glycation
}
func (s Simulator)GetBloodSugar() int  {
	return s.bloodSugar
}