package interfaces

import (
	"net/http"

	"github.com/eureka/domain"
	"github.com/gin-gonic/gin"
)

type DieServiceInteractor interface {
	GetAllDie() ([]domain.Die, error)
}

type DieHandler struct {
	DieServiceInteractor DieServiceInteractor
}

func (handler *DieHandler) GetAllDie(c *gin.Context) {

	allDie, err := handler.DieServiceInteractor.GetAllDie()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": allDie})
}

type DieOrderServiceInteractor interface {
	CreateDieOrder(domain.DieOrderLine) error
}

type DieOrderHandler struct {
	DieOrderServiceInteractor DieOrderServiceInteractor
}

func (handler *DieOrderHandler) CreateDieOrder(c *gin.Context) {

	err := handler.DieOrderServiceInteractor.CreateDieOrder()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": allDie})
}
