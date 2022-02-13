package app

import (
	"net/http"
	"time"

	"github.com/chalfel/go-lead-store/models"
	"github.com/gin-gonic/gin"
)

func HeathCheck(c *gin.Context) {
	heathCheck := &models.HeathCheckStatus{Status: "ok", Time: time.Now()}
	c.JSON(http.StatusOK, heathCheck)
}
