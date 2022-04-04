package application

import (
	"fmt"
	"testing"
)

type DamageEventData struct {
	Name   string
	Damage int
}

type FakeCar struct {
	damage int
}

func (f *FakeCar) TakeDamage(eventData interface{}) {
	damageObjt := eventData.(DamageEventData)
	fmt.Println("car take damage", damageObjt)
	f.damage += damageObjt.Damage
}

func TestEventManager(t *testing.T) {
	manager := NewEventManager()
	event := Event{
		Type: EVENT_A,
		Data: DamageEventData{
			Name:   "Max Verstappen",
			Damage: 20,
		},
	}
	car := FakeCar{damage: 0}
	manager.Register(car.TakeDamage, EVENT_A)
	manager.Queue(event)
	manager.Update()
	if car.damage != 20 {
		t.Fatal("Failed, expected damage not match", car.damage)
	}
}

func TestEventManager2(t *testing.T) {
	manager := NewEventManager()
	event := Event{
		Type: EVENT_A,
		Data: DamageEventData{
			Name:   "Max Verstappen",
			Damage: 20,
		},
	}
	car := FakeCar{damage: 0}
	manager.Register(car.TakeDamage, EVENT_A)
	manager.Queue(event)
	manager.Update()
	if car.damage != 20 {
		t.Fatal("Failed, expected damage not match", car.damage)
	}
}
