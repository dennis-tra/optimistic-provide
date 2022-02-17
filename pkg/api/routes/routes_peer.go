package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/controller"
)

type PeerRoute struct {
	ctrlr  *controller.PeerController
	router gin.IRouter
}

func NewPeerRoute(ctrlr *controller.PeerController, router gin.IRouter) PeerRoute {
	return PeerRoute{
		ctrlr:  ctrlr,
		router: router,
	}
}

func (p PeerRoute) Setup() {
	peers := p.router.Group("/peers")
	{
		peers.GET("/:id", p.ctrlr.Get)
	}
}
