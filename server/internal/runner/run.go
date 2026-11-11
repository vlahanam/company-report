package runner

import (
	"company-report/internal/config"
	"company-report/internal/delivery"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	env := config.GetEnv()

	db := config.SetupDB(&env.DbConfig)

	r := gin.Default()
	delivery.RouterSetup(db, r)
	port := fmt.Sprintf(":%s", env.ServerConfig.Port)
	r.Run(port)
}
