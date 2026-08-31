package test

// Path parameters equal to dot-segments ("." or "..") must be rejected
// client-side before any request is sent: the server normalizes them per
// RFC 3986, which would route the authenticated request to a different
// endpoint than intended.

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.mongodb.org/atlas-sdk/v20250312023/admin"
)

func TestDotSegmentPathParamsRejected(t *testing.T) {
	var hits int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer server.Close()

	sdk, err := admin.NewClient(admin.UseBaseURL(server.URL))
	require.NoError(t, err)
	ctx := context.Background()

	cases := []struct {
		name      string
		paramName string
		call      func() error
	}{
		{"DeleteIndexByName groupId=..", "groupId", func() error {
			_, err := sdk.AtlasSearchAPI.DeleteIndexByName(ctx, "..", "cluster0", "coll", "db", "idx").Execute()
			return err
		}},
		{"DeleteIndexByName collectionName=..", "collectionName", func() error {
			_, err := sdk.AtlasSearchAPI.DeleteIndexByName(ctx, "group1", "cluster0", "..", "db", "idx").Execute()
			return err
		}},
		{"DeleteIndexByName collectionName=.", "collectionName", func() error {
			_, err := sdk.AtlasSearchAPI.DeleteIndexByName(ctx, "group1", "cluster0", ".", "db", "idx").Execute()
			return err
		}},
		{"GetDatabaseUser username=..", "username", func() error {
			_, _, err := sdk.DatabaseUsersAPI.GetDatabaseUser(ctx, "group1", "admin", "..").Execute()
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			before := atomic.LoadInt32(&hits)
			err := tc.call()
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.paramName+" must not be a dot-segment path parameter")
			assert.Equal(t, before, atomic.LoadInt32(&hits), "rejected call must not reach the network")
		})
	}
}

func TestNonDotSegmentPathParamsStillWork(t *testing.T) {
	var capturedURI string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedURI = r.RequestURI
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{}"))
	}))
	defer server.Close()

	sdk, err := admin.NewClient(admin.UseBaseURL(server.URL))
	require.NoError(t, err)
	ctx := context.Background()

	cases := []struct {
		name       string
		collection string
		wantInURI  string
	}{
		{"dotted name", "my.coll", "/my.coll/"},
		{"three dots is not a dot-segment", "...", "/.../"},
		{"dots inside a name", "a..b", "/a..b/"},
		{"slashes still escaped, not a traversal", "a/../b", "/a%2F..%2Fb/"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			capturedURI = ""
			_, err := sdk.AtlasSearchAPI.DeleteIndexByName(ctx, "group1", "cluster0", tc.collection, "db", "idx").Execute()
			require.NoError(t, err)
			assert.Contains(t, capturedURI, tc.wantInURI)
		})
	}
}

func TestPrepareRequestRejectsDotSegments(t *testing.T) {
	sdk, err := admin.NewClient(admin.UseBaseURL("https://cloud.mongodb.com"))
	require.NoError(t, err)
	ctx := context.Background()

	badPaths := []string{
		"https://cloud.mongodb.com/api/atlas/v2/groups/g1/databaseUsers/admin/..",
		"https://cloud.mongodb.com/api/atlas/v2/groups/g1/./databaseUsers",
		"https://cloud.mongodb.com/api/atlas/v2/groups/g1/databaseUsers/%2E%2E",
	}
	for _, p := range badPaths {
		_, err := sdk.UntypedClient.PrepareRequest(ctx, p, http.MethodGet, nil, map[string]string{}, url.Values{}, url.Values{}, nil)
		require.Error(t, err, p)
		assert.Contains(t, err.Error(), "dot-segment")
	}

	req, err := sdk.UntypedClient.PrepareRequest(ctx,
		"https://cloud.mongodb.com/api/atlas/v2/groups/g1/databaseUsers",
		http.MethodGet, nil, map[string]string{}, url.Values{}, url.Values{}, nil)
	require.NoError(t, err)
	assert.Equal(t, "/api/atlas/v2/groups/g1/databaseUsers", req.URL.Path)
}
