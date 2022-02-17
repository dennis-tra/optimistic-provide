package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log"
	"github.com/pkg/errors"

	"github.com/dennis-tra/optimistic-provide/pkg/api/controller"
	"github.com/dennis-tra/optimistic-provide/pkg/api/routes"
	"github.com/dennis-tra/optimistic-provide/pkg/config"
	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

var log = logging.Logger("optprov")

// Run starts the REST API to control libp2p hosts.
func Run(ctx context.Context, cfg *config.Config) (*http.Server, error) {
	router := gin.Default()

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
	provideController := controller.NewProvideController(ctx, provideService)
	routingTableController := controller.NewRoutingTableController(ctx, rtService, hostService, rtRepo)

	v1 := router.Group("/v1")
	{
		routes.NewPeerRoute(peerController, v1).Setup()
		routes.NewHostRoute(hostController, v1).Setup()
		routes.NewProvideRoute(provideController, v1).Setup()
		routes.NewRoutingTableRoute(routingTableController, v1).Setup()
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
