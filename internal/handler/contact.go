package handler

import (
	"net/http"
	"strconv"

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

func ContactGetByIDHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	for _, contact := range contacts {
		if contact.ID == id {
			c.JSON(http.StatusOK, contact)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Kişi bulunamadı"})
}

func ContactUpdateHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	var updatedData model.Contact
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			contacts[i].Name = updatedData.Name
			contacts[i].Email = updatedData.Email
			contacts[i].Phone = updatedData.Phone

			c.JSON(http.StatusOK, contacts[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Kişi bulunamadı"})
}

func ContactDeleteHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Kişi silindi"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Kişi bulunamadı"})
}
