package core

type Actor struct {
}

type GameObject interface {
	Initialize()
	Update()
	Render()
}

type Component interface {
	Update()
}
