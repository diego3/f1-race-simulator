package gamelogic

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/TwiN/go-color"
	"github.com/diego3/golang-handson/concurrency/engine/application"
	"github.com/diego3/golang-handson/concurrency/engine/core"
	"github.com/diego3/golang-handson/concurrency/f1/network"
)

var eventManager *application.EventManager
var entityManager *core.EntityManager

type Game struct {
	LapDuration int //seconds
	MaxLaps     int
	Drivers     []*Driver
	GameLogic   GameLogic

	NextWorkManager *network.NetworkManager

	// this is the main ideia
	GameObjects []core.Actor
}

func (g *Game) Boot() {
	// load configs from file...
	g.Drivers = createDrivers()
	g.LapDuration = 6
	g.GameLogic = GameLogic{Lap: 1}
	g.MaxLaps = 71

	entityManager = core.NewEntityManager()
	eventManager = application.NewEventManager()
	eventManager.Register(network.OnLapSimulatedListener, application.EVENT_DRIVER_LAP_SIMULATED)

	//for _, gameObject := range g.GameObjects {
	//	gameObject.Initialize(g)
	//}
}

func (g *Game) GameLoop() {
	for {
		g.processInput()
		g.update()
		g.render()

		time.Sleep(time.Duration(g.LapDuration) * time.Second)

		// todo trigger EVENT_RACE_FINISHED
		// deve ter um jeito disso não ficar aqui!!!
		if g.GameLogic.Lap == g.MaxLaps {
			fmt.Printf("\nPODIUM:\n")
			fmt.Printf("P1: %s\n", g.Drivers[0].Name)
			fmt.Printf("P2: %s\n", g.Drivers[1].Name)
			fmt.Printf("P3: %s\n", g.Drivers[2].Name)
			break
		}
	}
}

func (g *Game) processInput() {}

func (g *Game) update() {
	// for _, gameObject := range g.GameObjects {
	// 	gameObject.Update(g)
	// }

	for _, driver := range g.Drivers {
		driver.Update(g)
	}

	g.GameLogic.Update(g)
}

func (g *Game) render() {
	fmt.Println("RENDER")
	command := exec.Command("clear")
	command.Stdout = os.Stdout
	command.Run()

	// for _, gameObject := range g.GameObjects {
	// 	gameObject.Render(g)
	// }

	g.GameLogic.Lap++
	msg1 := fmt.Sprintf("Grand Prix: %s\tLap %d\n\n", "Mônaco", g.GameLogic.Lap)
	fmt.Print(color.InBold(msg1))

	g.GameLogic.Render(g)
}
