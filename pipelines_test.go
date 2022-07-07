package swarm

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testPipelineID   = "01G6XSSVAAT8CNRGXPSBZSA1PY"
	testPipelineJSON = `{
  "id": "01G6XSSVAAT8CNRGXPSBZSA1PY",
  "name": "pipeline1",
  "steps": [
    {
      "function": "return message.hello == 'prod';",
      "outputs": [],
      "required": true,
      "type": "filter"
    }
  ],
  "outputs": [
    "01G2YYM9RRT8CNRGXPSBZSA1PY"
  ],
  "persistOutput": false,
  "stitchConfigs": null,
  "retryIntervalSeconds": 60,
  "maxRetries": -1
}`
)

var (
	testPipelineObj = &Pipeline{
		ID:   "01G6XSSVAAT8CNRGXPSBZSA1PY",
		Name: "pipeline1",
		Steps: []PipelineSteps{{
			Function: "return message.hello == 'prod';",
			Outputs:  []string{},
			Required: true,
			Type:     "filter",
		}},
		Outputs: []string{
			"01G2YYM9RRT8CNRGXPSBZSA1PY",
		},
		PersistOutput:        false,
		StitchConfigs:        nil,
		RetryIntervalSeconds: 60,
		MaxRetries:           -1,
	}
)

func TestPipelines_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/pipelines", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, `[`+testPipelineJSON+`]`)
	})

	ctx := context.Background()
	got, _, err := client.Pipelines.List(ctx)
	want := []*Pipeline{testPipelineObj}

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestPipelines_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/pipelines/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, testPipelineJSON)
	})

	ctx := context.Background()
	got, _, err := client.Pipelines.Get(ctx, testPipelineID)

	require.NoError(t, err)
	require.Equal(t, testPipelineObj, got)
}

func TestPipelines_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/pipelines", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, testPipelineJSON)
	})

	ctx := context.Background()
	got, _, err := client.Pipelines.Create(ctx, testPipelineObj)

	require.NoError(t, err)
	require.Equal(t, testPipelineObj, got)
}

func TestPipelines_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/pipelines", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "PUT", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, testPipelineJSON)
	})

	ctx := context.Background()
	got, _, err := client.Pipelines.Update(ctx, testPipelineObj)

	require.NoError(t, err)
	require.Equal(t, testPipelineObj, got)
}

func TestPipelines_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/pipelines/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	_, err := client.Pipelines.Delete(ctx, testPipelineID)

	require.NoError(t, err)
}

func TestPipelines_DeleteAll(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/pipelines/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	_, err := client.Pipelines.DeleteAll(ctx)

	require.NoError(t, err)
}
