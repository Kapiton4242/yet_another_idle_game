package monolith

type MonolithService struct {
	monolithRepo *monolithRepo
}

func NewMonolithService() *MonolithService {
	return &MonolithService{
		monolithRepo: newMonolithRepo(),
	}
}

func (monolithService *MonolithService) Get(id Id) *Monolith {
	return monolithService.monolithRepo.GetMonolith(id)
}

func (monolithService *MonolithService) Save(monolith *Monolith) (id Id, err error) {
	return monolithService.monolithRepo.SaveMonolith(monolith)
}
