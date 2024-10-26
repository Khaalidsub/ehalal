package vendors

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Vendors(c *gin.Context) {
	category_id := c.Query("category_id")
	var vendors []models.Vendor

	if category_id != "" {
		models.DB.Where("category_id = ?", category_id).Find(&vendors)
		c.IndentedJSON(http.StatusOK, vendors)
		return
	}

	models.DB.Find(&vendors)
	c.IndentedJSON(http.StatusOK, vendors)
}

func Create(c *gin.Context) {
	var newVendor models.Vendor

	if err := c.BindJSON(&newVendor); err != nil {
		return
	}
	models.DB.Create(&newVendor)
	c.IndentedJSON(http.StatusCreated, newVendor)
}

func Update(c *gin.Context) {
	id := c.Param(":id")
	var foundVendor models.Vendor
	err := models.DB.First(&foundVendor, id).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	var updatedVendor models.Vendor
	if err := c.BindJSON(&updatedVendor); err != nil {
		return
	}

	models.DB.Model(&foundVendor).Updates(&updatedVendor)
	c.IndentedJSON(http.StatusCreated, updatedVendor)
}

func Delete(c *gin.Context) {
	id := c.Param(":id")
	var foundVendor models.Vendor
	err := models.DB.First(&foundVendor, id).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	models.DB.Delete(&foundVendor)
	c.IndentedJSON(http.StatusNoContent, nil)
}
