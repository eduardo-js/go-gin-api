package router

import (
	v1 "github.com/eduardo-js/go-gin-api/src/controllers/v1"
	"github.com/gin-gonic/gin"
)

func RouteHandler() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")

	{
		apiV1.GET("pokemons", v1.Get)
	}

	return r
}
