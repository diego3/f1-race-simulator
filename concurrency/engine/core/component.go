package core

import "fmt"

type ActorComponent struct {
	Component
	Id    string
	Name  string
	Actor *Actor
}

type Component interface {
	Initialize()
	Update()
	GetName() string
}

func (ac ActorComponent) Initialize() {
	fmt.Println("ActorComponent::Initialize")
}

func (ac ActorComponent) Update() {
	fmt.Println("ActorComponent::Update")
}
