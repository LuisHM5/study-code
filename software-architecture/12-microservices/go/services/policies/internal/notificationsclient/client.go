// Package notificationsclient es el adaptador que habla con el servicio notifications.
package notificationsclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"arquitecturas/12-microservices/contracts"
	"arquitecturas/12-microservices/services/policies/internal/domain"
)

type Client struct {
	baseURL string
	http    *http.Client
}

func New(baseURL string) *Client {
	// Toda llamada saliente lleva timeout.
	return &Client{baseURL: baseURL, http: &http.Client{Timeout: 2 * time.Second}}
}

func (c *Client) PolicyIssued(ctx context.Context, p *domain.Policy) error {
	body, err := json.Marshal(contracts.PolicyIssuedNotification{PolicyID: p.ID, Number: p.Number})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/notifications", bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("notifications service: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return fmt.Errorf("notifications service: status %d", res.StatusCode)
	}
	return nil
}
