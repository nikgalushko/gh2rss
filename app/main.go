package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/nikgalushko/gh2rss/app/server"
)

var revision string

func main() {
	port := flag.Int("port", 8080, "Port for the HTTP server")
	db := flag.String("db", "cache.db", "Path to the cache file")
	logLevel := flag.String("log-level", "info", "Logging level (debug, info, warn, error)")
	flag.Parse()

	fmt.Println("gh2rss", revision)

	level := slog.LevelInfo
	switch *logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: level}))
	slog.SetDefault(logger)

	slog.Info("Starting server", "port", *port, "cache", *db)

	srv := server.New()
	err := srv.Run(*port)
	if err != nil {
		slog.Error("run server", slog.String("error", err.Error()))
	}
}
