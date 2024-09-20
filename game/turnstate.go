package game

const (
	BeforePlayerAction = iota
	PlayerTurn
	EnemiesTurn
	GameOver
)

type TurnState int

func GetNextState(state TurnState) TurnState {
	switch state {
	case BeforePlayerAction:
		return PlayerTurn
	case PlayerTurn:
		return EnemiesTurn
	case EnemiesTurn:
		return BeforePlayerAction
	case GameOver:
		return GameOver
	default:
		return PlayerTurn
	}
}
