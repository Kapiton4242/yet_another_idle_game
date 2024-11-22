package monolith

import (
	"errors"
	"sync/atomic"
)

type monolithRepo struct {
	monoliths   map[Id]*Monolith
	idGenerator atomic.Int32
}

func newMonolithRepo() *monolithRepo {
	return &monolithRepo{
		monoliths:   make(map[Id]*Monolith),
		idGenerator: atomic.Int32{},
	}
}

func (monolithRepo *monolithRepo) SaveMonolith(monolith *Monolith) (id Id, err error) {
	if monolith.Id == 0 {
		newId := monolithRepo.idGenerator.Add(1)
		monolith.Id = Id(newId)
	} else {
		_, exist := monolithRepo.monoliths[monolith.Id]

		if !exist {
			return 0, errors.New("monolith_not_found")
		}
	}

	monolithRepo.monoliths[monolith.Id] = monolith

	return monolith.Id, nil
}

func (monolithRepo *monolithRepo) GetMonolith(id Id) *Monolith {
	return monolithRepo.monoliths[id]
}
