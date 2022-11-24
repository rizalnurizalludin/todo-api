package routes

import (
	"github.com/rizalnurizalludin/todo-api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("", controllers.Home)
	r.GET("/activity-groups", controllers.FindActivities)
	r.POST("/activity-groups", controllers.CreateActivity)
	r.GET("/activity-groups/:id", controllers.FindActivity)
	r.PATCH("/activity-groups/:id", controllers.UpdateActivity)
	r.DELETE("activity-groups/:id", controllers.DeleteActivity)
	r.GET("/todo-items", controllers.FindTODOS)
	r.POST("/todo-items", controllers.CreateTODO)
	r.GET("/todo-items/:id", controllers.FindTODO)
	r.PATCH("/todo-items/:id", controllers.UpdateTODO)
	r.DELETE("todo-items/:id", controllers.DeleteTODO)
	return r
}
