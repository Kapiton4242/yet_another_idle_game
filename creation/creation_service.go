package creation

type CreationService struct {
	creationRepo *creationRepo
}

func NewCreationService() *CreationService {
	return &CreationService{
		creationRepo: newCreationRepo(),
	}
}

func (creationService *CreationService) Get(id Id) *Creation {
	return creationService.creationRepo.GetCreation(id)
}

func (creationService *CreationService) Save(creation *Creation) (id Id, err error) {
	return creationService.creationRepo.SaveCreation(creation)
}
