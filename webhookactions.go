package swarm

import (
	"context"
	"fmt"
	"net/http"
)

// WebhookActionsService handles communication to the webhookactions API endpoint
type WebhookActionsService service

const actionWebhookPath = "authenticated/webhookactions"

// WebhookAction is the generic type used for creating and returning webhooks
type WebhookAction struct {
	ID                    string                 `json:"id"`
	Name                  string                 `json:"name"`
	URL                   string                 `json:"url"`
	Method                string                 `json:"method"`
	Headers               []WebhookActionsHeader `json:"headers"`
	MaxConcurrentRequests int                    `json:"maxConcurrentRequests"`
	VerifyTLSCertificate  bool                   `json:"verifyTlsCertificate"`
	SuccessCodes          []int                  `json:"successCodes"`
	RetryIntervalSeconds  int                    `json:"retryIntervalSeconds"`
	MaxRetries            int                    `json:"maxRetries"`
}

// WebhookActionsHeader is a nested type for the WebhookActions struct
type WebhookActionsHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// List all action webhooks
func (s *WebhookActionsService) List(ctx context.Context) ([]*WebhookAction, *http.Response, error) {
	req, err := s.client.NewRequestWithBaseURL("GET", actionWebhookPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var a []*WebhookAction
	resp, err := s.client.DoRequest(ctx, req, &a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// Get an action webhook by ID
func (s *WebhookActionsService) Get(ctx context.Context, id string) (*WebhookAction, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", actionWebhookPath, id)
	req, err := s.client.NewRequestWithBaseURL("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	a := new(WebhookAction)
	resp, err := s.client.DoRequest(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// Create an action webhook
func (s *WebhookActionsService) Create(ctx context.Context, i *WebhookAction) (*WebhookAction, *http.Response, error) {
	req, err := s.client.NewRequestWithBaseURL("POST", actionWebhookPath, i)
	if err != nil {
		return nil, nil, err
	}

	a := new(WebhookAction)
	resp, err := s.client.DoRequest(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// Update an action webhook
func (s *WebhookActionsService) Update(ctx context.Context, i *WebhookAction) (*WebhookAction, *http.Response, error) {
	req, err := s.client.NewRequestWithBaseURL("PUT", actionWebhookPath, i)
	if err != nil {
		return nil, nil, err
	}

	a := new(WebhookAction)
	resp, err := s.client.DoRequest(ctx, req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// Delete an action webhook by ID
func (s *WebhookActionsService) Delete(ctx context.Context, id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", actionWebhookPath, id)
	req, err := s.client.NewRequestWithBaseURL("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	return resp, err
}

// DeleteAll action webhooks
func (s *WebhookActionsService) DeleteAll(ctx context.Context) (*http.Response, error) {
	path := fmt.Sprintf("%s/all", actionWebhookPath)
	req, err := s.client.NewRequestWithBaseURL("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	return resp, err
}
