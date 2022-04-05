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

func (e *EventManager) getActiveQueue() []Event {
	// switch between queues
	if e.CurrentActiveQueue == 0 {
		return e.Queue0
	}
	return e.Queue1
}

func (e *EventManager) clearQueue(queueNumber int) {
	var empty []Event
	if queueNumber == 0 {
		e.Queue0 = empty
	}
	if queueNumber == 1 {
		e.Queue1 = empty
	}
}

func (e *EventManager) advanceNextQueue() {
	e.CurrentActiveQueue = (e.CurrentActiveQueue + 1) % 2
}

func (e *EventManager) Update() {
	activeQueue := e.getActiveQueue()
	lastActiveQueue := e.CurrentActiveQueue
	e.advanceNextQueue()

	// for each event: find consumers by eventType to call they registered func
	for _, evt := range activeQueue {
		listeners := e.Consumers[evt.Type]
		if len(listeners) == 0 {
			fmt.Println("Zero listeners for eventType", evt.Type)
			continue
		}
		for _, callbackFunc := range listeners {
			callbackFunc(evt.Data)
		}
	}

	e.clearQueue(lastActiveQueue)
}
