package swarm

import (
	"context"
	"fmt"
	"net/http"
)

// PipelinesService handles communication to the pipelines API endpoint
type PipelinesService service

const pipelinesPath = "authenticated/pipelines"

// Pipeline is the generic type used for creating and returning Pipelines
type Pipeline struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Steps                []PipelineSteps        `json:"steps"`
	Outputs              []string               `json:"outputs"`
	PersistOutput        bool                   `json:"persistOutput"`
	StitchConfigs        []PipelineStitchConfig `json:"stitchConfigs"`
	RetryIntervalSeconds int                    `json:"retryIntervalSeconds"`
	MaxRetries           int                    `json:"maxRetries"`
}

// PipelineSteps is a nested type for the Pipeline struct
type PipelineSteps struct {
	Function string   `json:"function"`
	Outputs  []string `json:"outputs"`
	Required bool     `json:"required"`
	Type     string   `json:"type"`
}

// PipelineStitchConfig is a nested type for the Pipeline struct
type PipelineStitchConfig struct {
	StitchPipelineID string `json:"stitchPipelineId"`
	Key              string `json:"key"`
	TTL              int    `json:"ttl"`
}

// List all pipelines
func (s *PipelinesService) List(ctx context.Context) ([]*Pipeline, *http.Response, error) {
	u, err := s.client.BaseURL.Parse(pipelinesPath)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var p []*Pipeline
	resp, err := s.client.DoRequest(ctx, req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Get a pipeline by ID
func (s *PipelinesService) Get(ctx context.Context, id string) (*Pipeline, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", pipelinesPath, id)
	req, err := s.client.NewRequestWithBaseURL("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.DoRequest(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Create a pipeline
func (s *PipelinesService) Create(ctx context.Context, i *Pipeline) (*Pipeline, *http.Response, error) {
	req, err := s.client.NewRequestWithBaseURL("POST", pipelinesPath, i)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.DoRequest(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Update a pipeline
func (s *PipelinesService) Update(ctx context.Context, i *Pipeline) (*Pipeline, *http.Response, error) {
	req, err := s.client.NewRequestWithBaseURL("PUT", pipelinesPath, i)
	if err != nil {
		return nil, nil, err
	}

	p := new(Pipeline)
	resp, err := s.client.DoRequest(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Delete a pipeline by ID
func (s *PipelinesService) Delete(ctx context.Context, id string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", pipelinesPath, id)
	req, err := s.client.NewRequestWithBaseURL("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	return resp, err
}

// DeleteAll pipelines
func (s *PipelinesService) DeleteAll(ctx context.Context) (*http.Response, error) {
	path := fmt.Sprintf("%s/all", pipelinesPath)
	req, err := s.client.NewRequestWithBaseURL("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	return resp, err
}
