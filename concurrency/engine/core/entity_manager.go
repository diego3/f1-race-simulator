package core

type EntityManager struct {
	Entities map[string]GameObject
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities: make(map[string]GameObject),
	}
}

func (a *EntityManager) AddEntity(entity GameObject, name string) {
	a.Entities[name] = entity
}

func (a *EntityManager) GetByName(name string) *GameObject {
	v, exists := a.Entities[name]
	if exists {
		return &v
	}
	return nil
}
