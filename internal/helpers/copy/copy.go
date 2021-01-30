package copy

// Float32 Returns the pointer of a float32 value.
func Float32(v float32) *float32 {
	return &v
}

// Float64 Returns the pointer of a float64 value.
func Float64(v float64) *float64 {
	return &v
}

// Int Returns the pointer of a int value.
func Int(v int) *int {
	return &v
}

// Int32 Returns the pointer of a int32 value (rare but here).
func Int32(v int32) *int32 {
	return &v
}

// Int64 Returns the pointer of a int64 value (rare but here).
func Int64(v int64) *int64 {
	return &v
}

// String Returns the pointer of a string value.
func String(v string) *string {
	return &v
}