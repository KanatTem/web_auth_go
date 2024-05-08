package main

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"web_auth/internal/routes"
)

func main() {
	r := gin.Default()

	r.Static("/static", "web/static")
	path := filepath.Join("web", "templates", "**", "*.tmpl")

	r.LoadHTMLGlob(filepath.Join(path))

	r.GET("/dashboard", routes.Dashboard)

	r.Run(":8080")
}
