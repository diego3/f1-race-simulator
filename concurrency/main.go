package main

import (
	"math/rand"
	"time"

	v2 "github.com/diego3/golang-handson/concurrency/f1/v2"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	game := v2.Game{}
	game.Boot()
	game.GameLoop()
}
