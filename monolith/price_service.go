package monolith

type PriceService struct {
	monolithService *MonolithService
}

func NewPriceService(monolithService *MonolithService) *PriceService {
	return &PriceService{
		monolithService: monolithService,
	}
}

func (priceService *PriceService) GetPrice(monolith *Monolith, field Field) int {
	return 10
}
