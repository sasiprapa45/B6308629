package controller

import (
	"github.com/sasiprapa45/sa-65-exa/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

//--------------------------------------------------------------------------------------------------------

func CreateDoctor(c *gin.Context) {

	var doctor entity.Doctor

	if err := c.ShouldBindJSON(&doctor); err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}

	if err := entity.DB().Create(&doctor).Error; err != nil {
		   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		   return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}


//--------------------------------------------------------------------------------------------------------


func GetDoctor(c *gin.Context) {

	var doctor entity.Doctor

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM doctors WHERE id = ?", id).Scan(&doctor).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": doctor})
}

// GET /doctors
func ListDoctors(c *gin.Context) {

	var doctors []entity.Doctor

	if err := entity.DB().Raw("SELECT * FROM doctors").Scan(&doctors).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	c.JSON(http.StatusOK, gin.H{"data": doctors})
}

