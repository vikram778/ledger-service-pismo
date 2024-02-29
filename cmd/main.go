package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"ledger-service-pismo/pkg/config"
	"ledger-service-pismo/pkg/controller"
	"ledger-service-pismo/pkg/log"
	"ledger-service-pismo/pkg/postgres"
	"ledger-service-pismo/pkg/repository"
	service "ledger-service-pismo/pkg/service"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	log.SetLogLevel()

	log.Info("Starting server")

	configPath := "/pkg/config/config-local.yml" //config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatal("Loading config:", zap.Error(err))
	}

	psqlDB, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Postgresql init: %s", err))
	}
	defer psqlDB.Close()

	log.Info(fmt.Sprintf("PostgreSQL connected: %#v", psqlDB.Stats()))

	s := repository.NewDBOpsRepository(psqlDB)
	itemService := service.NewLedgerService(s)
	h := controllers.NewLedgerController(itemService)

	r := mux.NewRouter()

	r.HandleFunc("/accounts", h.CreateAccountHandler).Methods(http.MethodPost)
	r.HandleFunc("/accounts/{account_id}", h.GetAccountHandler).Methods(http.MethodGet)
	r.HandleFunc("/transactions", h.CreateTransactionHandler).Methods(http.MethodPost)

	log.Fatal("Starting Service Error", zap.Error(http.ListenAndServe(":3000", r)))
}
