package main

import (
	"math/rand"
	"time"

	"github.com/diego3/golang-handson/concurrency/f1/gamelogic"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	game := gamelogic.Game{}
	game.Boot()
	game.GameLoop()
}
