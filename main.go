package main

import (
	"github.com/szczynk/Sesi8/config"
	"github.com/szczynk/Sesi8/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := controllers.NewPersonInDB(db)

	router := gin.Default()

	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)

	router.Run(":3000")
}
