package f1

import "time"

type DriverTime struct {
	Driver *Driver
	Time   time.Duration
}

type DriverSum struct {
	driver *Driver
	time   time.Duration
}

type LapManager struct {
	Laps []DriverTime
}

func (lm *LapManager) Add(laps []DriverTime) {

}

type Driver struct {
	Name     string
	Min      int
	Max      int
	Tyres    Tyre
	PitStop1 PitStop
	PitStop2 PitStop
	PitStop3 PitStop
	Sum      time.Duration
}

func NewDriver(name string, min, max int) *Driver {
	return &Driver{
		Name:     name,
		Min:      min,
		Max:      max,
		Tyres:    compounds[Random(1, 3)-1],
		PitStop1: PitStop{Lap: Random(20, 30), Tyre: compounds[Random(1, 3)-1]},
	}
}

type Tyre struct {
	Counpond string
	Grip     int
	Age      int
	Color    string
}

type PitStop struct {
	Lap  int
	Tyre Tyre
}

const DRY = 1
const WET = 2

type GrandPrixWeekend struct {
	Qualify1 Qualifying
	Qualify2 Qualifying
	Qualify3 Qualifying
	Race     Race
}

type Qualifying struct {
	Pass              int
	DriversPass       []Driver
	DriversNotPass    []Driver
	Sequence          int
	Weather           int
	WeatherConditions []WeatherCondition
}

type Race struct {
	Name              string
	Laps              int
	WeatherConditions []WeatherCondition
	Weather           int
}

type WeatherCondition struct {
	Temperature int
	Wind        int
	DryFactor   int // 0=clear 100=rain
	ChangeLap   int
}
