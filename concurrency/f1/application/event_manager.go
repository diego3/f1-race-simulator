package application

import (
	"fmt"
)

type EventManager struct {
	Queue0             []Event
	Queue1             []Event                           // to avoid infinite consumer
	CurrentActiveQueue int                               // we switch between queues because one event can trigger another events
	Consumers          map[EventType][]func(interface{}) // map {event: method}
}

func NewEventManager() *EventManager {
	return &EventManager{
		CurrentActiveQueue: 0,
		Consumers:          make(map[EventType][]func(interface{})),
	}
}

type Event struct {
	Type EventType
	Data interface{}
}

type EventType int

const (
	EVENT_A = 1
	EVENT_B = 2
)

func (e *EventManager) Register(consumer func(interface{}), eventType EventType) {
	consumerList := e.Consumers[eventType]
	consumerList = append(consumerList, consumer)
	e.Consumers[eventType] = consumerList
}

func (e *EventManager) Queue(event Event) {
	// switch between queues
	if e.CurrentActiveQueue == 0 {
		e.Queue0 = append(e.Queue0, event)
	}
	if e.CurrentActiveQueue == 1 {
		e.Queue1 = append(e.Queue1, event)
	}
}

func (e *EventManager) Update() {
	// call all the listeners/consumers
	fmt.Println(e.Consumers)
	fmt.Println("queue-0", e.Queue0)
	fmt.Println("queue-1", e.Queue1)

	// for each event: find consumers by eventType to call they registered func
	for _, evt := range e.Queue0 {
		listeners := e.Consumers[evt.Type]
		if len(listeners) == 0 {
			fmt.Println("Zero listeners for eventType", evt.Type)
			continue
		}
		for _, callbackFunc := range listeners {
			callbackFunc(evt.Data)
		}
	}
	e.CurrentActiveQueue = (e.CurrentActiveQueue + 1) % 2
}
