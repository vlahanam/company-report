package delivery

import (
	usertrans "company-report/internal/module/user/transport"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouterSetup(db *gorm.DB, r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		usertrans.RegisterRoutes(v1, db)

		v1.GET("/posts", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": "v1", "data": "all posts"})
		})
	}
}

