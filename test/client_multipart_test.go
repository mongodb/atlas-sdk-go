package test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312026/admin"
)

const multipartEndpoint = "https://cloud.mongodb.com/api/atlas/v2/fake-endpoint"

func prepareMultipartRequest(t *testing.T, formParams url.Values) (*http.Request, error) {
	t.Helper()

	cfg := admin.NewConfiguration()
	cfg.HTTPClient = http.DefaultClient

	return admin.NewAPIClient(cfg).UntypedClient.PrepareRequest(
		context.Background(),
		multipartEndpoint,
		http.MethodPost,
		nil,
		map[string]string{"Content-Type": "multipart/form-data"},
		url.Values{},
		formParams,
		nil,
	)
}

func TestPrepareRequestMultipartFileUpload(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "payload.txt")
	payload := "multipart payload"
	require.NoError(t, os.WriteFile(filePath, []byte(payload), 0o600))

	formParams := url.Values{}
	formParams.Add("description", "some description")
	formParams.Add("@upload", filePath)

	req, err := prepareMultipartRequest(t, formParams)
	require.NoError(t, err)
	require.NotNil(t, req)

	contentType := req.Header.Get("Content-Type")
	boundary, ok := strings.CutPrefix(contentType, "multipart/form-data; boundary=")
	require.True(t, ok, "unexpected Content-Type header: %q", contentType)

	body, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	bodyText := string(body)

	assert.Contains(t, bodyText, `name="description"`)
	assert.Contains(t, bodyText, "some description")
	assert.Contains(t, bodyText, `name="upload"; filename="payload.txt"`)
	assert.Contains(t, bodyText, payload)
	assert.True(t, strings.HasSuffix(bodyText, "--"+boundary+"--\r\n"), "request body must end with the closing boundary")
}

// TestPrepareRequestMultipartContentLengthMatchesFinalizedBody checks that the
// closing boundary is written before Content-Length is measured.
func TestPrepareRequestMultipartContentLengthMatchesFinalizedBody(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "payload.txt")
	require.NoError(t, os.WriteFile(filePath, []byte("multipart payload"), 0o600))

	formParams := url.Values{}
	formParams.Add("@upload", filePath)

	req, err := prepareMultipartRequest(t, formParams)
	require.NoError(t, err)
	require.NotNil(t, req)

	body, err := io.ReadAll(req.Body)
	require.NoError(t, err)

	assert.Equal(t, int64(len(body)), req.ContentLength)
	assert.Equal(t, strconv.Itoa(len(body)), req.Header.Get("Content-Length"))
}

func TestPrepareRequestMultipartMissingFileReturnsError(t *testing.T) {
	formParams := url.Values{}
	formParams.Add("@upload", filepath.Join(t.TempDir(), "does-not-exist.txt"))

	req, err := prepareMultipartRequest(t, formParams)
	require.Error(t, err)
	assert.Nil(t, req)
	assert.True(t, os.IsNotExist(err), "expected the underlying file-open error to surface unmasked, got: %v", err)
}

func TestPrepareRequestMultipartEdgeCases(t *testing.T) {
	for _, tc := range []struct {
		name     string
		payloads []string
	}{
		{name: "fields only"},
		{name: "empty file", payloads: []string{""}},
		{name: "binary file", payloads: []string{"\x00\x01\xff\r\npayload\x00"}},
		{name: "multiple files", payloads: []string{"first payload", "second payload"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			formParams := url.Values{"tag": {"one", "two"}}
			dir := t.TempDir()

			var wantNames []string
			for i, payload := range tc.payloads {
				name := "payload-" + strconv.Itoa(i) + ".bin"
				path := filepath.Join(dir, name)
				require.NoError(t, os.WriteFile(path, []byte(payload), 0o600))

				wantNames = append(wantNames, name)
				formParams.Add("@upload", path)
			}

			req, err := prepareMultipartRequest(t, formParams)
			require.NoError(t, err)
			require.NotNil(t, req)
			defer req.Body.Close()

			body, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			assert.Equal(t, int64(len(body)), req.ContentLength)
			assert.Equal(t, strconv.Itoa(len(body)), req.Header.Get("Content-Length"))

			mediaType, params, err := mime.ParseMediaType(req.Header.Get("Content-Type"))
			require.NoError(t, err)
			assert.Equal(t, "multipart/form-data", mediaType)
			require.NotEmpty(t, params["boundary"])
			assert.True(t, bytes.HasSuffix(
				body,
				[]byte("\r\n--"+params["boundary"]+"--\r\n"),
			), "body must include its closing boundary")

			reader := multipart.NewReader(bytes.NewReader(body), params["boundary"])
			var gotNames, gotPayloads, gotTags []string

			for {
				part, err := reader.NextPart()
				if errors.Is(err, io.EOF) {
					break
				}
				require.NoError(t, err)

				payload, readErr := io.ReadAll(part)
				closeErr := part.Close()
				require.NoError(t, readErr)
				require.NoError(t, closeErr)

				switch part.FormName() {
				case "upload":
					// Inspect the raw filename parameter because FileName()
					// strips directories and could hide a leaked local path.
					_, disposition, err := mime.ParseMediaType(
						part.Header.Get("Content-Disposition"),
					)
					require.NoError(t, err)

					gotNames = append(gotNames, disposition["filename"])
					gotPayloads = append(gotPayloads, string(payload))

				case "tag":
					assert.Empty(t, part.FileName())
					gotTags = append(gotTags, string(payload))

				default:
					t.Fatalf("unexpected multipart field %q", part.FormName())
				}
			}

			assert.Equal(t, wantNames, gotNames)
			assert.Equal(t, tc.payloads, gotPayloads)
			assert.Equal(t, []string{"one", "two"}, gotTags)
		})
	}
}
