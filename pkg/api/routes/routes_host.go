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
	hosts := p.router.Group("/peers")
	{
		hosts.POST("/hosts", p.ctrlr.Create)
		hosts.GET("/hosts", p.ctrlr.List)
		hosts.GET("/hosts/:peerID", p.ctrlr.Get)
		hosts.DELETE("/hosts/:peerID", p.ctrlr.Stop)
		hosts.POST("/hosts/:peerID/bootstrap", p.ctrlr.Bootstrap)
		hosts.POST("/hosts/:peerID/dht/refresh", p.ctrlr.RefreshRoutingTable)
		hosts.POST("/hosts/:peerID/dht/save", p.ctrlr.SaveRoutingTable)
	}
}
