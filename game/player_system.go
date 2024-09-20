package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"tanks/data"
	components "tanks/world/component"
)

var fire bool
var reload bool

func TakePlayerAction(g *Game) {
	turnTaken := false

	players := g.WorldTags["players"]
	x := 0
	y := 0
	directionPlayer := ""
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
		directionPlayer = "up"
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
		directionPlayer = "down"
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
		directionPlayer = "left"
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
		directionPlayer = "right"
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		turnTaken = true
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if reload {
			fire = true
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyTab) {
		if !fire {
			reload = true
			fmt.Println("reload", reload)
		}
	}

	level := g.Map.CurrentLevel

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*components.Position)
		dir := result.Components[direction].(*components.Direction)

		index := level.GetIndexFromXY(pos.X+x, pos.Y+y)

		tile := level.Tiles[index]
		if tile.Blocked != true {
			switch directionPlayer {
			case "up":
				dir.Dir = 1
				break
			case "down":
				dir.Dir = 2
			case "left":
				dir.Dir = 3
			case "right":
				dir.Dir = 4
			}
			level.Tiles[level.GetIndexFromXY(pos.X, pos.Y)].Blocked = false

			pos.X += x
			pos.Y += y

			level.Tiles[index].Blocked = true
		} else if x != 0 || y != 0 {
			if level.Tiles[index].TileType != data.WALL {
				monsterPosition := components.Position{}
				fmt.Println("monster", monsterPosition)
			}
		}
	}

	if x != 0 || y != 0 || turnTaken {
		g.Turn = GetNextState(g.Turn)
		g.TurnCounter = 0
	}
}
