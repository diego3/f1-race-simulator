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
)

func init() {
	rand.Seed(time.Now().Unix())
}

var compounds = createTyresCompounds()

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

		fmt.Printf("Grand Prix: %s\tLap %d\n", grandPrixName, lap)

		times := simulateLapTime(lap, drivers)
		// sum current time position for each drive in this lap
		for _, driverLap := range times {
			currentDur := driverLap.time
			duration := mapSum[driverLap.driver.Name]
			mapSum[driverLap.driver.Name] = duration + currentDur
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
				driverLap.driver.Tyres.Counpond,
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
			pitStopTime = random(22500, 25000) // 2,5 sec to 5 sec
			driver.Tyres = driver.PitStop1.Tyre
		}
		if driver.PitStop2.Lap == currentLap {
			pitStopTime = random(22500, 25000) // 2,5 sec to 5 sec
			driver.Tyres = driver.PitStop2.Tyre
		}
		if driver.Tyres.Grip > 0 {
			tyreTime = driver.Tyres.Age * driver.Tyres.Grip
		}
		randTime := random(driver.Min+tyreTime+pitStopTime, driver.Max+tyreTime+pitStopTime)
		randDuration, _ := time.ParseDuration(fmt.Sprintf("%dms", randTime))
		// todo: simulate driver error. Example: rand % 13 = 0
		times = append(times, DriverTime{driver: driver, time: randDuration})
	}

	// fastest driver in this lap
	sort.Slice(times, func(i, j int) bool {
		return times[i].time.Microseconds() < times[j].time.Microseconds()
	})

	return times
}

func createDrivers() []*Driver {
	var drivers []*Driver

	hamilton := NewDriver("HAM", 997, 1250) // 253
	hamilton.Tyres = compounds[0]
	hamilton.PitStop1 = PitStop{Lap: 20, Tyre: compounds[1]}

	verstappen := NewDriver("VER", 998, 1245) // 247
	verstappen.Tyres = compounds[0]
	verstappen.PitStop1 = PitStop{Lap: 15, Tyre: compounds[0]}
	verstappen.PitStop2 = PitStop{Lap: 30, Tyre: compounds[1]}

	drivers = append(drivers, hamilton)
	drivers = append(drivers, verstappen)
	drivers = append(drivers, NewDriver("LEC", 999, 1700))
	drivers = append(drivers, NewDriver("SAI", 1010, 1810))
	drivers = append(drivers, NewDriver("RUS", 1200, 1850))
	drivers = append(drivers, NewDriver("PER", 1100, 1650))
	drivers = append(drivers, NewDriver("NOR", 1110, 1660))
	drivers = append(drivers, NewDriver("RIC", 1055, 1750))
	drivers = append(drivers, NewDriver("ALO", 1255, 1500))
	drivers = append(drivers, NewDriver("OCO", 1200, 1450))
	drivers = append(drivers, NewDriver("BOT", 1200, 1500)) //300 var
	drivers = append(drivers, NewDriver("ZOU", 1300, 1650)) //350 var
	drivers = append(drivers, NewDriver("MAG", 1215, 1750))
	drivers = append(drivers, NewDriver("MIC", 1565, 1890))
	drivers = append(drivers, NewDriver("ALB", 1225, 1400))
	drivers = append(drivers, NewDriver("LAT", 1400, 1500))
	drivers = append(drivers, NewDriver("GAS", 1100, 1500))
	drivers = append(drivers, NewDriver("TIS", 1400, 1800))
	drivers = append(drivers, NewDriver("STR", 1544, 1900))
	drivers = append(drivers, NewDriver("HUL", 1600, 1800))

	return drivers
}

func getDriverByName(drivers []*Driver, name string) (*Driver, error) {
	for _, driver := range drivers {
		if strings.Compare(driver.Name, name) == 0 {
			return driver, nil
		}
	}
	return nil, errors.New("driver not found")
}

func createTyresCompounds() map[int]Tyre {
	tyreMap := make(map[int]Tyre)
	tyreMap[0] = Tyre{Counpond: "Soft", Grip: 30, Age: 0}
	tyreMap[1] = Tyre{Counpond: "Medium", Grip: 50, Age: 0}
	tyreMap[2] = Tyre{Counpond: "Hard", Grip: 100, Age: 0}
	return tyreMap
}

type DriverTime struct {
	driver *Driver
	time   time.Duration
}

type DriverSum struct {
	driver *Driver
	time   time.Duration
}

type LapManager struct {
	Laps []DriverTime
}

func (lm *LapManager) Add(laps []DriverTime) {

}

type Driver struct {
	Name     string
	Min      int
	Max      int
	Tyres    Tyre
	PitStop1 PitStop
	PitStop2 PitStop
	PitStop3 PitStop
	Sum      time.Duration
}

func NewDriver(name string, min, max int) *Driver {
	return &Driver{
		Name:     name,
		Min:      min,
		Max:      max,
		Tyres:    compounds[random(1, 3)-1],
		PitStop1: PitStop{Lap: random(20, 30), Tyre: compounds[random(1, 3)-1]},
	}
}

type Tyre struct {
	Counpond string
	Grip     int
	Age      int
}

type PitStop struct {
	Lap  int
	Tyre Tyre
}

func random(min, max int) int {
	return rand.Intn(max-min+1) + min
}
