package dailycost

import (
	"fmt"
	"net/http"

	"github.com/etowett/datsimple/backend/data"
	"github.com/etowett/datsimple/backend/db"
	"github.com/etowett/datsimple/backend/logger"
	"github.com/etowett/datsimple/backend/services"
	"github.com/gin-gonic/gin"
)

func createDailyCost(
	dbManager *db.DBManager,
	costService services.DailyCostService,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		var costForm data.CreateDailyCostForm
		err := c.BindJSON(&costForm)
		if err != nil {
			msg := fmt.Sprintf("Error with the form given: %v", err)
			logger.Infof(msg)
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false, "message": msg,
			})
			return
		}
		logger.Infof("costForm: %+v", costForm)

		costData, err := costService.AddDailyCost(dbManager, &costForm)
		if err != nil {
			msg := fmt.Sprintf("Error creating daily cost [%+v]: %+v", costForm, err.Error())
			logger.Infof(msg)
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false, "message": msg,
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"success": true, "cost_data": costData})
		return
	}
}
