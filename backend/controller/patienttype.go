package controller

import (
	"github.com/sasiprapa45/sa-65-exa/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

//------------------------------------------------------------------------------
// GET /patientType

func ListPatient_type(c *gin.Context) {

	var patienttype []entity.PatientType

	if err := entity.DB().Raw("SELECT * FROM patient_types").Scan(&patienttype).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": patienttype})
}

func GetPatientType(c *gin.Context) {

	var patienttype entity.PatientType

	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&patienttype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patienttype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patienttype})

}
