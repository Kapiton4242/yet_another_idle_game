package battle

import "yet_another_idle_game/creation"

type BattleInitializeService struct {
	battleService           *BattleService
	battleProcessingService *battleProcessingService
}

func NewBattleInitializeService(battleService *BattleService, creationService *creation.CreationService) *BattleInitializeService {
	return &BattleInitializeService{
		battleService:           battleService,
		battleProcessingService: newBattleProcessingService(creationService, battleService),
	}
}

func (battleInitializeService *BattleInitializeService) InitiateBattle(members map[FractionId][]creation.Id) (id Id, err error) {
	newBattle := NewBattle(members)

	id, err = battleInitializeService.battleService.Save(newBattle)

	if err != nil {
		return 0, err
	}

	go battleInitializeService.battleProcessingService.battleProcessing(newBattle)()

	return id, nil
}
