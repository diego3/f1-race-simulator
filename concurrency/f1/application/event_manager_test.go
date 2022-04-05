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

func TestEventManagerSimpleEvent(t *testing.T) {
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
	if len(manager.Queue0) > 0 {
		t.Fatal("Failed, queue should be empty", len(manager.Queue0))
	}
}

func TestEventManagerWithRepeatedEvents(t *testing.T) {
	manager := NewEventManager()
	event1 := Event{
		Type: EVENT_A,
		Data: DamageEventData{
			Name:   "Max Verstappen",
			Damage: 20,
		},
	}
	car := FakeCar{damage: 0}
	manager.Register(car.TakeDamage, EVENT_A)

	manager.Queue(event1)
	manager.Queue(event1)

	manager.Update()
	if car.damage != 40 {
		t.Fatal("Failed, expected damage not match", car.damage)
	}
	if len(manager.Queue0) > 0 {
		t.Fatal("Failed, queue should be empty", len(manager.Queue0))
	}
}

func TestEventManagerForcingQueueCircuit(t *testing.T) {
	manager := NewEventManager()
	event1 := Event{
		Type: EVENT_A,
		Data: DamageEventData{
			Name:   "Event1",
			Damage: 10,
		},
	}

	var damageTotal int = 0
	onDamageListener := func(event interface{}) {
		damageEvt := event.(DamageEventData)

		damageTotal += damageEvt.Damage
		event2 := Event{
			Type: EVENT_A,
			Data: DamageEventData{
				Name:   "Listener-Event",
				Damage: 10,
			},
		}
		manager.Queue(event2)
	}

	manager.Register(onDamageListener, EVENT_A)

	manager.Queue(event1)

	manager.Update() // process queue0
	manager.Update() // process queue1
	manager.Update() // process queue0

	if damageTotal != 30 {
		t.Fatal("Damage Total expected = 30, got = ", damageTotal)
	}
}
