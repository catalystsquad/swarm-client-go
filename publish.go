package swarm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const publishPath = "authenticated/publish"

// PublishService handles communication to the publish API endpoint
type PublishService service

// Publish sends data to a specified pipeline by it's name. The data input is
// sent as the body of the request.
func (s *PublishService) Publish(ctx context.Context, pipelineName string, data interface{}) (*http.Response, error) {
	path := fmt.Sprintf("%s?name=%s", publishPath, url.QueryEscape(pipelineName))
	req, err := s.client.NewRequestWithCustomerURL("POST", path, data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// PublishByID sends data to a specified pipeline by it's ID. The data input is
// sent as the body of the request.
func (s *PublishService) PublishByID(ctx context.Context, pipelineID string, data interface{}) (*http.Response, error) {
	path := fmt.Sprintf("%s?id=%s", publishPath, url.QueryEscape(pipelineID))
	req, err := s.client.NewRequestWithCustomerURL("POST", path, data)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.DoRequest(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
