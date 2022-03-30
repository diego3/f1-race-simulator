package f1

import "time"

type DriverTime struct {
	driver *Driver
	time   time.Duration
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
		Tyres:    compounds[random(1, 3)-1],
		PitStop1: PitStop{Lap: random(20, 30), Tyre: compounds[random(1, 3)-1]},
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

type GrandPrix struct {
	Name              string
	Laps              int
	WeatherConditions []WeatherCondition
}

type WeatherCondition struct {
	Temperature int
	Wind        int
	DryFactor   int // 0=clear 100=rain
	ChangeLap   int
}
