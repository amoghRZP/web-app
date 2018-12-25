package routers

import (
	"github.com/amoghRZP/web-app/handlers"
	"github.com/amoghRZP/web-app/middlewares"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	//CUSTOM MIDDLEWARE BASIC AUTH
	r.Use(middlewares.LoginMiddleware)

	// BUILT IN AUTH MIDDLEWARE OF GIN
	/*authorized := r.Group("", gin.BasicAuth(gin.Accounts{
		"amogh": "amogh",
		"admin": "admin",
	}))*/

	r.POST("/v1/notes/create", handlers.CreateNotes)
	r.GET("/v1/notes/", handlers.GetNotes)
	r.GET("/v1/notes/:id", handlers.GetNoteById)
	r.PATCH("/v1/notes/:id", handlers.UpdateNoteById)
	r.DELETE("/v1/notes/:id", handlers.DeleteNoteById)
	r.DELETE("/v1/notes/", handlers.DeleteNotes)

	return r
}
