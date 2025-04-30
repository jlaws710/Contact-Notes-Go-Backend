package controller

import (
	"net/http"
	"notes-system/config"
	"notes-system/model"
	"notes-system/util"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	contactID := c.Param("id")
	var notes []model.Note

	config.DB.Where("contact_id = ?", contactID).Find(&notes)
	c.JSON(http.StatusOK, notes)
}

func CreateNote(c *gin.Context) {
	contactID := c.Param("id")
	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	input = util.NormalizeNoteInput(input)
	input["contact_id"] = contactID
	var note model.Note

	config.DB.Model(&model.Note{}).Create(&input)
	config.DB.Where("contact_id = ?", contactID).Last(&note)
	c.JSON(http.StatusCreated, note)
}
