package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/controller"
)

type RoutingTableRoute struct {
	ctrlr  *controller.RoutingTableController
	router gin.IRouter
}

func NewRoutingTableRoute(ctrlr *controller.RoutingTableController, router gin.IRouter) RoutingTableRoute {
	return RoutingTableRoute{
		ctrlr:  ctrlr,
		router: router,
	}
}

func (p RoutingTableRoute) Setup() {
	routingTables := p.router.Group("/hosts/:hostID/routing-tables")
	{
		routingTables.POST("/", p.ctrlr.Create)
		routingTables.GET("/", p.ctrlr.List)
		routingTables.GET("/:routingTableID", p.ctrlr.Get)
	}
}
