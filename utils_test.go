package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestAnyToString(t *testing.T) {
	values := []any{
		"hello",
		123,
		123.456,
		true,
		nil,
		int8(8),
		int16(16),
		int32(32),
		int64(64),
		uint8(8),
		uint16(16),
		uint32(32),
		uint64(64),
		float32(32.32),
		time.Now(),
		[]byte("byte slice"),
		[]int{1, 2, 3},
		map[string]int{"key1": 1, "key2": 2},
		struct{ Name string }{"value"},
	}

	for _, v := range values {

		fmt.Printf("String: %s\n", AnyToString(v))

	}
}
