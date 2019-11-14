package _gin

import (
	"github.com/gin-gonic/gin/binding"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Param string
}

func handler1(c *gin.Context) {
	req := &Req{}
	_ = c.ShouldBindBodyWith(req, binding.JSON)
	c.JSON(http.StatusOK, gin.H{
		"Return": req.Param,
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/post_test", handler1)

	return router
}
