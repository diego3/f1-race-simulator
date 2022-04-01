package v2

type GameObject interface {
	Initialize(game *Game)
	Update(game *Game)
	Render(game *Game)
}

type Component interface {
	Update(game *Game)
}
