package admin

import (
    "context"
    "testing"
)

func TestApiError_GetContextualError(t *testing.T) {
    // Create a sample API error
    detail := "Project not found"
    reason := "Resource not found"
    apiError := &ApiError{
        Detail:    &detail,
        Error:     404,
        ErrorCode: "RESOURCE_NOT_FOUND",
        Reason:    &reason,
    }

    // Test with request ID in context
    ctx := context.WithValue(context.Background(), "X-Request-ID", "debug-456")
    contextualError := apiError.GetContextualError(ctx)

    expected := "404 RESOURCE_NOT_FOUND Resource not found: Project not found (RequestID: debug-456)"
    if contextualError != expected {
        t.Errorf("Expected: %s\nGot: %s", expected, contextualError)
    }

    // Test without request ID in context
    ctx = context.Background()
    contextualError = apiError.GetContextualError(ctx)

    expected = "404 RESOURCE_NOT_FOUND Resource not found: Project not found"
    if contextualError != expected {
        t.Errorf("Expected: %s\nGot: %s", expected, contextualError)
    }
}