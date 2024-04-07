package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"test-crm/internal/config"
	handlermanager "test-crm/internal/handlers/manager"
	"test-crm/pkg/client/postgresql"
	"test-crm/pkg/logging"
	repeatable "test-crm/pkg/utils"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
)

func main() {
	cfg := config.GetConfig()

	logger := logging.GetLogger()
	postgresSQLClient, err := postgresql.NewClient(context.TODO(),  cfg.Storage)
	fmt.Println(postgresSQLClient)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	err = repeatable.CrateDir()
	start(handlermanager.Manager(postgresSQLClient, logger), cfg, postgresSQLClient)

}

func start(router *mux.Router, cfg *config.Config, pGPool *pgxpool.Pool) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	logger.Info("listen tcp")
	listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	fileServer := http.FileServer(http.Dir("./../../Uploads"))
	router.PathPrefix("/api/Uploads/").Handler(http.StripPrefix("/api/Uploads/", fileServer))
 
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*",
		},
	})

	handler := c.Handler(router)

	server := &http.Server{
		Handler:      handler,
		WriteTimeout: 5000 * time.Second,
		ReadTimeout:  5000 * time.Second,
	}

	fmt.Println(server)
	logger.Fatal(server.Serve(listener))
}
