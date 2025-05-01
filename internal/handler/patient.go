package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sumitst05/patiently/internal/service"
	"github.com/sumitst05/patiently/utils"
)

func CreatePatient(c *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Gender  string `json:"gender"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	claims := c.MustGet("claims").(*utils.Claims)

	patient, err := service.CreatePatient(
		claims.UserId,
		claims.Role,
		req.Name,
		req.Age,
		req.Gender,
		req.Address,
		req.Phone,
	)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func GetAllPatients(c *gin.Context) {
	claims := c.MustGet("claims").(*utils.Claims)

	patients, err := service.GetAllPatients(claims.Role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func GetPatientById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	claims := c.MustGet("claims").(*utils.Claims)

	patient, err := service.GetPatientById(uint(id), claims.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

func GetPatientRegistrationHistory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	claims := c.MustGet("claims").(*utils.Claims)

	history, err := service.GetPatientRegistrationHistory(uint(id), claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve registration history"})
		return
	}

	c.JSON(http.StatusOK, history)
}

func UpdatePatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Gender  string `json:"gender"`
		Address string `json:"address"`
		Phone   string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	claims := c.MustGet("claims").(*utils.Claims)

	updatedPatient, err := service.UpdatePatient(
		uint(id),
		claims.Role,
		req.Name,
		req.Age,
		req.Gender,
		req.Address,
		req.Phone,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, updatedPatient)
}

func DeletePatient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	claims := c.MustGet("claims").(*utils.Claims)
	err = service.DeletePatient(uint(id), claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}
