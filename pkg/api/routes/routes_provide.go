package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/controller"
)

type ProvideRoute struct {
	ctrlr  *controller.ProvideController
	router gin.IRouter
}

func NewProvideRoute(ctrlr *controller.ProvideController, router gin.IRouter) ProvideRoute {
	return ProvideRoute{
		ctrlr:  ctrlr,
		router: router,
	}
}

func (p ProvideRoute) Setup() {
	provides := p.router.Group("/provides")
	{
		provides.POST("/", p.ctrlr.Create)
	}
}
