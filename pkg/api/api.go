package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log"

	"github.com/dennis-tra/optimistic-provide/pkg/controller"
	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/services"
)

var log = logging.Logger("optprov")

// Start starts the REST API to control libp2p hosts.
func Start(ctx context.Context, host string, port string, dbc *db.Client) *http.Server {
	router := gin.Default()

	rts := services.NewRoutingTableService(dbc)
	hs := services.NewHostService(dbc)
	ps := services.NewProvideService(dbc)

	hostctl := controller.NewHostController(ctx, dbc, rts, hs)
	provctl := controller.NewProvideController(ctx, dbc, ps, rts, hs)

	v1 := router.Group("/v1")
	{
		v1.POST("/hosts", hostctl.Create)
		v1.GET("/hosts", hostctl.List)
		v1.GET("/hosts/:peerID", hostctl.Get)
		v1.DELETE("/hosts/:peerID", hostctl.Stop)
		v1.POST("/hosts/:peerID/bootstrap", hostctl.Bootstrap)
		v1.POST("/hosts/:peerID/dht/refresh", hostctl.RefreshRoutingTable)
		v1.POST("/hosts/:peerID/dht/save", hostctl.SaveRoutingTable)

		v1.POST("/provides", provctl.Create)
	}

	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv
}
