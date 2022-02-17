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
		hosts.GET("/:peerID", p.ctrlr.Get)
		hosts.DELETE("/:peerID", p.ctrlr.Stop)
		hosts.POST("/:peerID/bootstrap", p.ctrlr.Bootstrap)
		hosts.POST("/:peerID/dht/refresh", p.ctrlr.RefreshRoutingTable)
		hosts.POST("/:peerID/dht/save", p.ctrlr.SaveRoutingTable)
	}
}
