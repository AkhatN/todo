package routes

import (
	"net/http"
	"rest/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//Handlers ...

// Getlist reads gets todo list
func Getlist(c *gin.Context) {
	tds, err := models.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot show todo list"})
		return
	}

	c.JSON(http.StatusOK, tds)
}

// Createitem creates item
func Createitem(c *gin.Context) {
	newitem := &models.Todo{}

	if err := c.BindJSON(newitem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	newitem.Created_at = time.Now().Format("2006-02-01 15:04:05")

	if err := newitem.PostItem(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "item was not created"})
		return
	}

	c.JSON(http.StatusCreated, newitem)
}

// Updatelist updates specific item from the list
func Updatelist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	item := &models.Todo{}
	item.ID = id
	if err := item.UpdateItem(&id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

//Deleteitem deletes item from the list
func Deleteitem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	if err := models.DeleteItem(&id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item was deleted"})
}
