package handler

import (
	"net/http"

	"github.com/ardaeu/go-contacts-api/internal/model"
	"github.com/gin-gonic/gin"
)

var contacts []model.Contact
var currentID int64 = 1

func ContactCreateHandler(c *gin.Context) {
	var newContact model.Contact

	if err := c.ShouldBindJSON(&newContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	newContact.ID = currentID
	currentID++
	contacts = append(contacts, newContact)

	c.JSON(http.StatusCreated, newContact)
}

func ContactListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, contacts)
}
