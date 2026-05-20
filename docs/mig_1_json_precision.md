# JSON Number Precision: Migration Guide

This guide describes a breaking change introduced in `v20250312020` to fix silent integer precision loss in dynamic JSON fields. If you are using `v20250312019` or earlier, the old `float64` behavior still applies.

## What changed

The SDK's response decoder previously called `json.Unmarshal`, which maps all JSON numbers in untyped fields to `float64`. `float64` has a 53-bit mantissa, so any integer with absolute value above 2^53 (9007199254740992) is silently rounded on decode. A read-modify-write cycle would write the corrupted value back to Atlas with no error.

The decoder now uses `json.Decoder` with `UseNumber()`. JSON numbers in fields typed `any`, `[]any`, or `map[string]any` now decode as `json.Number` instead of `float64`, preserving the exact wire representation.

Fields with explicit numeric types (`int32`, `int64`, `float64`, etc.) are unaffected.

## Which fields are affected

Any model field typed `any`, `[]any`, or `map[string]any`. Examples:

- `StreamsProcessor.Pipeline` (`*[]any`) — stream processing pipeline stages
- `SearchMappings.Dynamic` (`any`) and `SearchMappings.Fields` (`*map[string]any`)
- `AtlasSearchAnalyzer`, `VectorSearchIndexDefinition`, `IndexOptions`, and related search model fields
- `ApiError.Parameters` and `ApiError.BadRequestDetail`

## Migration

If your code type-asserts a value from one of these fields to `float64`, you must update it to handle `json.Number`.

### Before

```go
pipeline := proc.GetPipeline()
stage := pipeline[0].(map[string]any)
match := stage["$match"].(map[string]any)
tenantID := match["tenantId"].(float64) // panics after upgrade
```

### After

```go
import "encoding/json"

pipeline := proc.GetPipeline()
stage := pipeline[0].(map[string]any)
match := stage["$match"].(map[string]any)
num := match["tenantId"].(json.Number)

// Convert to int64 for exact integer semantics:
tenantID, err := num.Int64()

// Or to float64 when you know the value fits in 53 bits:
approx, err := num.Float64()

// Or keep the string representation:
raw := num.String()
```

### Safe helper for mixed-type fields

When a field can hold either integers or floats, check with a type switch:

```go
switch v := match["value"].(type) {
case json.Number:
    i, err := v.Int64()    // prefer for integers
    f, err := v.Float64()  // fallback for floats
case string:
    // string value
}
```

## Number decoding behavior change and its effects

The SDK change is in deserialization: the decoder now produces `json.Number("10.0")` instead of `float64(10)` for a JSON `10.0` in a dynamic field. This has two downstream effects.

**Re-serialization**: `float64(10)` marshaled back to `"10"` (Go normalizes float64 to its shortest representation), but `json.Number("10.0")` marshals back to `"10.0"` (the wire representation is preserved). Code that deserializes a response and re-serializes it to a string will now produce `"10.0"` where it previously produced `"10"`.

**Direct comparison**: `float64(10.0) == float64(10)` was true before, but `json.Number("10.0") == json.Number("10")` is false after, since `json.Number` is just a string type.

Both effects surface in JSON marshaling round-trips where the Atlas API returns `10.0` for a value the client originally sent as `10`.

If you need `10` and `10.0` to compare as equal, normalize the string representation first by stripping a purely-zero fractional part, then compare:

```go
import (
    "encoding/json"
    "strings"
)

func normalizeNumber(n json.Number) json.Number {
    s := string(n)
    if idx := strings.IndexByte(s, '.'); idx != -1 && strings.TrimLeft(s[idx+1:], "0") == "" {
        return json.Number(s[:idx])
    }
    return n
}

func numbersEqual(a, b json.Number) bool {
    return normalizeNumber(a) == normalizeNumber(b)
}

// numbersEqual(json.Number("10"), json.Number("10.0")) == true
// numbersEqual(json.Number("9007199254740993"), json.Number("9007199254740992")) == false
```

## Testing your migration

Add round-trip tests with integers above 2^53 in dynamic fields to catch precision loss:

```go
const largeInt = "9007199254740993" // 2^53 + 1
payload := fmt.Sprintf(`{"pipeline":[{"$match":{"tenantId":%s}}]}`, largeInt)

var p admin.StreamsProcessor
require.NoError(t, json.Unmarshal([]byte(payload), &p))

stage := p.GetPipeline()[0].(map[string]any)
match := stage["$match"].(map[string]any)
num, ok := match["tenantId"].(json.Number)
require.True(t, ok)
require.Equal(t, largeInt, num.String())
```
