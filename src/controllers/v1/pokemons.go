package v1

import (
	"net/http"

	v1 "github.com/eduardo-js/go-gin-api/services/v1"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	resp := v1.Get()
	c.JSON(http.StatusOK, resp)
}
