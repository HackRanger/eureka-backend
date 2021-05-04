package usecases

import "github.com/eureka/domain"

type DieMasterRepository interface {
	GetAllDie() ([]domain.Die, error)
}

type DieRegistryService struct {
	DieMasterRepo domain.DieMasterRepository
}

func (service *DieRegistryService) GetAllDie() ([]domain.Die, error) {
	return service.DieMasterRepo.GetAllDie()
}

type DieOrderRepository interface {
	CreateDieOrder([]domain.DieOrderLine) error
	GetAllDieOrders() ([]domain.DieOrderLine, error)
	GenerateLotNumber() (int, error)
}

type DieOrderService struct {
	DieOrderRepo domain.DieOrderRepository
}

func (service *DieOrderService) CreateDieOrder(orders []domain.DieOrderLine) error {
	return service.DieOrderRepo.CreateDieOrder(orders)
}

func (service *DieOrderService) GetAllDieOrders() ([]domain.DieOrderLine, error) {
	return service.DieOrderRepo.GetAllDieOrders()
}

func (service *DieOrderService) GenerateLotNumber() (int, error) {
	lotNo, err := service.DieOrderRepo.GenerateLotNumber()
	if err != nil {
		return lotNo, err
	}
	return lotNo + 1, err
}
