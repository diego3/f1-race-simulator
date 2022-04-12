package core

import "fmt"

type Transform2DComp struct {
	ActorComponent
	Position Vec2
}

func NewTransform2D(x, y int) Transform2DComp {
	return Transform2DComp{
		Position: Vec2{x: x, y: y},
	}
}

func (r Transform2DComp) Initialize() {
	fmt.Println("Transform2DComp::Initialize")
}

func (r Transform2DComp) Update() {
	fmt.Println("Transform2DComp::Update")
}
