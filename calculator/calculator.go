package calculator

import "fmt"

var sharpness_multipliers = map[string]float64{
	"r": 0.5,
	"o": 0.75,
	"y": 1.0,
	"g": 1.05,
	"b": 1.2,
	"w": 1.32,
}

type Modifier struct {
	Name              string
	Additional_attack int
}

var modifiers = []Modifier{
	{"Base damage", 0},
	{"+ attack up L", 15 + 20},
	{"+ powercharm", 15 + 20 + 6},
	{"+ powertalon", 15 + 20 + 6 + 9},
}

func average_damage(weapon_damage float64, affinity float64) float64 {
	return weapon_damage * (1 + 0.25*affinity)
}

func Damage_calculator(weapon_damage int, sharpness string, affinity int) {
	base_damage := sharpness_multipliers[sharpness] * float64(weapon_damage)
	standard_affinity := 0.01 * float64(affinity)
	weakpoint_affinity := standard_affinity + 0.5

	for _, v := range modifiers {
		modified_damage := base_damage + float64(v.Additional_attack)
		fmt.Printf(
			"%v\t%.2f\t(%.2f)\n",
			v.Name,
			average_damage(modified_damage, standard_affinity),
			average_damage(modified_damage, weakpoint_affinity),
		)
	}
}
