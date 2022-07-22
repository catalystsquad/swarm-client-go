package swarm

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testPublishPipelineName = "testPipeline"
	testPublishPipelineID   = "ABC123"
)

func TestPublish_Publish(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/publish", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	data := struct {
		TestKey string
	}{
		TestKey: "TestValue",
	}
	// validate a generic publish attempt
	_, err := client.Publish.Publish(ctx, testPublishPipelineName, data)
	require.NoError(t, err)

	// ensure that query parameters are safely escaped
	_, err = client.Publish.Publish(ctx, "pipeline with spaces", data)
	require.NoError(t, err)
}

func TestPublish_PublishByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/publish", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	data := struct {
		TestKey string
	}{
		TestKey: "TestValue",
	}
	_, err := client.Publish.Publish(ctx, testPublishPipelineID, data)
	require.NoError(t, err)
}
