package game

import (
	"github.com/norendren/go-fov/fov"
	components "tanks/world/component"
	"tanks/world/component/path"
)

var fireEnemy bool

func UpdateEnemy(game *Game) {
	l := game.Map.CurrentLevel
	playerPosition := components.Position{}

	for _, plr := range game.World.Query(game.WorldTags["players"]) {
		pos := plr.Components[position].(*components.Position)
		playerPosition.X = pos.X
		playerPosition.Y = pos.Y
	}

	for _, result := range game.World.Query(game.WorldTags["enemies"]) {
		pos := result.Components[position].(*components.Position)
		//mon := result.Components[monster].(components.Monster)
		enemySees := fov.New()
		enemySees.Compute(l, pos.X, pos.Y, 8)
		if enemySees.IsVisible(playerPosition.X, playerPosition.Y) {
			astar := path.AStar{}
			pathMonster := astar.GetPath(l, pos, &playerPosition)
			if len(pathMonster) > 1 {
				nextTile := l.Tiles[l.GetIndexFromXY(pathMonster[1].X, pathMonster[1].Y)]
				if !nextTile.Blocked {
					l.Tiles[l.GetIndexFromXY(pos.X, pos.Y)].Blocked = false
					pos.X = pathMonster[1].X
					pos.Y = pathMonster[1].Y
					nextTile.Blocked = true
				}
			}
		}
	}
	game.Turn = PlayerTurn
}
