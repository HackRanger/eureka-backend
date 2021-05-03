package main

import (
	"fmt"
	"log"

	"github.com/eureka/infra"
	interfaces "github.com/eureka/interface"
	"github.com/eureka/usecases"
	"github.com/gin-gonic/gin"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "welcome123"
	dbname   = "eureka"
)

func main() {
	log.Println("Starting: Eureka Application")
	r := gin.Default()
	// Init DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db := infra.NewDBHandler(psqlInfo)
	defer db.Close()
	dieMgmt := r.Group("/dieManagement")

	di := new(usecases.DieRegistryService)
	di.DieMasterRepo = interfaces.NewDbDie(db)
	webServiceHandler := interfaces.DieHandler{}
	webServiceHandler.DieServiceInteractor = di

	dieMgmt.GET("/", webServiceHandler.GetAllDie)

	r.Run("0.0.0.0:8080")
	log.Println("Shutdown: Eureka Application")
}
