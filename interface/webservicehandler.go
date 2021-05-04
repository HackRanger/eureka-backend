package interfaces

import (
	"log"
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
	CreateDieOrder([]domain.DieOrderLine) error
	GetAllDieOrders() ([]domain.DieOrderLine, error)
}

type DieOrderHandler struct {
	DieOrderServiceInteractor DieOrderServiceInteractor
}

type Orders struct {
	AllOrderLines []domain.DieOrderLine `json:"orders" binding:"required"`
}

func (handler *DieOrderHandler) CreateDieOrder(c *gin.Context) {
	var o Orders

	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.DieOrderServiceInteractor.CreateDieOrder(o.AllOrderLines)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (handler *DieOrderHandler) GetAllDieOrders(c *gin.Context) {
	allDieOrders, err := handler.DieOrderServiceInteractor.GetAllDieOrders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": allDieOrders})
}
