package plugin

import "net/url"

type Plugin struct {
	Slug           string
	Name           string
	Version        string
	Description    string
	PluginHostPath *url.URL
}

type PluginOption func(*Plugin)

func WithName(name string) PluginOption {
	return func(p *Plugin) {
		p.Name = name
	}
}

func WithSlug(slug string) PluginOption {
	return func(p *Plugin) {
		p.Slug = slug
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
		p.PluginHostPath = hostPath
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
