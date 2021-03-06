package api

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/dennis-tra/optimistic-provide/pkg/api/controller"
	"github.com/dennis-tra/optimistic-provide/pkg/api/middlewares"
	"github.com/dennis-tra/optimistic-provide/pkg/config"
	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

// Run starts the REST API to control libp2p hosts.
func Run(ctx context.Context, cfg *config.Config) (*http.Server, error) {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(Logger(), gin.Recovery())

	gin.DebugPrintRouteFunc = func(method, absPath, handlerName string, handlerCount int) {
		log.WithFields(log.Fields{
			"handlerCount": handlerCount,
			"handlerName":  handlerName,
		}).Infof("%s %s", method, absPath)
	}

	dbclient, err := db.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "new db client")
	}

	peerRepo := repo.NewPeerRepo(dbclient)
	rtRepo := repo.NewRoutingTableRepo(dbclient)
	provideRepo := repo.NewProvideRepo(dbclient)
	retrieveRepo := repo.NewRetrievalRepo(dbclient)
	maRepo := repo.NewMultiAddressRepo(dbclient)
	iaRepo := repo.NewIPAddressRepo(dbclient)
	dialRepo := repo.NewDialRepo(dbclient)
	connRepo := repo.NewConnectionRepo(dbclient)
	fnRepo := repo.NewFindNodesRPCRepo(dbclient)
	gpRepo := repo.NewGetProvidersRepo(dbclient)
	apRepo := repo.NewAddProvidersRepo(dbclient)
	cpRepo := repo.NewCloserPeersRepo(dbclient)
	psRepo := repo.NewPeerStateRepo(dbclient)
	hostRepo := repo.NewHostRepo(dbclient)

	peerService := service.NewPeerService(peerRepo)
	rtService := service.NewRoutingTableService(peerService, rtRepo)
	hostService := service.NewHostService(peerService, rtService, hostRepo)
	maService := service.NewMultiAddressService(maRepo, iaRepo)
	dialService := service.NewDialService(peerService, maService, dialRepo)
	connService := service.NewConnectionService(peerService, maService, connRepo)
	fnService := service.NewFindNodesService(peerService, maService, fnRepo, cpRepo)
	apService := service.NewAddProvidersService(peerService, maService, apRepo, cpRepo)
	psService := service.NewPeerStateService(peerService, psRepo)
	gpService := service.NewGetProvidersService(peerService, maService, gpRepo)
	provideService := service.NewProvideService(peerService, hostService, rtService, maService, dialService, connService, fnService, psService, apService, provideRepo)
	retrievalService := service.NewRetrievalService(hostService, rtService, dialService, connService, gpService, psService, retrieveRepo)
	pmService := service.NewProvideMeasurementService(provideService)

	peerController := controller.NewPeerController(ctx, peerService)
	hostController := controller.NewHostController(ctx, hostService)
	provideController := controller.NewProvideController(ctx, provideService, hostService)
	retrievalController := controller.NewRetrievalController(ctx, retrievalService, hostService)
	routingTableController := controller.NewRoutingTableController(ctx, rtService, hostService)
	dialController := controller.NewDialController(ctx, dialService, provideService)
	connController := controller.NewConnectionController(ctx, connService, provideService)
	pmController := controller.NewProvideMeasurementController(ctx, hostService, pmService)

	hosts := router.Group("/hosts")
	{
		hosts.POST("", hostController.Create)
		hosts.GET("", hostController.List)

		hostID := hosts.Group("/:hostID")
		{
			hostID.Use(middlewares.HostID(hostService))
			hostID.GET("", hostController.Get)
			hostID.DELETE("", hostController.Archive)
			hostID.POST("/stop", hostController.Stop)
			hostID.POST("/start", hostController.Start)
			hostID.POST("/bootstrap", hostController.Bootstrap)
			hostID.GET("/routing-table", routingTableController.Current)

			provides := hostID.Group("provides")
			{
				provides.POST("", provideController.Create)
				provides.GET("", provideController.List)

				provideID := provides.Group("/:provideID")
				{
					provideID.Use(middlewares.ProvideID)
					provideID.GET("", provideController.Get)
					provideID.GET("/dials", dialController.List)
					provideID.GET("/graph", provideController.Graph)
				}
			}

			retrievals := hostID.Group("retrievals")
			{
				retrievals.POST("", retrievalController.Create)
				retrievals.GET("", retrievalController.List)

				retrievalID := retrievals.Group("/:retrievalID")
				{
					retrievalID.Use(middlewares.ProvideID)
					// provideID.GET("", retrievalController.Get)
				}
			}

			routingTables := hostID.Group("routing-tables")
			{
				routingTables.POST("", routingTableController.Create)
				routingTables.GET("", routingTableController.List)
				routingTables.GET("/listen", routingTableController.Listen)
				routingTables.POST("/refresh", routingTableController.Refresh)

				routingTableID := routingTables.Group("/:routingTableID")
				{
					routingTableID.Use(middlewares.RoutingTableID)
					routingTableID.GET("", routingTableController.Get)
				}
			}
		}
	}

	provides := router.Group("/provides")
	{
		provideID := provides.Group("/:provideID")
		{
			provideID.Use(middlewares.ProvideID)
			provideID.GET("/dials", dialController.List)
			provideID.GET("/connections", connController.List)
		}
	}

	peers := router.Group("/peers")
	{
		peerID := peers.Group("/:peerID")
		{
			peerID.Use(middlewares.PeerID)
			peerID.GET("", peerController.Get)
		}
	}

	provideMeasurements := router.Group("/provide-measurements")
	{
		provideMeasurements.POST("", pmController.Create)
		measurementID := provides.Group("/:measurementID")
		{
			measurementID.Use(middlewares.MeasurementID)
			measurementID.POST("/stop", pmController.Stop)
		}
	}

	srv := &http.Server{
		Addr:    cfg.HTTP.Host + ":" + cfg.HTTP.Port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return srv, nil
}
