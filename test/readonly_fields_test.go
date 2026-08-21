package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312023/admin"
)

// TestReadOnlyRequiredFieldsOmittedWhenUnset verifies that fields marked as both
// required and readOnly in the Atlas OpenAPI spec (e.g. Group.clusterCount) are
// omitted from serialized request bodies when they hold their zero value.
// Per OpenAPI 3.0 a readOnly property must not be sent in requests; sending the
// zero value can be rejected by the server, e.g. an empty string is not a valid
// enum value.
func TestReadOnlyRequiredFieldsOmittedWhenUnset(t *testing.T) {
	group := &admin.Group{}
	group.SetName("test-group")
	group.SetOrgId("65f2a1b3c4d5e6f7a8b9c0d1")

	raw, err := json.Marshal(group)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(raw, &payload))

	// readOnly required fields are omitted while unset
	assert.NotContains(t, payload, "clusterCount")
	// note: created is a time.Time struct, and encoding/json's omitempty cannot
	// omit zero-value structs, so it keeps being sent as "0001-01-01T00:00:00Z"
	// exactly as before this fix
	assert.Contains(t, payload, "created")
	// writable required fields are always sent, even with zero values
	assert.Contains(t, payload, "name")
	assert.Contains(t, payload, "orgId")
}

// TestReadOnlyRequiredStringFieldOmittedWhenUnset covers the string case:
// an unset readOnly required string (e.g. an enum such as
// ConnectedOrgConfig.orgId) must not be sent as "" in request bodies.
func TestReadOnlyRequiredStringFieldOmittedWhenUnset(t *testing.T) {
	config := &admin.ConnectedOrgConfig{}
	config.SetDomainRestrictionEnabled(false)

	raw, err := json.Marshal(config)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(raw, &payload))

	assert.NotContains(t, payload, "orgId")
	assert.Contains(t, payload, "domainRestrictionEnabled")
}

// TestReadOnlyRequiredFieldsSentWhenSet verifies that an explicitly set
// required readOnly field is still serialized.
func TestReadOnlyRequiredFieldsSentWhenSet(t *testing.T) {
	group := &admin.Group{}
	group.SetClusterCount(3)

	raw, err := json.Marshal(group)
	require.NoError(t, err)

	var payload map[string]any
	require.NoError(t, json.Unmarshal(raw, &payload))

	assert.Equal(t, float64(3), payload["clusterCount"])
}

// TestReadOnlyRequiredFieldsUnmarshal verifies that responses carrying
// required readOnly fields still populate them.
func TestReadOnlyRequiredFieldsUnmarshal(t *testing.T) {
	const response = `{"clusterCount":2,"created":"2026-08-21T10:00:00Z","name":"g","orgId":"o"}`

	var group admin.Group
	require.NoError(t, json.Unmarshal([]byte(response), &group))

	assert.Equal(t, int64(2), group.GetClusterCount())
	assert.Equal(t,
		time.Date(2026, 8, 21, 10, 0, 0, 0, time.UTC),
		group.GetCreated().UTC(),
	)
}
