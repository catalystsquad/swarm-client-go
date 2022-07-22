package swarm

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testWebhookActionID   = "01G7816ZYJT8CNRGXPSBZSA1PY"
	testWebhookActionJSON = `{
  "id": "01G7816ZYJT8CNRGXPSBZSA1PY",
  "name": "webhook1",
  "url": "https://example.com/mywebhook/123",
  "method": "POST",
  "headers": [],
  "maxConcurrentRequests": 2,
  "verifyTlsCertificate": true,
  "successCodes": [
    200
  ],
  "retryIntervalSeconds": 60,
  "maxRetries": -1
}`
)

var (
	testWebhookActionObj = &WebhookAction{
		ID:                    "01G7816ZYJT8CNRGXPSBZSA1PY",
		Name:                  "webhook1",
		URL:                   "https://example.com/mywebhook/123",
		Method:                "POST",
		Headers:               []WebhookActionsHeader{},
		MaxConcurrentRequests: 2,
		VerifyTLSCertificate:  true,
		SuccessCodes: []int{
			200,
		},
		RetryIntervalSeconds: 60,
		MaxRetries:           -1,
	}
)

func TestWebhookActions_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/webhookactions", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, `[`+testWebhookActionJSON+`]`)
	})

	ctx := context.Background()
	got, _, err := client.WebhookActions.List(ctx)
	want := []*WebhookAction{testWebhookActionObj}

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestWebhookActions_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/webhookactions/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, testWebhookActionJSON)
	})

	ctx := context.Background()
	got, _, err := client.WebhookActions.Get(ctx, testWebhookActionID)

	require.NoError(t, err)
	require.Equal(t, testWebhookActionObj, got)
}

func TestWebhookActions_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/webhookactions", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, testWebhookActionJSON)
	})

	ctx := context.Background()
	got, _, err := client.WebhookActions.Create(ctx, testWebhookActionObj)

	require.NoError(t, err)
	require.Equal(t, testWebhookActionObj, got)
}

func TestWebhookActions_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/webhookactions/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "PUT", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, testWebhookActionJSON)
	})

	ctx := context.Background()
	got, _, err := client.WebhookActions.Update(ctx, testWebhookActionID, testWebhookActionObj)

	require.NoError(t, err)
	require.Equal(t, testWebhookActionObj, got)
}

func TestWebhookActions_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/webhookactions/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	_, err := client.WebhookActions.Delete(ctx, testWebhookActionID)

	require.NoError(t, err)
}

func TestWebhookActions_DeleteAll(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/webhookactions/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	_, err := client.WebhookActions.DeleteAll(ctx)

	require.NoError(t, err)
}
