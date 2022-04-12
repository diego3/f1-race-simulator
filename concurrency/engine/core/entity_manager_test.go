package core

import "testing"

func TestGetByNameFunction(t *testing.T) {
	manager := NewEntityManager()

	boss1 := RobotBoss{Name: "boss1"}
	manager.AddEntity(boss1, "boss1")

	boss1Get := manager.GetByName("boss1")
	if boss1Get == nil {
		t.Fatal("Excepted boss1 instance, but it wasnt found")
	}
	//manager.LoadFromJson("filename.json")
}

type RobotBoss struct {
	Name string
}

func (r RobotBoss) Initialize(game *Game) {

}

func (r RobotBoss) Update(game *Game) {

}

func (r RobotBoss) Render(game *Game) {

}
