package categories

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Categories(c *gin.Context) {
	var categories []models.Category
	models.DB.Find(&categories)
	c.IndentedJSON(http.StatusOK, categories)
}

func Create(c *gin.Context) {
	var newCategory models.Category

	if err := c.BindJSON(&newCategory); err != nil {
		return
	}
	models.DB.Create(&newCategory)
	c.IndentedJSON(http.StatusCreated, newCategory)
}

func Update(c *gin.Context) {
	id := c.Param(":id")
	var foundCategory models.Category
	err := models.DB.First(&foundCategory, id).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	var updatedCategory models.Category
	if err := c.BindJSON(&updatedCategory); err != nil {
		return
	}

	models.DB.Model(&foundCategory).Updates(&updatedCategory)
	c.IndentedJSON(http.StatusCreated, updatedCategory)
}

func Delete(c *gin.Context) {
	id := c.Param(":id")
	var foundCategory models.Category
	err := models.DB.First(&foundCategory, id).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	models.DB.Delete(&foundCategory)
	c.IndentedJSON(http.StatusNoContent, nil)
}
