package swarm

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testPublishPipelineID = "ABC123"
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
	_, err := client.Publish.Publish(ctx, testPublishPipelineID, data)
	require.NoError(t, err)
}
