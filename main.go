package main

import (
	"context"
	"embed"
	"log/slog"
	"net/url"
	"os/signal"
	"sync"
	"syscall"

	"github.com/choffmann/green-ecolution-demo-plugin/internal/server"
	"github.com/choffmann/green-ecolution-demo-plugin/pkg/plugin"
)

//go:embed ui/dist/**/*
var f embed.FS

var (
	username = "demo_plugin"
	password = "demo_plugin"
)

func main() {
	pluginPath, err := url.Parse("http://localhost:8080/")
	if err != nil {
		panic(err)
	}

	p := plugin.Plugin{
		Slug:        "demo_plugin",
    Name:        "Demo Plugin",
		Version:     "v1.0.0",
		Description: "This is a demo plugin for the Green Ecolution platform to showcase the plugin system",
		PluginHostPath:    pluginPath,
	}

	http := server.NewServer(
		server.WithPort(8080),
		server.WithPluginFS(f),
		server.WithPlugin(p),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err = http.Run(ctx); err != nil {
			slog.Error("Error while running http server", "error", err)
		}
	}()

	hostPath, err := url.Parse("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	worker, err := plugin.NewPluginWorker(
		plugin.WithHost(hostPath),
		plugin.WithPlugin(p),
	)

	_, err = worker.Register(ctx, username, password)
	if err != nil {
		panic(err)
	}

	if err := worker.RunHeartbeat(ctx); err != nil {
		slog.Error("Failed to send heartbeat", "error", err)
	}

	wg.Wait()
}
