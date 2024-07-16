package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tonymj76/savannah/handlers"
)

func SetRouter(rs *handlers.RestService) *gin.Engine {
	router := gin.Default()
	apiGroupRoute := router.Group("/api")
	apiGroupRoute.Group("fetch/:owner/:repo", rs.FetchAndSaveRepoData)
	return router
}
