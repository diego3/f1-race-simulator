package gamelogic

import (
	"time"

	"github.com/diego3/golang-handson/concurrency/engine/core"
)

type Driver struct {
	Name           string
	Min            int
	Max            int
	Components     []core.Component
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

type DriverAttributes struct {
	Aggresive    int
	Pace         int //more pace more lap time consistency
	ExperienceXp int //how more xp more pace at race, less prob error
	Skills       int //more greater less chances to make mistakes
}

func NewDriver(name string, min, max int) *Driver {
	driver := &Driver{
		Name:     name,
		Min:      min,
		Max:      max,
		Tyres:    compounds[Random(1, 3)-1],
		PitStop1: PitStop{Lap: Random(20, 30), Tyre: compounds[Random(1, 3)-1]},
	}
	var components []core.Component
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
		component.Update()
	}
}

// draft: possible events
func (d *Driver) OnRaceStarts() {

}

func (d *Driver) OnRaceEnds() {

}

func (d *Driver) OnSafetyCarsEnter() {

}

func (d *Driver) OnSafetyCarsExit() {

}

func (d *Driver) OnVirtalSafetyCarStarts() {

}

func (d *Driver) OnVirtalSafetyCarEnds() {

}

func (d *Driver) OnReceivePitCommand(data interface{}) {

}
