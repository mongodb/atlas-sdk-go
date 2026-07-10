package test

import (
	"encoding/json"
	"testing"
	"time"

	"go.mongodb.org/atlas-sdk/v20250312021/admin"
)

type testCase[T comparable] struct {
	name         string
	val          T
	defaultValue T
	wantNil      bool
}

//nolint:thelper // used to minimize duplication in test
func assertSetOrNil[T comparable](t *testing.T, tc testCase[T]) {
	t.Run(tc.name, func(t *testing.T) {
		ptr := admin.SetOrNil(tc.val, tc.defaultValue)
		if gotNil := ptr == nil; gotNil != tc.wantNil {
			t.Errorf("got nil %t, want %t", gotNil, tc.wantNil)
		}
	})
}

func TestSetOrNil(t *testing.T) {
	assertSetOrNil(t, testCase[int]{
		name:         "non default int",
		val:          1,
		defaultValue: 0,
		wantNil:      false,
	})

	assertSetOrNil(t, testCase[int]{
		name:         "default int",
		val:          0,
		defaultValue: 0,
		wantNil:      true,
	})

	assertSetOrNil(t, testCase[string]{
		name:         "non default string",
		val:          "hello",
		defaultValue: "",
		wantNil:      false,
	})

	assertSetOrNil(t, testCase[string]{
		name:         "default string",
		val:          "",
		defaultValue: "",
		wantNil:      true,
	})
}

func TestStringToTime(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string

		want    time.Time
		wantErr string
	}{
		{
			name:  "valid time",
			input: "2023-07-18T16:12:23Z",
			want: time.Date(
				2023, 7, 18,
				16, 12, 23, 0,
				time.UTC,
			),
		},
		{
			name:  "valid time with millis",
			input: "2023-07-18T16:12:23.456Z",
			want: time.Date(
				2023, 7, 18,
				16, 12, 23, 456_000_000,
				time.UTC,
			),
		},
		{
			name:    "invalid time",
			input:   "invalid",
			wantErr: `parsing time "invalid" as "2006-01-02T15:04:05.999999999Z07:00": cannot parse "invalid" as "2006"`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotErr := ""
			got, err := admin.StringToTime(tc.input)
			if err != nil {
				gotErr = err.Error()
			}

			if gotErr != tc.wantErr {
				t.Errorf("want error %q, got %q", tc.wantErr, gotErr)
			}
			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

func unmarshalToMap(t *testing.T, b []byte) map[string]any {
	t.Helper()
	var out map[string]any
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}
	return out
}

func TestNullFieldsMarshalJSON(t *testing.T) {
	t.Run("empty NullFields marshals normally", func(t *testing.T) {
		entry := admin.NewNetworkPermissionEntry()
		entry.SetIpAddress("1.2.3.4")
		entry.SetComment("a comment")

		b, err := json.Marshal(entry)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := unmarshalToMap(t, b)
		if got["ipAddress"] != "1.2.3.4" {
			t.Errorf("want ipAddress %q, got %v", "1.2.3.4", got["ipAddress"])
		}
		if got["comment"] != "a comment" {
			t.Errorf("want comment %q, got %v", "a comment", got["comment"])
		}
	})

	t.Run("single null field overrides its value", func(t *testing.T) {
		entry := admin.NewNetworkPermissionEntry()
		entry.SetIpAddress("1.2.3.4")
		entry.SetComment("a comment")
		entry.SetCommentNil()

		b, err := json.Marshal(entry)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := unmarshalToMap(t, b)
		if v, ok := got["comment"]; !ok || v != nil {
			t.Errorf("want comment null, got %v (present=%t)", v, ok)
		}
		if got["ipAddress"] != "1.2.3.4" {
			t.Errorf("want ipAddress %q, got %v", "1.2.3.4", got["ipAddress"])
		}
	})

	t.Run("multiple null fields are all overridden", func(t *testing.T) {
		entry := admin.NewNetworkPermissionEntry()
		entry.SetIpAddress("1.2.3.4")
		entry.SetComment("a comment")
		entry.SetCommentNil()
		entry.SetCidrBlockNil()

		b, err := json.Marshal(entry)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := unmarshalToMap(t, b)
		if v, ok := got["comment"]; !ok || v != nil {
			t.Errorf("want comment null, got %v (present=%t)", v, ok)
		}
		if v, ok := got["cidrBlock"]; !ok || v != nil {
			t.Errorf("want cidrBlock null, got %v (present=%t)", v, ok)
		}
	})

	t.Run("unknown field name in NullFields is ignored", func(t *testing.T) {
		entry := admin.NewNetworkPermissionEntry()
		entry.SetIpAddress("1.2.3.4")
		entry.NullFields = append(entry.NullFields, "NotARealField")

		b, err := json.Marshal(entry)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := unmarshalToMap(t, b)
		if _, ok := got["NotARealField"]; ok {
			t.Errorf("did not expect unknown field to appear in output, got %v", got)
		}
		if got["ipAddress"] != "1.2.3.4" {
			t.Errorf("want ipAddress %q, got %v", "1.2.3.4", got["ipAddress"])
		}
	})

	t.Run("setting a value after SetXxxNil clears the null override", func(t *testing.T) {
		entry := admin.NewNetworkPermissionEntry()
		entry.SetCommentNil()
		entry.SetComment("a comment")

		b, err := json.Marshal(entry)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := unmarshalToMap(t, b)
		if got["comment"] != "a comment" {
			t.Errorf("want comment %q, got %v", "a comment", got["comment"])
		}
	})

	t.Run("calling SetXxxNil twice does not duplicate the field name", func(t *testing.T) {
		entry := admin.NewNetworkPermissionEntry()
		entry.SetCommentNil()
		entry.SetCommentNil()

		count := 0
		for _, f := range entry.NullFields {
			if f == "Comment" {
				count++
			}
		}
		if count != 1 {
			t.Errorf("want Comment to appear once in NullFields, got %d (%v)", count, entry.NullFields)
		}
	})

	t.Run("array-typed fields also support SetXxxNil", func(t *testing.T) {
		entry := admin.NewNetworkPermissionEntry()
		entry.SetLinks([]admin.Link{})
		entry.SetLinksNil()

		b, err := json.Marshal(entry)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := unmarshalToMap(t, b)
		if v, ok := got["links"]; !ok || v != nil {
			t.Errorf("want links null, got %v (present=%t)", v, ok)
		}
	})
}
