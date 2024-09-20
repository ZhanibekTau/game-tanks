package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"tanks/game"
	"tanks/world/board"
)

func main() {
	l := board.NewLevel()
	levels := make([]board.Level, 0)
	levels = append(levels, l)
	d := board.BattleField{Name: "default", Levels: levels}
	battleFields := make([]board.BattleField, 0)
	battleFields = append(battleFields, d)

	gameMap := board.NewGameMap(l, battleFields)
	escWorld, tags := game.InitializeWorld(gameMap.CurrentLevel)

	g := game.NewGame(gameMap, escWorld, tags)
	ebiten.SetWindowResizingMode(2)
	ebiten.SetWindowTitle("Tanks")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
