package main

import (
	"context"
	"embed"
	"log"
	"log/slog"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/choffmann/green-ecolution-demo-plugin/internal/server"
	"github.com/green-ecolution/green-ecolution-backend/pkg/plugin"
	"github.com/joho/godotenv"
)

//go:embed all:ui/dist
var f embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	hostPathEnv := os.Getenv("HOST_PATH")

	pluginPath, err := url.Parse("http://localhost:8080/")
	if err != nil {
		panic(err)
	}

	p := plugin.Plugin{
		Slug:           "demo_plugin",
		Name:           "Demo Plugin",
		Version:        "v1.0.0",
		Description:    "This is a demo plugin for the Green Ecolution platform to showcase the plugin system",
		PluginHostPath: pluginPath,
	}

	http := server.NewServer(
		server.WithPort(8080),
		server.WithPluginFS(f),
		server.WithPlugin(p),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err = http.Run(ctx); err != nil {
			slog.Error("Error while running http server", "error", err)
		}
	}()

	hostPath, err := url.Parse(hostPathEnv)
	if err != nil {
		panic(err)
	}

	worker, err := plugin.NewPluginWorker(
		plugin.WithHost(hostPath),
		plugin.WithPlugin(p),
		plugin.WithHostAPIVersion("v1"),
	)

	_, err = worker.Register(ctx, clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	go func() {
		defer wg.Done()
		if err := worker.RunHeartbeat(ctx); err != nil {
			slog.Error("Failed to send heartbeat", "error", err)
		}
	}()

	wg.Wait()
}
