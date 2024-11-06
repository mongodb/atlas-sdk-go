package test

import (
	"testing"
	"time"

	"go.mongodb.org/atlas-sdk/v20241023002/admin"
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
