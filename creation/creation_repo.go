package creation

import (
	"errors"
	"sync/atomic"
)

type creationRepo struct {
	creations   map[Id]*Creation
	idGenerator atomic.Int32
}

func newCreationRepo() *creationRepo {
	return &creationRepo{
		creations:   make(map[Id]*Creation),
		idGenerator: atomic.Int32{},
	}
}

func (creationRepo *creationRepo) SaveCreation(creation *Creation) (id Id, err error) {
	if creation.id == 0 {
		newId := creationRepo.idGenerator.Add(1)
		creation.id = Id(newId)
	} else {
		_, exist := creationRepo.creations[creation.id]

		if !exist {
			return 0, errors.New("creation_not_found")
		}
	}

	creationRepo.creations[creation.id] = creation

	return creation.id, nil
}

func (creationRepo *creationRepo) GetCreation(id Id) *Creation {
	return creationRepo.creations[id]
}
