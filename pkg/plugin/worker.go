package plugin

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type PluginWorkerConfig struct {
	plugin     Plugin
	host       *url.URL
	interval   time.Duration
	client     *http.Client
}

type PluginWorker struct {
	cfg PluginWorkerConfig
}

type PluginWorkerOption func(*PluginWorkerConfig)

func WithClient(client *http.Client) PluginWorkerOption {
	return func(cfg *PluginWorkerConfig) {
		cfg.client = client
	}
}

func WithHost(host *url.URL) PluginWorkerOption {
	return func(cfg *PluginWorkerConfig) {
		cfg.host = host
	}
}

func WithPlugin(plugin Plugin) PluginWorkerOption {
  return func(cfg *PluginWorkerConfig) {
    cfg.plugin = plugin
  }
}

func WithInterval(interval time.Duration) PluginWorkerOption {
	return func(cfg *PluginWorkerConfig) {
		cfg.interval = interval
	}
}

func (c *PluginWorkerConfig) IsValid() bool {
	return c.host != nil && c.plugin.HostPath != nil && c.interval > 0 && c.client != nil && c.plugin.Name != ""
}

var defaultCfg = PluginWorkerConfig{
	client:   http.DefaultClient,
	interval: 500 * time.Millisecond,
}

func NewPluginWorker(opts ...PluginWorkerOption) (*PluginWorker, error) {
	cfg := defaultCfg
	for _, opt := range opts {
		opt(&cfg)
	}
	if !cfg.IsValid() {
		return nil, errors.New("invalid config")
	}

	return &PluginWorker{cfg: cfg}, nil
}

func (w *PluginWorker) RunHeartbeat(ctx context.Context) error {
	ticker := time.NewTicker(w.cfg.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := w.Heartbeat(ctx); err != nil {
				return err
			}
		}
	}
}
