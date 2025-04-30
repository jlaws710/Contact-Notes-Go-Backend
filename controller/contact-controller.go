package controller

import (
	"net/http"
	"notes-system/config"
	"notes-system/model"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	var contacts []model.Contact

	config.DB.Find(&contacts)
	c.JSON(http.StatusOK, contacts)
}

func CreateContact(c *gin.Context) {
	var contact model.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	config.DB.Create(&contact)
	c.JSON(http.StatusCreated, contact)
}

func UpdateContact(c *gin.Context) {
	id := c.Param("id")
	var contact model.Contact
	if err := config.DB.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})

		return
	}
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	config.DB.Save(&contact)
	c.JSON(http.StatusOK, contact)
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&model.Contact{}, id)
	c.Status(http.StatusNoContent)
}
