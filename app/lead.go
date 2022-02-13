package app

import (
	"net/http"

	"github.com/chalfel/go-lead-store/models"
	"github.com/chalfel/go-lead-store/repositories"
	"github.com/chalfel/go-lead-store/usecases"
	"github.com/gin-gonic/gin"
)

func GetAllLeads(c *gin.Context, server *Server) {
	usecase := usecases.GetAllLeadsUseCase{Repo: &repositories.MongoLeadRepository{Db: server.DB}}
	leads, err := usecase.Handle(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, leads)
}

func InsertLead(c *gin.Context, server *Server) {
	var lead models.Lead
	if err := c.BindJSON(&lead); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	usecase := usecases.InsertLeadUseCase{Repo: &repositories.MongoLeadRepository{Db: server.DB}}
	err := usecase.Handle(c.Request.Context(), lead)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lead)
}
