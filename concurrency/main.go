package main

import (
	"math/rand"
	"time"

	"github.com/diego3/golang-handson/concurrency/f1/gamelogic"
	"github.com/diego3/golang-handson/concurrency/f1/network"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	game := gamelogic.Game{}
	game.Boot()
	go game.GameLoop()

	network.WebServer()
}
