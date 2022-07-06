package swarm

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testAPIToken = "ABC123DEF467"
)

func TestAPITokens_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/apitokens", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "GET", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, `["`+testAPIToken+`"]`)
	})

	ctx := context.Background()
	got, _, err := client.APITokens.List(ctx)
	var tokenStruct APIToken = testAPIToken
	want := []*APIToken{&tokenStruct}

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestAPITokens_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/apitokens", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "POST", r.Method)
		require.Contains(t, r.Header, "Authorization")
		fmt.Fprint(w, `"`+testAPIToken+`"`)
	})

	ctx := context.Background()
	got, _, err := client.APITokens.Create(ctx)
	var tokenStruct APIToken = testAPIToken
	want := &tokenStruct

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestAPITokens_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/apitokens/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	_, err := client.APITokens.Delete(ctx, testAPIToken)

	require.NoError(t, err)
}

func TestAPITokens_DeleteAll(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/authenticated/apitokens/", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "DELETE", r.Method)
		require.Contains(t, r.Header, "Authorization")
	})

	ctx := context.Background()
	_, err := client.APITokens.DeleteAll(ctx)

	require.NoError(t, err)
}
