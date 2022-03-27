package Handlers

import (
	"net/http"

	"gorm-db/Config"
	"gorm-db/Models"

	"github.com/gin-gonic/gin"
)

// In the Go GIN WebFramework, controllers are more like like Handlers;

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PEMs service is healthy!",
	})
}

func CreateVerse(c *gin.Context) {

	// initialize a VerseDAO object
	var input Models.VerseDAO

	// checks to determine if binding happened successfully; if it did not send back a "BadRequest" status to the client
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if binding happened successfully, use values from the VerseDAO object and assign them to the concrete Verse object to be saved to the database
	verse := Models.Verse{Book: input.Book, Chapter: input.Chapter, Verse: input.Verse, Text: input.Text, Translation_Name: input.Translation_Name, CategoryID: input.CategoryID}

	// save the verse object to the database
	Config.DB.Create(&verse)

	// let the client know everything is good :)
	c.JSON(http.StatusOK, gin.H{"data": verse})

}

func GetAllVerses(c *gin.Context) {

	// initialize an array of Verses
	var Verse []Models.Verse

	// "Find" gets all the items from a particular table
	result := Config.DB.Find(&Verse)
	println(result.RowsAffected) // returns count of records found
	println(result.Error)

	// let the client know everything is good and send back the array full of Verse
	c.JSON(http.StatusOK, Verse)

}

func CreateCategory(c *gin.Context) {

	// initialize a CategoryDAO object
	var input Models.CategoryDAO

	// checks to determine if binding happened successfully; if it did not send back a "BadRequest" status to the client
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if binding happened successfully, use values from the CategoryDAO object and assign them to the concrete Category object to be saved to the database
	category := Models.Category{Name: input.Name}

	// save the category object to the database
	Config.DB.Create(&category)

	// let the client know everything is good :)
	c.JSON(http.StatusOK, gin.H{"data": category})

}

func GetAllCategories(c *gin.Context) {
	var Category []Models.Category

	result := Config.DB.Find(&Category)
	println(result.RowsAffected) // returns count of records found
	println(result.Error)

	c.JSON(http.StatusOK, Category)
}

// func GetATodo(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var todo Models.Todo
// 	err := Models.GetATodo(&todo, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, todo)
// 	}
// }

// func UpdateATodo(c *gin.Context) {
// 	var todo Models.Todo
// 	id := c.Params.ByName("id")
// 	err := Models.GetATodo(&todo, id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, todo)
// 	}
// 	c.BindJSON(&todo)
// 	err = Models.UpdateATodo(&todo, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, todo)
// 	}
// }

// func DeleteATodo(c *gin.Context) {
// 	var todo Models.Todo
// 	id := c.Params.ByName("id")
// 	err := Models.DeleteATodo(&todo, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
// 	}
// }
