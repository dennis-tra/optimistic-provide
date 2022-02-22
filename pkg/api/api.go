package api

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log"
	"github.com/pkg/errors"

	"github.com/dennis-tra/optimistic-provide/pkg/api/controller"
	"github.com/dennis-tra/optimistic-provide/pkg/api/middlewares"
	"github.com/dennis-tra/optimistic-provide/pkg/config"
	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

var log = logging.Logger("optprov")

// Run starts the REST API to control libp2p hosts.
func Run(ctx context.Context, cfg *config.Config) (*http.Server, error) {
	router := gin.Default()

	router.Use(cors.Default())

	dbclient, err := db.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "new db client")
	}

	peerRepo := repo.NewPeerRepo(dbclient)
	rtRepo := repo.NewRoutingTableRepo(dbclient)
	provideRepo := repo.NewProvideRepo(dbclient)
	maRepo := repo.NewMultiAddressRepo(dbclient)
	iaRepo := repo.NewIPAddressRepo(dbclient)
	dialRepo := repo.NewDialRepo(dbclient)
	connRepo := repo.NewConnectionRepo(dbclient)
	fnRepo := repo.NewFindNodesRepo(dbclient)
	cpRepo := repo.NewCloserPeersRepo(dbclient)
	psRepo := repo.NewPeerStateRepo(dbclient)

	peerService := service.NewPeerService(peerRepo)
	rtService := service.NewRoutingTableService(peerService, rtRepo)
	hostService := service.NewHostService(peerService, rtService)
	maService := service.NewMultiAddressService(maRepo, iaRepo)
	dialService := service.NewDialService(peerService, maService, dialRepo)
	connService := service.NewConnectionService(peerService, maService, connRepo)
	fnService := service.NewFindNodesService(peerService, fnRepo, cpRepo)
	psService := service.NewPeerStateService(peerService, psRepo)
	provideService := service.NewProvideService(peerService, hostService, rtService, maService, dialService, connService, fnService, psService, provideRepo)

	peerController := controller.NewPeerController(ctx, peerService)
	hostController := controller.NewHostController(ctx, hostService)
	provideController := controller.NewProvideController(ctx, provideService, hostService)
	routingTableController := controller.NewRoutingTableController(ctx, rtService, hostService)

	hosts := router.Group("/hosts")
	{
		hosts.POST("/", hostController.Create)
		hosts.GET("/", hostController.List)

		hostID := hosts.Group("/:hostID")
		{
			hostID.Use(middlewares.HostID(hostService))
			hostID.GET("/", hostController.Get)
			hostID.DELETE("/", hostController.Stop)
			hostID.POST("/bootstrap", hostController.Bootstrap)

			provides := hostID.Group("provides")
			{
				provides.POST("/", provideController.Create)
			}

			routingTables := hostID.Group("routing-tables")
			{
				routingTables.POST("/", routingTableController.Create)
				routingTables.GET("/", routingTableController.List)
				routingTables.GET("/listen", routingTableController.Listen)
				routingTables.POST("/refresh", routingTableController.Refresh)

				routingTableID := routingTables.Group("/:routingTableID")
				{
					routingTableID.Use(middlewares.RoutingTableID)
					routingTableID.GET("/", routingTableController.Get)
				}
			}
		}
	}

	peers := router.Group("/peers")
	{
		peerID := peers.Group("/:peerID")
		{
			peerID.Use(middlewares.PeerID)
			peerID.GET("/", peerController.Get)
		}
	}

	srv := &http.Server{
		Addr:    cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv, nil
}
