package v2

type TyresComponent struct {
	Driver *Driver
}

func (t *TyresComponent) Update(game *Game) {
	var tyreTime int = 0

	t.Driver.Tyres.Age++
	if t.Driver.Tyres.Grip > 0 {
		tyreTime = t.Driver.Tyres.Age * t.Driver.Tyres.Grip
	}

	t.Driver.MaxCalculated = tyreTime
	t.Driver.MinCalculated = tyreTime
}
