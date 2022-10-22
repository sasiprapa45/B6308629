package main

import (
	"github.com/sasiprapa45/sa-65-exa/controller"

	"github.com/sasiprapa45/sa-65-exa/entity"

	"github.com/gin-gonic/gin"

)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
  
		  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  
		  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
  
		  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
  
		  c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
  
   
  
		  if c.Request.Method == "OPTIONS" {
  
			c.AbortWithStatus(204)
  
			return
  
		  }
  
   
  
		  c.Next()
  
	}
  
  }

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/patients", controller.ListPatients)
	r.GET("/patients/:id", controller.GetPatient)
	r.POST("/cpatients", controller.CreatePatient)

	r.GET("/doctors", controller.ListDoctors)
	r.GET("/doctors/:id", controller.GetDoctor)
	r.POST("/doctor", controller.CreateDoctor)
	
	r.GET("/symptoms", controller.ListSymptoms)
	r.GET("/symptoms/:id", controller.GetSymptoms)

	r.GET("/patient_types", controller.ListPatient_type)
	r.GET("/patient_types/:id", controller.GetPatientType)

	// Run the server

	r.Run()

}
