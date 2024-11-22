package battle

import (
	"yet_another_idle_game/creation"
)

type Battle struct {
	Id           Id
	BattleStatus BattleStatus
	Members      map[FractionId][]creation.Id
}

type Id int

type FractionId int

const (
	PLAYER FractionId = iota
	ENEMY
)

type BattleStatus string

const (
	PROCESSING BattleStatus = "processing"
	WIN        BattleStatus = "win"
	LOSE       BattleStatus = "lose"
)

func NewBattle(members map[FractionId][]creation.Id) *Battle {
	return &Battle{
		BattleStatus: PROCESSING,
		Members:      members,
	}
}
