package transport

import (
	"company-report/internal/module/user/business"
	"company-report/internal/module/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewCreateUserRepo(db)
	biz := business.NewCreateUserBiz(repo)

	users := rg.Group("/users")
	{
		users.POST("", CreateUser(biz))
	}
}
