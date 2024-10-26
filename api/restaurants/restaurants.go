package restaurants

import (
	"api/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var restaurants = []models.Restaurant{
//     {ID: "1", Name: "Blue Train", Address: "John Coltrane", Rating: 5},
//     {ID: "2", Name: "Jeru", Address: "Gerry Mulligan", Rating: 1},
//     {ID: "3", Name: "Sarah Vaughan and Clifford Brown", Address: "Sarah Vaughan", Rating: 3},
// }

func Restaurants(c *gin.Context) {
	var restaurants []models.Restaurant
	models.DB.Find(&restaurants)
	c.IndentedJSON(http.StatusOK, restaurants)
}

func Create(c *gin.Context) {
	var newRestaurant models.Restaurant

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newRestaurant); err != nil {
		return
	}
	models.DB.Create(&newRestaurant)
	c.IndentedJSON(http.StatusCreated, newRestaurant)
}

func Update(c *gin.Context) {
	id := c.Param(":id")
	var foundRestaurant models.Restaurant
	err := models.DB.First(&foundRestaurant, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle record not found error...
		c.IndentedJSON(http.StatusNotFound, errors.New("Not Found"))
	}

	var updatedRestaurant models.Restaurant
	if err := c.BindJSON(&updatedRestaurant); err != nil {
		return
	}

	models.DB.Model(&foundRestaurant).Updates(&updatedRestaurant)
	c.IndentedJSON(http.StatusCreated, updatedRestaurant)
}

func Delete(c *gin.Context) {
	id := c.Param(":id")
	var foundRestaurant models.Restaurant
	err := models.DB.First(&foundRestaurant, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle record not found error...
		c.IndentedJSON(http.StatusNotFound, errors.New("Not Found"))
	}

	var updatedRestaurant models.Restaurant
	if err := c.BindJSON(&updatedRestaurant); err != nil {
		return
	}

	models.DB.Delete(&foundRestaurant, id)
	c.Status(200)
}
