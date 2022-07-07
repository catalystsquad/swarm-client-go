package swarm

import (
	"context"
	"fmt"
	"net/http"
)

// APITokensService handles communication to the apitokens API endpoint
type APITokensService service

const apiTokensPath = "authenticated/apitokens"

// APIToken type, currently just a string as no other metadata exists
type APIToken string

// List all API tokens
func (s *APITokensService) List(ctx context.Context) ([]*APIToken, *http.Response, error) {
	req, err := s.client.NewRequestWithBaseURL("GET", apiTokensPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var t []*APIToken
	resp, err := s.client.DoRequest(ctx, req, &t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

// Create an API token
func (s *APITokensService) Create(ctx context.Context) (*APIToken, *http.Response, error) {
	req, err := s.client.NewRequestWithBaseURL("POST", apiTokensPath, nil)
	if err != nil {
		return nil, nil, err
	}

	t := new(APIToken)
	resp, err := s.client.DoRequest(ctx, req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

// Delete an API token by value
func (s *APITokensService) Delete(ctx context.Context, token APIToken) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", apiTokensPath, token)
	req, err := s.client.NewRequestWithBaseURL("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	return resp, err
}

// DeleteAll API tokens
func (s *APITokensService) DeleteAll(ctx context.Context) (*http.Response, error) {
	path := fmt.Sprintf("%s/all", apiTokensPath)
	req, err := s.client.NewRequestWithBaseURL("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	return resp, err
}
