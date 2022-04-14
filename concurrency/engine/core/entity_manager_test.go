package core

import "testing"

func TestGetByNameFunction(t *testing.T) {
	manager := NewEntityManager()

	boss1 := NewActor("boss1")
	manager.AddEntity(*boss1)

	boss1Get := manager.GetByName("boss1")
	if boss1Get == nil {
		t.Fatal("Expected boss1 instance, but it wasnt found")
	}
	//manager.LoadFromJson("filename.json")
}

func TestNewActor(t *testing.T) {
	manager := NewEntityManager()

	boss1 := NewActor("   ")
	manager.AddEntity(*boss1)

	boss1Get := manager.GetByName("boss1")
	if boss1Get != nil {
		t.Fatal("Expected empty actor")
	}
}

func TestAddOneComponent(t *testing.T) {
	manager := NewEntityManager()

	boss1 := NewActor("boss1")
	transform := NewTransform2D(0, 0)
	boss1.AddComponent(transform)

	manager.AddEntity(*boss1)

	boss1Get := manager.GetByName("boss1")
	if boss1Get == nil {
		t.Fatal("Expected boss1 instance, but it wasnt found")
	}

	if len(boss1Get.Components) != 1 {
		t.Fatal("Expected 1 component, bug result wast = ", len(boss1.Components))
	}
}

func TestGetComponent(t *testing.T) {
	manager := NewEntityManager()

	boss1 := NewActor("boss1")
	transform := NewTransform2D(10, 20)
	boss1.AddComponent(transform)
	manager.AddEntity(*boss1)

	boss1Get := manager.GetByName("boss1")

	transform2d := boss1Get.GetComponent("Transform2D")
	if transform2d == nil {
		t.Fatal("Expected a transform component")
	}
	// this is how you can "cast" from the base "class"
	theRealTransform := transform2d.(Transform2DComp)
	if theRealTransform.Position.x != 10 {
		t.Fatal("Exception real transform objetc with Vector2.x = 10, got", theRealTransform.Position.x)
	}
	if theRealTransform.Position.y != 20 {
		t.Fatal("Exception real transform objetc with Vector2.y = 20, got", theRealTransform.Position.y)
	}
}
