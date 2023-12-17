package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/web_router"
)

var (
	_ servicemesh.Service           = &exampleRouteInitializer{}
	_ web_router.IsRouteInitializer = &exampleRouteInitializer{}
)

type exampleRouteInitializer struct{}

func (s *exampleRouteInitializer) Init(mesh servicemesh.Mesh) {
	// nothing to do
}

func (s *exampleRouteInitializer) Name() string {
	return "Example"
}

func (s *exampleRouteInitializer) InitRoutes(group *gin.RouterGroup) {
	group.GET("foo", s.exampleHandler)
}

func (s *exampleRouteInitializer) exampleHandler(c *gin.Context) {
	c.String(http.StatusOK, "bar")
}
