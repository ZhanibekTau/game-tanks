package game

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"tanks/data"
	"tanks/utils"
	"tanks/world/board"
	components "tanks/world/component"
)

var (
	position       *ecs.Component
	bulletPosition *ecs.Component
	renderable     *ecs.Component
	name           *ecs.Component
	direction      *ecs.Component
	weapon         *ecs.Component
	movable        *ecs.Component
	enemy          *ecs.Component
)

func InitializeWorld(startingLevel board.Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable = manager.NewComponent()
	name = manager.NewComponent()
	direction = manager.NewComponent()
	bulletPosition = manager.NewComponent()
	weapon = manager.NewComponent()
	enemy = manager.NewComponent()
	enemiesPositionX := utils.GetDiceRoll(19)
	enemiesPositionY := utils.GetDiceRoll(6)

	x, y := 13, 9
	d := 0
	player := CreateNewPlayerEntity(manager, x, y, d)
	weapon = CreateNewBulletEntity(manager, x, y, d)
	for {
		index := startingLevel.GetIndexFromXY(enemiesPositionX, enemiesPositionY)
		tileType := startingLevel.Tiles[index].TileType
		if tileType == data.FLOOR {
			enemy = CreateEnemy(manager, enemiesPositionX, enemiesPositionY, d)
			break
		} else {
			enemiesPositionX = utils.GetDiceRoll(19)
			enemiesPositionY = utils.GetDiceRoll(6)
		}
	}

	players := ecs.BuildTag(player, position, name, direction, renderable)
	tags["players"] = players

	renderables := ecs.BuildTag(renderable, position, direction, weapon)
	tags["renderables"] = renderables

	weapons := ecs.BuildTag(weapon, bulletPosition, name, direction, renderable)
	tags["weapons"] = weapons

	enemies := ecs.BuildTag(enemy, position, name, direction, renderable)
	tags["enemies"] = enemies

	return manager, tags
}

func CreateNewPlayerEntity(manager *ecs.Manager, x, y, d int) *ecs.Component {
	tags := make(map[string]ecs.Tag)
	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}
	player := manager.NewComponent()

	manager.NewEntity().
		AddComponent(player, components.Player{}).
		AddComponent(renderable, &components.Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, components.Movable{}).
		AddComponent(position, &components.Position{
			X: x,
			Y: y,
		}).
		AddComponent(direction, &components.Direction{
			Dir: d,
		}).
		AddComponent(name, &components.Name{Label: "Player"})

	players := ecs.BuildTag(player, position, name, direction, renderable)
	tags["players"] = players

	return player
}

func CreateNewBulletEntity(manager *ecs.Manager, x, y, d int) *ecs.Component {
	tags := make(map[string]ecs.Tag)
	weaponImg, _, err := ebitenutil.NewImageFromFile("assets/bullet.png")
	if err != nil {
		log.Fatal(err)
	}

	manager.NewEntity().
		AddComponent(weapon, &components.Weapon{
			Name:          "Bullet",
			MinimumDamage: 10,
			MaximumDamage: 20,
			ToHitBonus:    3,
			WasGunshot:    true,
		}).
		AddComponent(renderable, &components.Renderable{
			Image: weaponImg,
		}).
		AddComponent(movable, components.Movable{}).
		AddComponent(direction, &components.Direction{
			Dir: d,
		}).
		AddComponent(bulletPosition, &components.BulletPosition{
			X: float64(x),
			Y: float64(y),
		}).
		AddComponent(name, &components.Name{Label: "Weapons"})

	weapons := ecs.BuildTag(weapon, bulletPosition, name, direction, renderable)
	tags["weapons"] = weapons

	return weapon
}

func CreateEnemy(manager *ecs.Manager, x, y, d int) *ecs.Component {
	tags := make(map[string]ecs.Tag)
	enemyImg, _, err := ebitenutil.NewImageFromFile("assets/enemy.png")
	if err != nil {
		log.Fatal(err)
	}

	manager.NewEntity().
		AddComponent(enemy, &components.Enemy{}).
		AddComponent(renderable, &components.Renderable{
			Image: enemyImg,
		}).
		AddComponent(movable, components.Movable{}).
		AddComponent(direction, &components.Direction{
			Dir: d,
		}).
		AddComponent(position, &components.Position{
			X: x,
			Y: y,
		}).
		AddComponent(name, &components.Name{Label: "Enemy"})

	enemies := ecs.BuildTag(weapon, position, name, direction, renderable)
	tags["enemies"] = enemies

	return enemy
}
