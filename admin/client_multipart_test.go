package admin

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func newMultipartTestClient() *APIClient {
	cfg := &Configuration{
		HTTPClient: http.DefaultClient,
		Servers: ServerConfigurations{
			{
				URL: "https://cloud.mongodb.com",
			},
		},
		UserAgent: "multipart-test",
	}
	return NewAPIClient(cfg)
}

func TestAddFile_WritesFilePayload(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "payload.txt")
	payload := "multipart payload"
	if err := os.WriteFile(filePath, []byte(payload), 0o600); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	err := addFile(w, "upload", filePath)
	if err != nil {
		t.Fatalf("expected addFile to succeed, got error: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("failed to close multipart writer: %v", err)
	}

	got := body.String()
	if !strings.Contains(got, `name="upload"; filename="payload.txt"`) {
		t.Fatalf("multipart body missing file part header, got: %q", got)
	}
	if !strings.Contains(got, payload) {
		t.Fatalf("multipart body missing file payload, got: %q", got)
	}
}

func TestPrepareRequest_MultipartAtFileSuccess(t *testing.T) {
	client := newMultipartTestClient()

	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "payload.txt")
	payload := "multipart payload"
	if err := os.WriteFile(filePath, []byte(payload), 0o600); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	headerParams := map[string]string{
		"Content-Type": "multipart/form-data",
	}
	formParams := url.Values{}
	formParams.Add("@upload", filePath)

	req, err := client.prepareRequest(
		context.Background(),
		"https://cloud.mongodb.com/api/atlas/v2/fake-endpoint",
		http.MethodPost,
		nil,
		headerParams,
		url.Values{},
		formParams,
		nil,
	)
	if err != nil {
		t.Fatalf("expected prepareRequest multipart path to succeed, got: %v", err)
	}
	if req == nil {
		t.Fatal("expected non-nil request")
	}

	if ctype := req.Header.Get("Content-Type"); !strings.HasPrefix(ctype, "multipart/form-data; boundary=") {
		t.Fatalf("unexpected Content-Type header: %q", ctype)
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("failed to read request body: %v", err)
	}
	bodyText := string(bodyBytes)
	if !strings.Contains(bodyText, `name="upload"; filename="payload.txt"`) {
		t.Fatalf("request body missing file part header, got: %q", bodyText)
	}
	if !strings.Contains(bodyText, payload) {
		t.Fatalf("request body missing file payload, got: %q", bodyText)
	}
}
