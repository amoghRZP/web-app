package handlers

import (
	"fmt"
	"github.com/amoghRZP/web-app/dbconfig"
	"github.com/amoghRZP/web-app/models"
	"github.com/gin-gonic/gin"
	"time"
)

func GetHome(c *gin.Context) {
	c.String(200, "Welcome to Notes !!")
}

func CreateNotes(c *gin.Context) {

	var notes models.Notes
	c.BindJSON(&notes)
	t := time.Now()
	notes.CreatedAt = t.Format(time.UnixDate)
	notes.UpdatedAt = t.Format(time.UnixDate)

	dbconfig.DB.Create(&notes)
	c.JSON(200, notes)
}

func GetNotes(c *gin.Context) {
	var notes []models.Notes
	if err := dbconfig.DB.Find(&notes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, notes)
	}
}

func GetNoteById(c *gin.Context) {

	id := c.Params.ByName("id")
	var notes models.Notes
	if err := dbconfig.DB.Where("id = ?", id).First(&notes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, notes)
	}
}

func UpdateNoteById(c *gin.Context) {

	id := c.Params.ByName("id")
	var notes models.Notes
	if err := dbconfig.DB.Where("id = ?", id).First(&notes).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.BindJSON(&notes)
	}
	t := time.Now()
	notes.UpdatedAt = t.Format(time.UnixDate)
	dbconfig.DB.Save(&notes)

	c.JSON(200, notes)

}

func DeleteNoteById(c *gin.Context) {
	id := c.Params.ByName("id")
	var notes models.Notes
	d := dbconfig.DB.Where("id = ?", id).Delete(&notes)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func DeleteNotes(c *gin.Context) {
	var notes []models.Notes
	d := dbconfig.DB.Delete(&notes)
	fmt.Println(d)

	c.JSON(200, gin.H{"Records": "deleted"})
}
