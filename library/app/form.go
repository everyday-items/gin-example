package app

import (
	"net/http"

	"github.com/everyday-items/gin-example/library/e"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data using Gin's built-in validator
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.ShouldBind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	return http.StatusOK, e.SUCCESS
}
