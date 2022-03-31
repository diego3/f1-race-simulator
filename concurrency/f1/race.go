package f1

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var compounds = createTyresCompounds()

/*
Game States:
	MENU
	Q1
	Q2
	Q3
	RACE

Race States:
 RUNNING
 SAFETY_CAR
 VIRTUAL_SAFETY_CAR
 YELLOW_FLAG

Allow send game commands:
    switch to plan B
	box box!

*/
func RaceStart() {
	drivers := createDrivers()
	mapSum := make(map[string]time.Duration)

	for _, driver := range drivers {
		mapSum[driver.Name] = time.Duration(0)
	}

	//start := time.Now()
	var lapDuration int = 6 //seconds
	var totalLaps int = 51
	var grandPrixName = "Mônaco"
	lap := 1
	//lapManager := LapManager{}
	for {
		command := exec.Command("clear")
		command.Stdout = os.Stdout
		command.Run()

		msg1 := fmt.Sprintf("Grand Prix: %s\tLap %d\n", grandPrixName, lap)
		fmt.Print(color.InBold(msg1))

		times := simulateLapTime(lap, drivers)
		go WriteJson(fmt.Sprintf("data/lap-%d.json", lap), times)

		// sum current time position for each drive in this lap
		for _, driverLap := range times {
			currentDur := driverLap.Time
			duration := mapSum[driverLap.Driver.Name]
			mapSum[driverLap.Driver.Name] = duration + currentDur
		}

		var driversSum []DriverSum
		for k, v := range mapSum {
			d, _ := getDriverByName(drivers, k)
			driversSum = append(driversSum, DriverSum{driver: d, time: v})
		}

		// fastest driver the race
		sort.Slice(driversSum, func(i, j int) bool {
			return driversSum[i].time.Microseconds() < driversSum[j].time.Microseconds()
		})

		// print current grid positions
		for pos, driverLap := range driversSum {
			var diff time.Duration
			if pos == 0 {
				diff = time.Duration(0)
			} else {
				// previous drive time
				index := pos - 1
				maxDriversLen := len(drivers) - 1
				if index >= maxDriversLen {
					index = maxDriversLen
				}
				previousDriverSum := driversSum[index].time
				diff = driverLap.time - previousDriverSum
			}
			fmt.Printf("%d %s\tTyre:%s(%d)\tDiff: %.3f\tTime: %.3f\n",
				pos+1,
				driverLap.driver.Name,
				color.Colorize(driverLap.driver.Tyres.Color, driverLap.driver.Tyres.Counpond),
				driverLap.driver.Tyres.Age,
				diff.Seconds(),
				driverLap.time.Seconds())
		}
		time.Sleep(time.Duration(lapDuration) * time.Second)
		//total = time.Since(start)

		if lap >= totalLaps {
			fmt.Printf("WINNER: %s\n", driversSum[0].driver.Name)
			break
		}
		lap++
	}
}

func simulateLapTime(currentLap int, drivers []*Driver) []DriverTime {
	var times []DriverTime
	for _, driver := range drivers {
		var tyreTime int = 0
		var pitStopTime int = 0

		driver.Tyres.Age++

		if driver.PitStop1.Lap == currentLap {
			pitStopTime = Random(22500, 25000) // 2,5 sec to 5 sec
			driver.Tyres = driver.PitStop1.Tyre
		}
		if driver.PitStop2.Lap == currentLap {
			pitStopTime = Random(22500, 25000) // 2,5 sec to 5 sec
			driver.Tyres = driver.PitStop2.Tyre
		}
		if driver.Tyres.Grip > 0 {
			tyreTime = driver.Tyres.Age * driver.Tyres.Grip
		}
		randTime := Random(driver.Min+tyreTime+pitStopTime, driver.Max+tyreTime+pitStopTime)
		randDuration, _ := time.ParseDuration(fmt.Sprintf("%dms", randTime))

		// todo: simulate driver error. Example: rand % 13 = 0
		// todo: simulate driver crash and safety car

		times = append(times, DriverTime{Driver: driver, Time: randDuration})
	}

	// fastest driver in this lap
	sort.Slice(times, func(i, j int) bool {
		return times[i].Time.Microseconds() < times[j].Time.Microseconds()
	})

	return times
}

func getDriverByName(drivers []*Driver, name string) (*Driver, error) {
	for _, driver := range drivers {
		if strings.Compare(driver.Name, name) == 0 {
			return driver, nil
		}
	}
	return nil, errors.New("driver not found")
}
