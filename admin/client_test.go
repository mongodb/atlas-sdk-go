package admin

import (
    "context"
    "net/http"
    "testing"
)

func TestPrepareRequest_WithRequestID(t *testing.T) {
    client := NewAPIClient(NewConfiguration())

    // Test with request ID in context
    ctx := context.WithValue(context.Background(), "X-Request-ID", "test-request-123")
    req, err := client.prepareRequest(ctx, "/api/test", "GET", nil, nil, nil, nil, nil)

    if err != nil {
        t.Fatalf("PrepareRequest failed: %v", err)
    }

    // Check that request ID header was added
    requestID := req.Header.Get("X-Request-ID")
    if requestID != "test-request-123" {
        t.Errorf("Expected X-Request-ID header 'test-request-123', got '%s'", requestID)
    }
}

func TestPrepareRequest_WithoutRequestID(t *testing.T) {
    client := NewAPIClient(NewConfiguration())

    // Test without request ID in context (should work as before)
    ctx := context.Background()
    req, err := client.prepareRequest(ctx, "/api/test", "GET", nil, nil, nil, nil, nil)

    if err != nil {
        t.Fatalf("PrepareRequest failed: %v", err)
    }

    // Check that no request ID header was added
    requestID := req.Header.Get("X-Request-ID")
    if requestID != "" {
        t.Errorf("Expected no X-Request-ID header, got '%s'", requestID)
    }
}