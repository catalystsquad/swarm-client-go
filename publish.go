package swarm

import (
	"context"
	"fmt"
	"net/http"
)

const publishPath = "authenticated/publish"

// PublishService handles communication to the publish API endpoint
type PublishService service

// Publish sends data to a specified pipeline. The data input is sent as the
// body of the request.
func (s *PublishService) Publish(ctx context.Context, pipelineID string, data interface{}) (*http.Response, error) {
	path := fmt.Sprintf("%s?pipelineId=%s", publishPath, pipelineID)
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
