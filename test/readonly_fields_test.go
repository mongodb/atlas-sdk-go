package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312023/admin"
)

// CLOUDP-437811: properties that are required + readOnly in the OpenAPI spec
// must not be serialized in request bodies when unset. The generator emits
// omitempty for them, so their zero values are omitted from marshaled payloads.
func TestRequiredReadOnlyFieldsOmittedWhenUnset(t *testing.T) {
	group := &admin.Group{
		Name:  "my-project",
		OrgId: "5a0a1e7e0f2912c554080adc",
	}

	out, err := json.Marshal(group)
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, json.Unmarshal(out, &m))

	assert.NotContains(t, m, "clusterCount")
	// Known limitation: omitempty cannot omit zero-value structs, so unset
	// time.Time read-only fields are still serialized.
	assert.Contains(t, m, "created")
	assert.Equal(t, "my-project", m["name"])
	assert.Equal(t, "5a0a1e7e0f2912c554080adc", m["orgId"])
}

func TestRequiredReadOnlyFieldsSentWhenSet(t *testing.T) {
	group := &admin.Group{
		ClusterCount: 3,
		Name:         "my-project",
		OrgId:        "5a0a1e7e0f2912c554080adc",
	}

	out, err := json.Marshal(group)
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, json.Unmarshal(out, &m))

	assert.Equal(t, float64(3), m["clusterCount"])
}

func TestRequiredReadOnlyFieldsUnmarshaledFromResponses(t *testing.T) {
	payload := `{"clusterCount":2,"created":"2026-01-01T00:00:00Z","name":"my-project","orgId":"5a0a1e7e0f2912c554080adc"}`

	var group admin.Group
	require.NoError(t, json.Unmarshal([]byte(payload), &group))

	assert.Equal(t, int64(2), group.ClusterCount)
	assert.Equal(t, time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), group.Created)
	assert.Equal(t, "my-project", group.Name)
}
