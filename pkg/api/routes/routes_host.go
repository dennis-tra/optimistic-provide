package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/controller"
)

type HostRoute struct {
	ctrlr  *controller.HostController
	router gin.IRouter
}

func NewHostRoute(ctrlr *controller.HostController, router gin.IRouter) HostRoute {
	return HostRoute{
		ctrlr:  ctrlr,
		router: router,
	}
}

func (p HostRoute) Setup() {
	hosts := p.router.Group("/hosts")
	{
		hosts.POST("/", p.ctrlr.Create)
		hosts.GET("/", p.ctrlr.List)
		hosts.GET("/:hostID", p.ctrlr.Get)
		hosts.DELETE("/:hostID", p.ctrlr.Stop)
		hosts.POST("/:hostID/bootstrap", p.ctrlr.Bootstrap)
	}
}
