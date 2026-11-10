package runner

import (
	"company-report/internal/config"
	"company-report/internal/delivery"

	"github.com/gin-gonic/gin"
)

func Run() {
	env := config.GetEnv()

	config.SetupDB(&env.DbConfig)

	r := gin.Default()
	delivery.RouterSetup(r)
	r.Run(":8080")
}
