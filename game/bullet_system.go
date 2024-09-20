package game

import (
	"tanks/data"
	components "tanks/world/component"
)

func UpdateBullet(g *Game) {
	for _, player := range g.World.Query(g.WorldTags["players"]) {
		posPlayer := player.Components[position].(*components.Position)
		dirPLayer := player.Components[direction].(*components.Direction)
		if fire {
			Fire(g, posPlayer, dirPLayer)
		}
	}

	//for _, enemyBullet := range g.World.Query(g.WorldTags["enemies"]) {
	//	enemyPos := enemyBullet.Components[position].(*components.Position)
	//	enemyDIr := enemyBullet.Components[direction].(*components.Direction)
	//	if fireEnemy {
	//		Fire(g, enemyPos, enemyDIr)
	//	}
	//}

	if len(g.World.Query(g.WorldTags["weapons"])) == 0 && reload {
		for _, player := range g.World.Query(g.WorldTags["players"]) {
			posPlayer := player.Components[position].(*components.Position)
			dirPLayer := player.Components[direction].(*components.Direction)
			CreateNewBulletEntity(g.World, posPlayer.X, posPlayer.Y, dirPLayer.Dir)
			reload = false
		}
	}
}

func Fire(g *Game, posPlayer *components.Position, dirPLayer *components.Direction) {
	level := g.Map.CurrentLevel
	for _, bullet := range g.World.Query(g.WorldTags["weapons"]) {
		pos := bullet.Components[bulletPosition].(*components.BulletPosition)
		wpn := bullet.Components[weapon].(*components.Weapon)
		dir := bullet.Components[direction].(*components.Direction)

		index := level.GetIndexFromXY(int(posPlayer.X), int(posPlayer.Y))
		tile := level.Tiles[index]
		if wpn.WasGunshot {
			pos.Y = float64(posPlayer.Y)
			pos.X = float64(posPlayer.X)
			dir.Dir = dirPLayer.Dir
			wpn.WasGunshot = false
		}

		if tile.TileType != data.WALL {
			switch dir.Dir {
			case 1:
				pos.Y = pos.Y - 0.3
			case 2:
				pos.Y = pos.Y + 0.3
			case 3:
				pos.X = pos.X - 0.3
			case 4:
				pos.X = pos.X + 0.3
			}

			if level.Tiles[level.GetIndexFromXY(int(pos.X), int(pos.Y))].TileType == data.WALL || pos.X == 0 && pos.Y == 0 {
				level.Tiles[level.GetIndexFromXY(int(pos.X), int(pos.Y))].TileType = data.FLOOR
				level.Tiles[level.GetIndexFromXY(int(pos.X), int(pos.Y))].Image = level.GetFloorImage()
				level.Tiles[level.GetIndexFromXY(int(pos.X), int(pos.Y))].Blocked = false

				g.World.DisposeEntity(bullet.Entity)
				fire = false
			} else if level.Tiles[level.GetIndexFromXY(int(pos.X), int(pos.Y))].TileType == data.CEMENT {
				g.World.DisposeEntity(bullet.Entity)
				fire = false
			}
		}
	}
}
