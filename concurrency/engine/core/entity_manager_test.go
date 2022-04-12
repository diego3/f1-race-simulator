package core

import "testing"

func TestGetByNameFunction(t *testing.T) {
	manager := NewEntityManager()

	boss1 := NewActor("boss1")
	manager.AddEntity(*boss1)

	boss1Get := manager.GetByName("boss1")
	if boss1Get == nil {
		t.Fatal("Excepted boss1 instance, but it wasnt found")
	}
	//manager.LoadFromJson("filename.json")
}

func TestNewActor(t *testing.T) {
	manager := NewEntityManager()

	boss1 := NewActor("   ")
	manager.AddEntity(*boss1)

	boss1Get := manager.GetByName("boss1")
	if boss1Get != nil {
		t.Fatal("Excepted empty actor")
	}
}

func TestAddComponent(t *testing.T) {
	manager := NewEntityManager()

	boss1 := NewActor("boss1")
	transform := NewTransform2D(0, 0)
	boss1.AddComponent(transform)

	manager.AddEntity(*boss1)

	boss1Get := manager.GetByName("boss1")
	if boss1Get == nil {
		t.Fatal("Excepted boss1 instance, but it wasnt found")
	}

	if len(boss1Get.Components) != 1 {
		t.Fatal("Expected 1 component, bug result wast = ", len(boss1.Components))
	}
}
