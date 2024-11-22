package monolith

import "yet_another_idle_game/creation"

type Monolith struct {
	Id                Id
	CreationId        creation.Id
	SoulEnergyPerTick int
	SoulEnergy        int
}

type Id int

type Field string

const (
	HP                Field = "HP"
	DamagePerHit      Field = "DamagePerHit"
	SoulEnergyPerTick Field = "SoulEnergyPerTick"
)

func (monolith *Monolith) UpgradeSoulEnergyPerTick() {
	monolith.SoulEnergyPerTick += 1
}

func (monolith *Monolith) IncreaseSoulEnergyByTick() {
	monolith.SoulEnergy += monolith.SoulEnergyPerTick
}

func (monolith *Monolith) ChangeSoulEnergy(delta int) {
	monolith.SoulEnergy += delta
}
