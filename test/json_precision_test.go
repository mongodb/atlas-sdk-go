package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312021/admin"
)

// TestDecodePreservesLargeIntegers verifies that integers above 2^53 in dynamic
// fields (any, []any, map[string]any) are not silently rounded to float64 by the
// SDK's response decoder (admin/client.go decode function).
func TestDecodePreservesLargeIntegers(t *testing.T) {
	// 2^53 + 1 = 9007199254740993, the first integer float64 cannot represent exactly.
	const payload = `{"name":"test-proc","pipeline":[{"$match":{"tenantId":9007199254740993}}]}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.atlas.2024-05-30+json")
		fmt.Fprint(w, payload)
	}))
	defer srv.Close()

	sdk, err := admin.NewClient(
		admin.UseDigestAuth("key", "secret"),
		admin.UseBaseURL(srv.URL),
	)
	require.NoError(t, err)

	proc, httpResp, err := sdk.StreamsAPI.GetStreamProcessor(t.Context(), "groupId", "tenantName", "test-proc").Execute()
	if httpResp != nil {
		defer httpResp.Body.Close()
	}
	require.NoError(t, err)

	pipeline := proc.GetPipeline()
	require.Len(t, pipeline, 1)

	stage, ok := pipeline[0].(map[string]any)
	require.True(t, ok, "pipeline stage must be map[string]any")

	match, ok := stage["$match"].(map[string]any)
	require.True(t, ok, "$match must be map[string]any")

	tenantID, ok := match["tenantId"]
	require.True(t, ok, "tenantId must be present")

	num, ok := tenantID.(json.Number)
	assert.True(t, ok, "tenantId must be json.Number, not float64 — large integer precision loss detected")
	assert.Equal(t, "9007199254740993", num.String())
}
