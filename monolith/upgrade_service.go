package monolith

import "yet_another_idle_game/creation"

type UpgradeService struct {
	monolithService *MonolithService
	priceService    *PriceService
	creationService *creation.CreationService
}

func NewUpgradeService(
	monolithService *MonolithService,
	priceService *PriceService,
	creationService *creation.CreationService,
) *UpgradeService {
	return &UpgradeService{
		monolithService: monolithService,
		priceService:    priceService,
		creationService: creationService,
	}
}

func (upgradeService *UpgradeService) UpgradeDamagePerHit(monolithId Id) {
	monolithToUpgrade := upgradeService.monolithService.Get(monolithId)
	creationToUpgrade := upgradeService.creationService.Get(monolithToUpgrade.CreationId)

	creationToUpgrade.UpgradeDamage()
	price := upgradeService.priceService.GetPrice(monolithToUpgrade, DamagePerHit)
	monolithToUpgrade.ChangeSoulEnergy(-price)
}

func (upgradeService *UpgradeService) UpgradeHP(monolithId Id) {
	monolithToUpgrade := upgradeService.monolithService.Get(monolithId)
	creationToUpgrade := upgradeService.creationService.Get(monolithToUpgrade.CreationId)

	creationToUpgrade.UpgradeHP()
	price := upgradeService.priceService.GetPrice(monolithToUpgrade, HP)
	monolithToUpgrade.ChangeSoulEnergy(-price)
}

func (upgradeService *UpgradeService) UpgradeSoulEnergyPerTick(monolithId Id) {
	monolithToUpgrade := upgradeService.monolithService.Get(monolithId)

	monolithToUpgrade.UpgradeSoulEnergyPerTick()
	price := upgradeService.priceService.GetPrice(monolithToUpgrade, SoulEnergyPerTick)
	monolithToUpgrade.ChangeSoulEnergy(-price)
}
