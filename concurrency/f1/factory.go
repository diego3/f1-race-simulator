package f1

import "github.com/TwiN/go-color"

const SOFT_COMPOUND = 0
const MEDIUM_COMPOUND = 1
const HARD_COMPOUND = 2

const MONACO_GRAND_PRIX = "MONACO"

func createDrivers() []*Driver {
	var drivers []*Driver

	hamilton := NewDriver("HAM", 997, 1250) // 253
	hamilton.Tyres = compounds[SOFT_COMPOUND]
	hamilton.PitStop1 = PitStop{Lap: 20, Tyre: compounds[MEDIUM_COMPOUND]}

	verstappen := NewDriver("VER", 998, 1245) // 247
	verstappen.Tyres = compounds[SOFT_COMPOUND]
	verstappen.PitStop1 = PitStop{Lap: 15, Tyre: compounds[SOFT_COMPOUND]}
	verstappen.PitStop2 = PitStop{Lap: 30, Tyre: compounds[MEDIUM_COMPOUND]}

	leclerc := NewDriver("LEC", 999, 1350)
	leclerc.Tyres = compounds[SOFT_COMPOUND]
	leclerc.PitStop1 = PitStop{Lap: -1, Tyre: compounds[SOFT_COMPOUND]}

	drivers = append(drivers, hamilton)
	drivers = append(drivers, verstappen)
	drivers = append(drivers, leclerc)
	drivers = append(drivers, NewDriver("SAI", 1010, 1810))
	drivers = append(drivers, NewDriver("RUS", 1200, 1850))
	drivers = append(drivers, NewDriver("PER", 1100, 1650))
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
	tyreMap[SOFT_COMPOUND] = Tyre{Counpond: "Soft", Grip: 30, Age: 0, Color: color.Red}
	tyreMap[MEDIUM_COMPOUND] = Tyre{Counpond: "Medium", Grip: 50, Age: 0, Color: color.Yellow}
	tyreMap[HARD_COMPOUND] = Tyre{Counpond: "Hard", Grip: 100, Age: 0, Color: color.Bold}
	return tyreMap
}

func createGrandPrix(name string) *GrandPrix {
	prixMap := make(map[string]*GrandPrix)

	var conditions []WeatherCondition
	cond1 := WeatherCondition{Temperature: random(19, 40), ChangeLap: 0}
	cond2 := WeatherCondition{Temperature: random(15, 30), ChangeLap: random(30, 71)}
	conditions = append(conditions, cond1)
	conditions = append(conditions, cond2)

	monaco := &GrandPrix{Name: "Mônaco", Laps: 71, WeatherConditions: conditions}

	prixMap[MONACO_GRAND_PRIX] = monaco

	return prixMap[name]
}
