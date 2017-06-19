package config

import "time"

type Simulator struct {
	MinBloodSugar            float64
	BloodSugarLimitToEntreesGlycation float64
	GlycationEncrimentBy int
	FoodLoock time.Duration
	ExerciseLock time.Duration
}

func NewSimulatorConfig() Simulator{

	return Simulator{
		MinBloodSugar:80,
		BloodSugarLimitToEntreesGlycation:150,
		GlycationEncrimentBy:1,
		FoodLoock: time.Hour * 2,
		ExerciseLock:time.Hour,
	}
}