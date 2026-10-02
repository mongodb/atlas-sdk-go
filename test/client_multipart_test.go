package test

import (
	"context"
	"io"
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
