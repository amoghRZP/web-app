package routers

import (
	"github.com/amoghRZP/web-app/handlers"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", handlers.GetHome)
	r.POST("/v1/notes/create", handlers.CreateNotes)
	r.GET("/v1/notes", handlers.GetNotes)
	r.GET("/v1/notes/:id", handlers.GetNoteById)
	r.PUT("/v1/notes/:id", handlers.UpdateNoteById)
	r.DELETE("/v1/notes/:id", handlers.DeleteNoteById)
	r.DELETE("v1/notes", handlers.DeleteNotes)

	return r;
}
