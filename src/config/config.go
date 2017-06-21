package config

import "time"

type Simulator struct {
	MinBloodSugar            float64
	BloodSugarLimitToEncrimentGlycation float64
	GlycationEncrimentBy int
	FoodLock time.Duration
	ExerciseLock time.Duration
}

func NewSimulatorConfig() Simulator{

	return Simulator{
		MinBloodSugar:80,
		BloodSugarLimitToEncrimentGlycation:150,
		GlycationEncrimentBy:1,
		FoodLock: time.Hour * 2,
		ExerciseLock:time.Hour,
	}
}