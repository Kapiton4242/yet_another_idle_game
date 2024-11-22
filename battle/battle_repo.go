package battle

import (
	"errors"
	"sync/atomic"
)

type battleRepo struct {
	battles     map[Id]*Battle
	idGenerator atomic.Int32
}

func newBattleRepo() *battleRepo {
	return &battleRepo{
		battles:     make(map[Id]*Battle),
		idGenerator: atomic.Int32{},
	}
}

func (battleRepo *battleRepo) SaveBattle(battle *Battle) (id Id, err error) {
	if battle.Id == 0 {
		newId := battleRepo.idGenerator.Add(1)
		battle.Id = Id(newId)
	} else {
		_, exist := battleRepo.battles[battle.Id]

		if !exist {
			return 0, errors.New("battle_not_found")
		}
	}

	battleRepo.battles[battle.Id] = battle

	return battle.Id, nil
}

func (battleRepo *battleRepo) GetBattle(id Id) *Battle {
	return battleRepo.battles[id]
}
