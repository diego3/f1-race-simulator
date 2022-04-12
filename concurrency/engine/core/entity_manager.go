package core

import (
	"log"
	"strings"
)

// "github.com/diego3/golang-handson/concurrency/engine/core"

type EntityManager struct {
	Entities map[string]Actor
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities: make(map[string]Actor),
	}
}

func (a *EntityManager) AddEntity(actor Actor) {
	if len(strings.TrimSpace(actor.Name)) == 0 {
		log.Println("Warn: Trying to add new actor without name")
		return
	}
	a.Entities[actor.Name] = actor
}

func (a *EntityManager) GetByName(name string) *Actor {
	v, exists := a.Entities[name]
	if exists {
		return &v
	}
	return nil
}
