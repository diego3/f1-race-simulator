package gamelogic

type GameObject interface {
	Initialize(game *Game)
	Update(game *Game)
	Render(game *Game)
}

type Component interface {
	Update(game *Game)
}

// todo: move this guys to its files
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
