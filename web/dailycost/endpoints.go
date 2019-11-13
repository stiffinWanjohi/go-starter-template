package dailycost

import (
	"github.com/etowett/datsimple/backend/db"
	"github.com/etowett/datsimple/backend/services"
	"github.com/gin-gonic/gin"
)

func AddEndpoints(
	r *gin.RouterGroup,
	dbManager *db.DBManager,
	costService services.DailyCostService,
) {
	r.POST("add", createDailyCost(dbManager, costService))
}
