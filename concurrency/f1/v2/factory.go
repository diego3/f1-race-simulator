package v2

import (
	"github.com/TwiN/go-color"
)

var compounds = createTyresCompounds()

func createDrivers() []*Driver {
	var drivers []*Driver

	hamilton := NewDriver("HAM", 997, 1250) // 253
	hamilton.Tyres = compounds[SOFT_COMPOUND]
	hamilton.PitStop1 = PitStop{Lap: 20, Tyre: compounds[MEDIUM_COMPOUND]}

	verstappen := NewDriver("VER", 998, 1245) // 247
	verstappen.Tyres = compounds[HARD_COMPOUND]
	verstappen.PitStop1 = PitStop{Lap: 35, Tyre: compounds[SOFT_COMPOUND]}

	drivers = append(drivers, hamilton)
	drivers = append(drivers, verstappen)
	drivers = append(drivers, NewDriver("LEC", 999, 1350))
	drivers = append(drivers, NewDriver("SAI", 1010, 1810))
	drivers = append(drivers, NewDriver("RUS", 1200, 1850))
	drivers = append(drivers, NewDriver("PER", 1000, 1550))
	drivers = append(drivers, NewDriver("NOR", 1110, 1660))
	drivers = append(drivers, NewDriver("RIC", 1055, 1750))
	drivers = append(drivers, NewDriver("ALO", 1255, 1500))
	drivers = append(drivers, NewDriver("OCO", 1200, 1450))
	drivers = append(drivers, NewDriver("BOT", 1200, 1500)) //300 var
	drivers = append(drivers, NewDriver("ZOU", 1300, 1650)) //350 var
	drivers = append(drivers, NewDriver("MAG", 1215, 1750))
	drivers = append(drivers, NewDriver("MIC", 1565, 1890))
	drivers = append(drivers, NewDriver("ALB", 1225, 1400))
	drivers = append(drivers, NewDriver("LAT", 1400, 1500))
	drivers = append(drivers, NewDriver("GAS", 1100, 1500))
	drivers = append(drivers, NewDriver("TIS", 1400, 1800))
	drivers = append(drivers, NewDriver("STR", 1544, 1900))
	drivers = append(drivers, NewDriver("HUL", 1600, 1800))

	return drivers
}

func createTyresCompounds() map[int]Tyre {
	tyreMap := make(map[int]Tyre)
	tyreMap[SOFT_COMPOUND] = Tyre{Compound: "Soft", Grip: 15, Age: 0, Color: color.Red}
	tyreMap[MEDIUM_COMPOUND] = Tyre{Compound: "Medium", Grip: 25, Age: 0, Color: color.Yellow}
	tyreMap[HARD_COMPOUND] = Tyre{Compound: "Hard", Grip: 50, Age: 0, Color: color.Bold}
	return tyreMap
}
