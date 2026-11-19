package transport

import (
	"context"
	"net/http"

	"company-report/internal/common"
	"company-report/internal/module/user/dto"
	"company-report/internal/module/user/entity"

	"github.com/gin-gonic/gin"
)

type CreateUserBusiness interface {
	CreateUser(ctx context.Context, req *dto.CreateUserReqDTO) (*entity.User, error)
}

func CreateUser(biz CreateUserBusiness) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateUserReqDTO

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCodeBadRequest)
			return
		}

		if err := req.Validate(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.ErrCodeValidation.WithDetailValidation(err))
			return
		}

		user, err := biz.CreateUser(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, common.NewCreatedResponse("user", "Created user successfully").WithData(user))
	}
}
