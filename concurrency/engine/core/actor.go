package core

type Actor struct {
	Name       string
	Id         string
	Components map[string]Component
}

func NewActor(name string) *Actor {
	return &Actor{
		Name:       name,
		Components: map[string]Component{},
	}
}

func (a *Actor) AddComponent(component Component) {
	// ????
	a.Components[""] = component
}

func (a *Actor) GetComponent(name string) Component {
	v, exists := a.Components[name]
	if !exists {
		return nil
	}
	return v
}
