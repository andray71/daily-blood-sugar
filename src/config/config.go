package config

import "time"

type Simulator struct {
	MinBloodSugar            int
	BloodSugarLimitToEntreesGlycation int
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