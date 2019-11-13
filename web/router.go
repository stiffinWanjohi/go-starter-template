package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/etowett/datsimple/backend/db"
	"github.com/etowett/datsimple/backend/logger"
	"github.com/etowett/datsimple/backend/services"
	"github.com/etowett/datsimple/backend/web/dailycost"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	*gin.Engine
}

func (router *AppRouter) Run() {
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "9000"
	}
	logger.Infof("datsimple starting, listening on :%v", appPort)
	router.Engine.Run(fmt.Sprintf("0.0.0.0:%s", appPort))
}

func BuildRouter(
	dbManager *db.DBManager,
	costService services.DailyCostService,
) *AppRouter {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	addHealthEndpoints(router)

	dailyCostRoutes := router.Group("/daily-cost")
	{
		dailycost.AddEndpoints(dailyCostRoutes, dbManager, costService)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Endpoint Not Found"})
	})
	return &AppRouter{router}
}

func addHealthEndpoints(r *gin.Engine) {
	r.HEAD("/health", ping())
	r.GET("/health", ping())
}

func ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	}
}
