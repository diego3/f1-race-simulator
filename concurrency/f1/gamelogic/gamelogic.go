package gamelogic

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

type GameLogic struct {
	Lap               int
	FastestLap        Driver
	Laps              []Lap
	FastestCurrentLap *Driver
	SlowestCurrentLap *Driver
}

func (g *GameLogic) Update(game *Game) {
	drivers := game.Drivers

	g.gridSort(drivers)
	g.checkFastestLap(drivers)
	g.checkFastestCurrentLap(drivers)

	// calculate Diff time between drivers
	for index, driver := range game.Drivers {
		if index == 0 {
			driver.PrevDriver = nil
			driver.DiffTime = time.Duration(0)
		} else {
			prevIndex := index - 1
			maxIndex := len(game.Drivers) - 1
			if index >= maxIndex {
				prevIndex = maxIndex
			}
			driver.PrevDriver = game.Drivers[prevIndex]

			driver.DiffTime = driver.SumRaceTime - driver.PrevDriver.SumRaceTime
		}
	}
}

func (g *GameLogic) Render(game *Game) {
	var grid []string
	for pos, driverLap := range game.Drivers {
		lap := NewLap(g.Lap, pos+1, driverLap)
		g.Laps = append(g.Laps, lap)

		go WriteJson("data/laps.json", g.Laps)

		curTimeLabel := "CurTime"
		if strings.Compare(g.FastestCurrentLap.Name, driverLap.Name) == 0 {
			curTimeLabel = color.InGreen("CurTime")
		}
		if strings.Compare(g.SlowestCurrentLap.Name, driverLap.Name) == 0 {
			curTimeLabel = color.InRed("CurTime")
		}
		line := fmt.Sprintf("%d %s\tTyre:%s(%d)\tDiff: %.3f\tTime: %.3f\t%s: %.3f\n",
			pos+1,
			driverLap.Name,
			color.Colorize(driverLap.Tyres.Color, driverLap.Tyres.Compound),
			driverLap.Tyres.Age,
			driverLap.DiffTime.Seconds(),
			driverLap.SumRaceTime.Seconds(),
			curTimeLabel,
			driverLap.CurrentLapTime.Seconds())
		grid = append(grid, line)
	}
	for _, line := range grid {
		fmt.Print(line)
	}
	fmt.Printf("\nFastest Lap: %s Time: %.3f\n", g.FastestLap.Name, g.FastestLap.CurrentLapTime.Seconds())
}

func (g *GameLogic) checkFastestLap(drivers []*Driver) {
	var driversCopy []*Driver
	for _, d := range drivers {
		driversCopy = append(driversCopy, &Driver{Name: d.Name, CurrentLapTime: d.CurrentLapTime})
	}

	sort.Slice(driversCopy, func(i, j int) bool {
		return driversCopy[i].CurrentLapTime.Microseconds() < driversCopy[j].CurrentLapTime.Microseconds()
	})

	fastest := driversCopy[0]
	if g.FastestLap.CurrentLapTime == 0 {
		g.FastestLap = *fastest
	}

	if fastest.CurrentLapTime < g.FastestLap.CurrentLapTime {
		g.FastestLap = *fastest
	}
}

func (g *GameLogic) checkFastestCurrentLap(drivers []*Driver) {
	var driversCopy []*Driver
	for _, d := range drivers {
		driversCopy = append(driversCopy, &Driver{Name: d.Name, CurrentLapTime: d.CurrentLapTime})
	}

	sort.Slice(driversCopy, func(i, j int) bool {
		return driversCopy[i].CurrentLapTime.Microseconds() < driversCopy[j].CurrentLapTime.Microseconds()
	})

	g.FastestCurrentLap = driversCopy[0]
	g.SlowestCurrentLap = driversCopy[len(driversCopy)-1]
}

func (g *GameLogic) gridSort(drivers []*Driver) {
	// fastest driver the race
	sort.Slice(drivers, func(i, j int) bool {
		return drivers[i].SumRaceTime.Microseconds() < drivers[j].SumRaceTime.Microseconds()
	})
}
