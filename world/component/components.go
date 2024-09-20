package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Health struct {
	MaxHealth     int
	CurrentHealth int
}

type Player struct {
}

type Enemy struct {
}

type Position struct {
	X int
	Y int
}

type BulletPosition struct {
	X float64
	Y float64
}

type Direction struct {
	Key ebiten.Key
	Dir int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}

type Weapon struct {
	Name          string
	MinimumDamage int
	MaximumDamage int
	ToHitBonus    int
	WasGunshot    bool
}

type Armor struct {
	Name       string
	Defense    int
	ArmorClass int
}

type Name struct {
	Label string
}

type UserMessage struct {
	AttackMessage    string
	DeadMessage      string
	GameStateMessage string
}

func (p *Position) GetManhattanDistance(other *Position) int {
	xDist := math.Abs(float64(p.X - other.X))
	yDist := math.Abs(float64(p.Y - other.Y))
	return int(xDist) + int(yDist)
}

func (p *Position) IsEqual(other *Position) bool {
	return (p.X == other.X && p.Y == other.Y)
}
