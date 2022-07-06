package swarm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURLv1             = "https://api.swarmiolabs.com/v1/"
	customerURLTemplatev1 = "https://%s.api.swarmiolabs.com/v1/"
)

// Client is the primary interface for all Swarm service handlers
type Client struct {
	apiKey     string
	httpClient *http.Client

	// Base URL for most API requests
	BaseURL *url.URL

	// Customer URL for Publish API requests
	CustomerURL *url.URL

	// Exposes services via the client
	APITokens      *APITokensService
	Pipelines      *PipelinesService
	Publish        *PublishService
	WebhookActions *WebhookActionsService
}

// Generic service to implement in all service types
type service struct {
	client *Client
}

// NewClient is a constructor for Client
func NewClient(customerID string, apiKey string) *Client {
	baseURL, err := url.Parse(baseURLv1)
	if err != nil {
		panic(err)
	}

	customerURL, err := url.Parse(fmt.Sprintf(customerURLTemplatev1, customerID))
	if err != nil {
		panic(err)
	}

	c := &Client{
		BaseURL:     baseURL,
		apiKey:      apiKey,
		CustomerURL: customerURL,
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	c.APITokens = &APITokensService{client: c}
	c.Pipelines = &PipelinesService{client: c}
	c.Publish = &PublishService{client: c}
	c.WebhookActions = &WebhookActionsService{client: c}

	return c
}

// NewRequestWithBaseURL builds an http.Request using the BaseURL.
func (s *Client) NewRequestWithBaseURL(method string, path string, body interface{}) (*http.Request, error) {
	u, err := s.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}
	return s.NewRequest(method, u, body)
}

// NewRequestWithCustomerURL builds an http.Request using the CustomerURL.
func (s *Client) NewRequestWithCustomerURL(method string, path string, body interface{}) (*http.Request, error) {
	u, err := s.CustomerURL.Parse(path)
	if err != nil {
		return nil, err
	}
	return s.NewRequest(method, u, body)
}

// NewRequest builds an http.Request object. The body parameter will
// automatically be encoded to json to send in a request.
func (s *Client) NewRequest(method string, u *url.URL, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	return req, nil
}

// DoRequest will execute an http.Request. The entire http.Response will be
// returned. The JSON response will be decoded into the value pointed to v.
func (s *Client) DoRequest(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		// convert the body to a string for reporting errors easily
		b, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			return resp, readErr
		}
		return resp, fmt.Errorf("http response %d: %#v", resp.StatusCode, string(b))
	}

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}
