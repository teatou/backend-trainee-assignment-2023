package app

import (
	"backend-trainee-assignment-2023/internal/config"
	"backend-trainee-assignment-2023/internal/myapi"
	"backend-trainee-assignment-2023/internal/myloader"
	"backend-trainee-assignment-2023/internal/mylogic"
	"backend-trainee-assignment-2023/pkg/mylogger"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Avito(configFilename string) error {
	cfg, err := config.LoadConfig(configFilename)
	if err != nil {
		return fmt.Errorf("uploading config error: %w", err)
	}

	logger, err := mylogger.NewZapLogger(cfg.Logger.Level)
	if err != nil {
		return fmt.Errorf("making mylogger error: %w", err)
	}
	defer logger.Sync()

	loader, err := myloader.NewPgLoader(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName)
	if err != nil {
		return fmt.Errorf("making myloader error: %w", err)
	}
	logic := mylogic.NewLogic(loader)

	api := myapi.NewApi(logic, logger)

	r := mux.NewRouter()

	userR := r.PathPrefix("/user").Methods(http.MethodPost).Subrouter()
	userR.HandleFunc("/add", api.AddUser)
	userR.HandleFunc("/remove", api.RemoveUser)

	segmentR := r.PathPrefix("/segment").Methods(http.MethodPost).Subrouter()
	segmentR.HandleFunc("/add", api.AddSegment)
	segmentR.HandleFunc("/remove", api.RemoveSegment)

	userSegmentsR := r.PathPrefix("/usersegments").Methods(http.MethodPost).Subrouter()
	userSegmentsR.HandleFunc("/update", api.UpdateUserSegments)
	userSegmentsR.HandleFunc("/get", api.GetUserActiveSegments)

	logger.Infof("server starts on port, %d", cfg.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), r)
	if err != nil {
		logger.Fatalf("server is down: %v", err)
	}

	return nil
}
