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
