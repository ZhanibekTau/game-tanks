package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"tanks/data"
	"tanks/world/board"
	components "tanks/world/component"
)

func ProcessPlayerRenderer(g *Game, level board.Level, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["players"]) {
		pos := result.Components[position].(*components.Position)
		img := result.Components[renderable].(*components.Renderable).Image
		dir := result.Components[direction].(*components.Direction)

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]

		op := &ebiten.DrawImageOptions{}
		w, h := img.Bounds().Dx(), img.Bounds().Dy()
		op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
		switch dir.Dir {
		case 1:
			op.GeoM.Rotate(0)
		case 2:
			op.GeoM.Rotate(math.Pi)
		case 3:
			op.GeoM.Rotate(-math.Pi / 2)
		case 4:
			op.GeoM.Rotate(math.Pi / 2)
		}

		op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)

		op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(img, op)
	}
}

func ProcessWeaponRenderer(g *Game, level board.Level, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["weapons"]) {
		if fire {
			dirWpn := result.Components[direction].(*components.Direction)
			posWpn := result.Components[bulletPosition].(*components.BulletPosition)

			imgWpn := result.Components[renderable].(*components.Renderable).Image
			w, h := imgWpn.Bounds().Dx(), imgWpn.Bounds().Dy()

			indexWpn := level.GetIndexFromXY(int(posWpn.X), int(posWpn.Y))
			tileWpn := level.Tiles[indexWpn]
			opWpn := &ebiten.DrawImageOptions{}

			tileCenterX := float64(tileWpn.PixelX) + float64(data.TileWidth)/2.0
			tileCenterY := float64(tileWpn.PixelY) + float64(data.TileHeight)/2.0

			opWpn.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
			switch dirWpn.Dir {
			case 1:
				opWpn.GeoM.Rotate(0)
			case 2:
				opWpn.GeoM.Rotate(math.Pi)
			case 3:
				opWpn.GeoM.Rotate(-math.Pi / 2)
			case 4:
				opWpn.GeoM.Rotate(math.Pi / 2)
			}

			opWpn.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)

			opWpn.GeoM.Translate(tileCenterX-float64(w)/2.0, tileCenterY-float64(h)/2.0)
			screen.DrawImage(imgWpn, opWpn)
		}
	}
}

func ProcessEnemyRenderer(g *Game, level board.Level, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["enemies"]) {
		pos := result.Components[position].(*components.Position)
		img := result.Components[renderable].(*components.Renderable).Image
		dir := result.Components[direction].(*components.Direction)

		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]

		op := &ebiten.DrawImageOptions{}
		w, h := img.Bounds().Dx(), img.Bounds().Dy()
		op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
		switch dir.Dir {
		case 1:
			op.GeoM.Rotate(0)
		case 2:
			op.GeoM.Rotate(math.Pi)
		case 3:
			op.GeoM.Rotate(-math.Pi / 2)
		case 4:
			op.GeoM.Rotate(math.Pi / 2)
		}

		op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)

		op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(img, op)
	}
}
