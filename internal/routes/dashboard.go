package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_auth/internal/models"
)

func Dashboard(c *gin.Context) {
	apps := []models.AppView{
		{ID: "1", Name: "Billing System", Description: "Handles billing", Roles: []string{"admin", "viewer"}, UserCount: 5},
		{ID: "2", Name: "HR System", Description: "Handles HR", Roles: []string{"editor"}, UserCount: 3},
	}

	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
		"Title":      "Dashboard",
		"Apps":       apps,
		"TotalUsers": 12,
	})
}
