package plugin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (w *PluginWorker) Register(ctx context.Context, username, password string) (*Token, error) {
	reqBody := RegisterRequest{
    Slug: w.cfg.plugin.Slug,
		Name: w.cfg.plugin.Name,
		Path: w.cfg.plugin.PluginHostPath.String(),
    Version: w.cfg.plugin.Version,
    Description: w.cfg.plugin.Description,
		Auth: AuthRequest{
			Username: username,
			Password: password,
		},
	}

	buf, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	registerPath := fmt.Sprintf("%s://%s/api/v1/plugin/register", w.cfg.host.Scheme, w.cfg.host.Host)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, registerPath, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := w.cfg.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to register plugin")
	}

	var tokenResp ClientTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

  token := &Token{
    AccessToken:  tokenResp.AccessToken,
    RefreshToken: tokenResp.RefreshToken,
  }

  w.cfg.token = token
  return token, nil
}

func (w *PluginWorker) Heartbeat(ctx context.Context) error {
	registerPath := fmt.Sprintf("%s://%s/api/v1/plugin/%s/heartbeat", w.cfg.host.Scheme, w.cfg.host.Host, w.cfg.plugin.Slug)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, registerPath, nil)
	if err != nil {
		return err
	}

	resp, err := w.cfg.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send heartbeat")
	}

	return nil
}
