package gamelogic

const SOFT_COMPOUND = 0
const MEDIUM_COMPOUND = 1
const HARD_COMPOUND = 2

type Tyre struct {
	Compound string
	Grip     int
	Age      int
	Color    string
}

type PitStop struct {
	Lap  int
	Tyre Tyre
}
