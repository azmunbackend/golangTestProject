package handlermanager

import (
	"test-crm/internal/client/seller"
	sellerdb "test-crm/internal/client/seller/db"
	"test-crm/internal/client/user"
	userdb "test-crm/internal/client/user/db"
	"test-crm/internal/client/worker"
	workerdb "test-crm/internal/client/worker/db"
	"test-crm/pkg/client/postgresql"
	"test-crm/pkg/logging"

	"github.com/gorilla/mux"
)

const (
	sellerURL = "/seller"   //satyjy
	userURL = "/user"
	workerURL = "/worker"
)

func Manager(client postgresql.Client,  logger *logging.Logger) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//Yonekey CRUD
	SellerRouterManager := router.PathPrefix(sellerURL).Subrouter()
	SellerRouterRepository := sellerdb.NewRepository(client, logger)
	SellerRouterHandler := seller.NewHandler(SellerRouterRepository, logger)
	SellerRouterHandler.Register(SellerRouterManager)
	
	//Login Register
	UserRouterManager := router.PathPrefix(userURL).Subrouter()
	UserRouterRepository := userdb.NewRepository(client, logger)
	UserRouterHandler := user.NewHandler(UserRouterRepository, logger)
	UserRouterHandler.Register(UserRouterManager)

	//worker
	WorkerRouterManager := router.PathPrefix(workerURL).Subrouter()
	WorkerRouterRepository := workerdb.NewRepository(client, logger)
	WorkerHandler := worker.NewHandler(WorkerRouterRepository, logger)
	WorkerHandler.Register(WorkerRouterManager)
	
	return router
}
