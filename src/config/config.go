package config


type Simulator struct {
	MinBloodSugar            int
	BloodSugarLimitToEntreesGlycation int
	GlycationEncrimentBy int
}

func NewSimulatorConfig() Simulator{

	return Simulator{
		MinBloodSugar:80,
		BloodSugarLimitToEntreesGlycation:150,
		GlycationEncrimentBy:1,
	}
}