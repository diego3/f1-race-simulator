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

type Car struct {
	Engine          Engine
	Tyres           []Tyre
	AeroConfig      AeroConfig
	CurrentVelocity int
}

func (c *Car) Velocity() {
	// use engine + tyres + aero
	c.CurrentVelocity = 0
}

type AeroConfig struct {
	Front int //higher more fast at streight sectors, lower more fast at turn sectors
	Rear  int
}

type Engine struct {
	ActorId     string
	Temperature int
	MinTemp     int
	MaxTemp     int
	Batery      Batery
	Life        int // 100 is healthy
}

type Batery struct {
	MaxCharge     int
	CurrentCharge int
	TimeToRecover int
	IsOn          bool
}

type Track struct {
	Sectors []TrackSector
}

const SECTOR_TYPE_TURN = 1
const SECTOR_TYPE_STREIGHT = 2

type TrackSector struct {
	Type      int
	IsDRSZone bool
}
