package data

const (
	ScreenWidth  = 20
	ScreenHeight = 15
	TileWidth    = 24
	TileHeight   = 24
	//UIHeight     = 10
	LevelHeight          = 15
	AccelerationConstant = 1
)

type TileType int

const (
	WALL TileType = iota
	FLOOR
	CEMENT
)
