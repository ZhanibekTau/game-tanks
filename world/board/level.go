package board

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"tanks/data"
	"tanks/utils"
	components "tanks/world/component"
)

var (
	floor  *ebiten.Image
	wall   *ebiten.Image
	cement *ebiten.Image
)

type Level struct {
	Tiles []*MapTile
	Rooms []components.Rect
}

func NewLevel() Level {
	l := Level{}
	loadTileImages()
	tiles := l.CreateTiles()
	l.Tiles = tiles
	l.GenerateLevelTiles()
	return l
}

type MapTile struct {
	PixelX     int
	PixelY     int
	Blocked    bool
	Image      *ebiten.Image
	IsRevealed bool
	TileType   data.TileType
}

func (level *Level) GetIndexFromXY(x int, y int) int {
	return (y * data.ScreenWidth) + x
}

func (level *Level) CreateTiles() []*MapTile {
	tiles := make([]*MapTile, data.LevelHeight*data.ScreenWidth)
	index := 0
	for x := 0; x < data.ScreenWidth; x++ {
		for y := 0; y < data.LevelHeight; y++ {
			if x == 0 || y == 0 || x == data.ScreenWidth-1 || y == data.LevelHeight-1 {
				index = level.GetIndexFromXY(x, y)
				tile := MapTile{
					PixelX:     x * data.TileWidth,
					PixelY:     y * data.TileHeight,
					Blocked:    true,
					Image:      cement,
					IsRevealed: false,
					TileType:   data.CEMENT,
				}
				tiles[index] = &tile
			} else {
				index = level.GetIndexFromXY(x, y)
				tile := MapTile{
					PixelX:     x * data.TileWidth,
					PixelY:     y * data.TileHeight,
					Blocked:    true,
					Image:      wall,
					IsRevealed: false,
					TileType:   data.WALL,
				}
				tiles[index] = &tile
			}
		}
	}
	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image) {
	for x := 0; x < data.ScreenWidth; x++ {
		for y := 0; y < data.LevelHeight; y++ {
			idx := level.GetIndexFromXY(x, y)
			tile := level.Tiles[idx]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
			level.Tiles[idx].IsRevealed = true
		}
	}
}

func (level *Level) createRoom(room components.Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = data.FLOOR
			level.Tiles[index].Image = floor
		}
	}
}

func (level *Level) GenerateLevelTiles() {
	var x, y int

	tiles := level.CreateTiles()
	level.Tiles = tiles
	tileX, tileY := 0, 0
	for _, blockType := range IntGrid {
		if blockType == 0 {
			x = tileX
			y = tileY
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = data.FLOOR
			level.Tiles[index].Image = floor
		}
		tileX++
		if tileX >= data.ScreenWidth {
			tileY++
			tileX = 0
		}
	}
}
func (level *Level) createHorizontalTunnel(x1 int, x2 int, y int) {
	for x := utils.Min(x1, x2); x < utils.Max(x1, x2)+1; x++ {
		index := level.GetIndexFromXY(x, y)
		if index > 0 && index < data.ScreenWidth*data.LevelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = data.FLOOR
			level.Tiles[index].Image = floor
		}
	}
}
func (level *Level) createVerticalTunnel(y1 int, y2 int, x int) {
	for y := utils.Min(y1, y2); y < utils.Max(y1, y2)+1; y++ {
		index := level.GetIndexFromXY(x, y)

		if index > 0 && index < data.ScreenWidth*data.LevelHeight {
			level.Tiles[index].Blocked = false
			level.Tiles[index].TileType = data.FLOOR
			level.Tiles[index].Image = floor
		}
	}
}

func (level Level) InBounds(x, y int) bool {
	if x < 0 || x > data.ScreenWidth || y < 0 || y > data.LevelHeight {
		return false
	}
	return true
}

func (level Level) IsOpaque(x, y int) bool {
	idx := level.GetIndexFromXY(x, y)
	return level.Tiles[idx].TileType == data.WALL || level.Tiles[idx].TileType == data.CEMENT
}

func loadTileImages() {
	if floor != nil && wall != nil {
		return
	}
	var err error

	floor, _, err = ebitenutil.NewImageFromFile("assets/floor.png")
	if err != nil {
		log.Fatal(err)
	}

	wall, _, err = ebitenutil.NewImageFromFile("assets/wall.png")
	if err != nil {
		log.Fatal(err)
	}

	cement, _, err = ebitenutil.NewImageFromFile("assets/cement.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (level *Level) GetFloorImage() *ebiten.Image {
	return floor
}
