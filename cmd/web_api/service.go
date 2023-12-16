package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/servicesmesh-examples/services/web_router"
)

var (
	_ servicemesh.Service           = &exampleRouteInitializer{}
	_ web_router.IsRouteInitializer = &exampleRouteInitializer{}
)

type exampleRouteInitializer struct{}

func (s *exampleRouteInitializer) Init(mesh servicemesh.M) {
	// nothing to do
}

func (s *exampleRouteInitializer) Name() string {
	return "Example"
}

func (s *exampleRouteInitializer) InitRoutes(group *gin.RouterGroup) {
	group.GET("foobar", s.exampleHandler)
}

func (s *exampleRouteInitializer) exampleHandler(c *gin.Context) {
	c.String(http.StatusOK, "foobar is foobar")
}
