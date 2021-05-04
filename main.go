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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	log.Println("Starting: Eureka Application")
	r := gin.Default()

	r.Use(CORSMiddleware())

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

	doi := new(usecases.DieOrderService)
	doi.DieOrderRepo = interfaces.NewDbDieOrder(db)
	dieOrderHadler := interfaces.DieOrderHandler{}
	dieOrderHadler.DieOrderServiceInteractor = doi
	dieMgmt.POST("/dieOrder", dieOrderHadler.CreateDieOrder)

	r.Run("0.0.0.0:8080")

	log.Println("Shutdown: Eureka Application")
}
