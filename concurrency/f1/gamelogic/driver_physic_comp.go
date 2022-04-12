package gamelogic

import (
	"fmt"
	"time"
)

type PhysicsComponent struct {
	Driver *Driver
}

func (p *PhysicsComponent) Update() {
	// todo: simulate driver crash and safety car

	newMin := p.Driver.Min + p.Driver.MinCalculated
	newMax := p.Driver.Max + p.Driver.MaxCalculated
	randTime := Random(newMin, newMax)
	randDuration, _ := time.ParseDuration(fmt.Sprintf("%dms", randTime))
	p.Driver.CurrentLapTime = randDuration
	p.Driver.SumRaceTime = p.Driver.SumRaceTime + p.Driver.CurrentLapTime
}
