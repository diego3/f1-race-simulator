package v2

import (
	"fmt"
	"time"
)

type PhysicsComponent struct {
	Driver *Driver
}

func (p *PhysicsComponent) Update(game *Game) {
	var tyreTime int = 0
	var pitStopTime int = 0

	p.Driver.Tyres.Age++

	if p.Driver.PitStop1.Lap == game.GameLogic.Lap {
		pitStopTime = Random(22500, 25000) // 2,5 sec to 5 sec
		p.Driver.Tyres = p.Driver.PitStop1.Tyre
	}
	if p.Driver.PitStop2.Lap == game.GameLogic.Lap {
		pitStopTime = Random(22500, 25000) // 2,5 sec to 5 sec
		p.Driver.Tyres = p.Driver.PitStop2.Tyre
	}
	if p.Driver.Tyres.Grip > 0 {
		tyreTime = p.Driver.Tyres.Age * p.Driver.Tyres.Grip
	}

	// todo: simulate driver error. Example: rand % 13 = 0
	// todo: simulate driver crash and safety car

	newMin := p.Driver.Min + tyreTime + pitStopTime
	newMax := p.Driver.Max + tyreTime + pitStopTime
	randTime := Random(newMin, newMax)
	randDuration, _ := time.ParseDuration(fmt.Sprintf("%dms", randTime))
	p.Driver.CurrentLapTime = randDuration
	p.Driver.SumRaceTime = p.Driver.SumRaceTime + p.Driver.CurrentLapTime
}
