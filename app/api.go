package app

import "github.com/gin-gonic/gin"

func NewAPI(s *Server) {
	s.Router.GET("/status", HeathCheck)
	s.Router.GET("/leads", func(c *gin.Context) {
		GetAllLeads(c, s)
	})
	s.Router.POST("/leads", func(c *gin.Context) {
		InsertLead(c, s)
	})
}
