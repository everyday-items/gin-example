package v1

import (
	"net/http"

	"github.com/everyday-items/gin-example/library/app"
	"github.com/everyday-items/gin-example/library/e"
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
