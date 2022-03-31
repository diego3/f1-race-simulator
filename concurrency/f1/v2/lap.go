package v2

import (
	"fmt"
)

type Lap struct {
	Lap         int
	Pos         string
	DriverName  string
	Tyre        string
	Diff        string
	Time        string
	CurrentTime string
}

func NewLap(lap, pos int, driver *Driver) Lap {
	return Lap{
		Lap:         lap,
		Pos:         fmt.Sprintf("%d", pos),
		DriverName:  driver.Name,
		Tyre:        fmt.Sprintf("%s(%d)", driver.Tyres.Compound, driver.Tyres.Age),
		Diff:        fmt.Sprintf("%.3f", driver.DiffTime.Seconds()),
		Time:        fmt.Sprintf("%.3f", driver.SumRaceTime.Seconds()),
		CurrentTime: fmt.Sprintf("%.3f", driver.CurrentLapTime.Seconds()),
	}
}

func (l *Lap) String() string {
	return fmt.Sprintf("%s %s\tTyre:%s\tDiff: %s\tTime: %s\tCurTime:%s\n",
		l.Pos,
		l.DriverName,
		l.Tyre,
		l.Diff,
		l.Time,
		l.CurrentTime,
	)
}
