package board

type GameMap struct {
	BattleField  []BattleField
	CurrentLevel Level
}

func NewGameMap(l Level, battleFields []BattleField) GameMap {
	//Return a new game map of a single board for now
	return GameMap{
		BattleField:  battleFields,
		CurrentLevel: l,
	}
}
