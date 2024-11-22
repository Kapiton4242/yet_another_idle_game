package creation

type Creation struct {
	id           Id
	Name         string
	HP           int
	MaxHP        int
	DamagePerHit int
	AttackSpeed  int
}

type Id int

func (creation *Creation) UpgradeDamage() {
	creation.DamagePerHit += 1
}

func (creation *Creation) UpgradeHP() {
	creation.HP += 1
	creation.MaxHP += 1
}

func (creation *Creation) GetDamage(damage int) {
	creation.HP -= damage
}

func (creation *Creation) IsAlive() bool {
	return creation.HP > 0
}
