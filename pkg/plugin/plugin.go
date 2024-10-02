package plugin

import "net/url"

type Plugin struct {
	Name        string
	Version     string
	Description string
  HostPath    *url.URL
}

type PluginOption func(*Plugin)

func WithName(name string) PluginOption {
  return func(p *Plugin) {
    p.Name = name
  }
}

func WithVersion(version string) PluginOption {
  return func(p *Plugin) {
    p.Version = version
  }
}

func WithDescription(description string) PluginOption {
  return func(p *Plugin) {
    p.Description = description
  }
}

func WithHostPath(hostPath *url.URL) PluginOption {
  return func(p *Plugin) {
    p.HostPath = hostPath
  }
}

var defaultPlugin = Plugin{
  Version: "develop",
}

func NewPlugin(opts ...PluginOption) *Plugin {
  p := &Plugin{}
  for _, opt := range opts {
    opt(p)
  }
  return p
}
