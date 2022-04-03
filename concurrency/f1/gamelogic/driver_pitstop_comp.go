package gamelogic

type PitStopComponent struct {
	Driver *Driver
}

func (p *PitStopComponent) Update(game *Game) {
	if p.Driver.PitStop1.Lap == game.GameLogic.Lap {
		pitStopTime := Random(22500, 25000)
		p.Driver.MinCalculated += pitStopTime
		p.Driver.MaxCalculated += pitStopTime
		p.Driver.Tyres = p.Driver.PitStop1.Tyre
		return
	}

	if p.Driver.PitStop2.Lap == game.GameLogic.Lap {
		pitStopTime := Random(22500, 25000)
		p.Driver.MinCalculated += pitStopTime
		p.Driver.MaxCalculated += pitStopTime
		p.Driver.Tyres = p.Driver.PitStop2.Tyre
		return
	}

	p.Driver.MinCalculated += 0
	p.Driver.MaxCalculated += 0
}
