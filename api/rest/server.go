package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"template/api/rest/template"
	"net/http"
)

func BuildServer() *gin.Engine {
	server := gin.New()
	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization"},
	}))
	template.RegisterRoutes(server)
	return server
}

func HttpBuildServer() *http.Server {
	server := BuildServer()
	s := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}
	return s
}
