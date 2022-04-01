package v2

import (
	"time"
)

type Driver struct {
	Name           string
	Min            int
	Max            int
	Components     []Component
	CurrentLapTime time.Duration
	DiffTime       time.Duration // diff time from ahead driver
	SumRaceTime    time.Duration
	Tyres          Tyre
	PitStop1       PitStop
	PitStop2       PitStop
	PitStop3       PitStop
	PrevDriver     *Driver
	MinCalculated  int
	MaxCalculated  int
}

func NewDriver(name string, min, max int) *Driver {
	driver := &Driver{
		Name:     name,
		Min:      min,
		Max:      max,
		Tyres:    compounds[Random(1, 3)-1],
		PitStop1: PitStop{Lap: Random(20, 30), Tyre: compounds[Random(1, 3)-1]},
	}
	var components []Component
	physics := PhysicsComponent{
		Driver: driver,
	}
	pits := PitStopComponent{
		Driver: driver,
	}
	tyresComp := TyresComponent{
		Driver: driver,
	}

	components = append(components, &tyresComp)
	components = append(components, &pits)
	components = append(components, &physics)

	driver.Components = components
	return driver
}

func (d *Driver) Update(game *Game) {
	for _, component := range d.Components {
		component.Update(game)
	}
}
