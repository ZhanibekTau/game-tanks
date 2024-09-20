package game

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"tanks/data"
	"tanks/world/board"
)

type Game struct {
	Map         board.GameMap
	World       *ecs.Manager
	WorldTags   map[string]ecs.Tag
	Turn        TurnState
	TurnCounter int
}

func NewGame(gameMap board.GameMap, world *ecs.Manager, tags map[string]ecs.Tag) *Game {
	g := &Game{
		Map:         gameMap,
		World:       world,
		WorldTags:   tags,
		Turn:        PlayerTurn,
		TurnCounter: 0,
	}
	return g
}

func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 18 {
		TakePlayerAction(g)
	}

	UpdateBullet(g)

	if g.Turn == EnemiesTurn {
		UpdateEnemy(g)
	}

	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessPlayerRenderer(g, level, screen)
	ProcessWeaponRenderer(g, level, screen)
	ProcessEnemyRenderer(g, level, screen)
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return data.TileWidth * data.ScreenWidth, data.TileHeight * data.ScreenHeight
}
