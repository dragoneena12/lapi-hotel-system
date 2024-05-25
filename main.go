package main

import (
	"log/slog"
	"os"

	"github.com/dragoneena12/lapi-hotel-system/auth"
	"github.com/dragoneena12/lapi-hotel-system/config"
	"github.com/dragoneena12/lapi-hotel-system/controller"
	"github.com/dragoneena12/lapi-hotel-system/db"
	"github.com/dragoneena12/lapi-hotel-system/graph"
	"github.com/dragoneena12/lapi-hotel-system/server"
)

func main() {
	cfg := config.NewConfig()
	logLevel := slog.LevelInfo
	if cfg.Debug {
		logLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	slog.SetDefault(logger)

	d, err := db.NewDBConnection(*cfg)
	if err != nil {
		slog.Error("failed to connect to database: %v", err)
		os.Exit(1)
	}
	stayRepository := db.NewStayRepository(d)
	hotelRepository := db.NewHotelRepository(d)
	handler := graph.NewHandler(auth.NewJWTAuthController(), controller.NewStayController(stayRepository, hotelRepository), controller.NewHotelController(hotelRepository), auth.JWTHasRole)
	server := server.NewServer(*cfg, handler, auth.AuthMiddleWare(cfg.Auth0Domain))
	err = server.Start()
	if err != nil {
		slog.Error("failed to start server: %v", err)
		os.Exit(1)
	}
}
