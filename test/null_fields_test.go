package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312028/admin"
)

// TestNullFieldsHonouredInNestedValueStructs verifies that SetXxxNil() still
// produces an explicit JSON null when the model is nested by value inside
// another model. A nested struct value is not addressable inside
// marshalWithNullFields, so a pointer-receiver MarshalJSON is skipped and the
// null is silently dropped.
func TestNullFieldsHonouredInNestedValueStructs(t *testing.T) {
	tests := []struct {
		name   string
		direct func() any
		nested func() any
		want   string
	}{
		{
			name: "BackupRestoreJob.Delivery",
			direct: func() any {
				d := admin.BackupRestoreJobDelivery{}
				d.SetMaxDownloadsNil()
				return &d
			},
			nested: func() any {
				d := admin.BackupRestoreJobDelivery{}
				d.SetMaxDownloadsNil()
				return &admin.BackupRestoreJob{Delivery: d}
			},
			want: `"maxDownloads":null`,
		},
		{
			name: "SearchIndexUpdateRequest.Definition",
			direct: func() any {
				d := admin.SearchIndexUpdateRequestDefinition{}
				d.SetAnalyzerNil()
				return &d
			},
			nested: func() any {
				d := admin.SearchIndexUpdateRequestDefinition{}
				d.SetAnalyzerNil()
				return &admin.SearchIndexUpdateRequest{Definition: d}
			},
			want: `"analyzer":null`,
		},
		{
			name: "OrgUserRequest.Roles",
			direct: func() any {
				r := admin.OrgUserRolesRequest{}
				r.SetGroupRoleAssignmentsNil()
				return &r
			},
			nested: func() any {
				r := admin.OrgUserRolesRequest{}
				r.SetGroupRoleAssignmentsNil()
				return &admin.OrgUserRequest{Roles: r}
			},
			want: `"groupRoleAssignments":null`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			direct, err := json.Marshal(tc.direct())
			require.NoError(t, err)
			assert.Contains(t, string(direct), tc.want, "top-level null missing")

			nested, err := json.Marshal(tc.nested())
			require.NoError(t, err)
			assert.Contains(t, string(nested), tc.want, "nested value struct dropped the explicit null")
		})
	}
}
