package admin

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDecodePreservesLargeIntegers verifies that integers above 2^53 in dynamic
// fields (any, []any, map[string]any) are not silently rounded to float64.
func TestDecodePreservesLargeIntegers(t *testing.T) {
	// 2^53 + 1 = 9007199254740993, the first integer float64 cannot represent exactly.
	const payload = `{"name":"test-proc","pipeline":[{"$match":{"tenantId":9007199254740993}}]}`

	var proc StreamsProcessor
	c := &APIClient{}
	err := c.decode(&proc, io.NopCloser(strings.NewReader(payload)), "application/json")
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
