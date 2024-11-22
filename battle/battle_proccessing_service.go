package battle

import (
	"time"
	"yet_another_idle_game/creation"
)

type battleProcessingService struct {
	creationService *creation.CreationService
	battleService   *BattleService
}

func newBattleProcessingService(creationService *creation.CreationService, battleService *BattleService) *battleProcessingService {
	return &battleProcessingService{creationService: creationService, battleService: battleService}
}

func (battleProcessingService *battleProcessingService) battleProcessing(newBattle *Battle) func() {
	return func() {
		for {
			if battleProcessingService.battleRoundProcessing(newBattle) {
				return
			}

			time.Sleep(time.Second * 1)
		}
	}
}

func (battleProcessingService *battleProcessingService) battleRoundProcessing(newBattle *Battle) (battleEnd bool) {
	characterId := newBattle.Members[PLAYER][0]
	character := battleProcessingService.creationService.Get(characterId)

	enemyId := newBattle.Members[PLAYER][0]
	enemy := battleProcessingService.creationService.Get(enemyId)

	character.GetDamage(enemy.DamagePerHit)
	_, _ = battleProcessingService.creationService.Save(character)

	if !character.IsAlive() {
		newBattle.BattleStatus = LOSE
		_, _ = battleProcessingService.battleService.Save(newBattle)
		return true
	}

	enemy.GetDamage(character.DamagePerHit)
	_, _ = battleProcessingService.creationService.Save(enemy)

	if !enemy.IsAlive() {
		newBattle.BattleStatus = WIN
		_, _ = battleProcessingService.battleService.Save(newBattle)
		return true
	}

	return false
}
