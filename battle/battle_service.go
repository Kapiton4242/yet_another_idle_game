package battle

import "yet_another_idle_game/creation"

type BattleService struct {
	battleRepo *battleRepo
}

func NewBattleService() *BattleService {
	return &BattleService{
		battleRepo: newBattleRepo(),
	}
}

func (battleService *BattleService) Get(id Id) *Battle {
	return battleService.battleRepo.GetBattle(id)
}

func (battleService *BattleService) Save(battle *Battle) (id Id, err error) {
	return battleService.battleRepo.SaveBattle(battle)
}

func (battleService *BattleService) GetBattles(id creation.Id) ([]*Battle, error) {
	battles := make([]*Battle, 0)

	for _, battle := range battleService.battleRepo.battles {
		for _, member := range battle.Members[PLAYER] {
			if member == id {
				battles = append(battles, battle)
			}
		}
	}

	return battles, nil
}
