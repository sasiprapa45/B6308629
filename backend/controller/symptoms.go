package controller

import (
	"github.com/sasiprapa45/sa-65-exa/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

//------------------------------------------------------------------------------------------------------------------

// GET /symptoms

func ListSymptoms(c *gin.Context) {

	var symptoms []entity.Symptoms

	if err := entity.DB().Raw("SELECT * FROM symptoms").Scan(&symptoms).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": symptoms})
}
func GetSymptoms(c *gin.Context) {

	var symptoms entity.Symptoms
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&symptoms); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Symptoms not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": symptoms})

}
