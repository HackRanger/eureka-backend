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
