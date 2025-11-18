package gokeys

import (
	"fmt"
	"testing"
)

func makeMap(n int) map[string]string {
	m := make(map[string]string, n)
	for i := range n {
		m[fmt.Sprintf("key-%d", i)] = "value"
	}
	return m
}

// Benchmark for the manual iteration approach
func BenchmarkGetKeysWithPreAllocate(b *testing.B) {
	m := makeMap(1000000)
	b.ReportAllocs()

	for b.Loop() {
		_ = getKeysWithPreAllocate(m)
	}
}

// Benchmark for maps.Keys + slices.Collect approach
func BenchmarkGetKeysWithCollection(b *testing.B) {
	m := makeMap(1000000)
	b.ReportAllocs()

	for b.Loop() {
		_ = getKeysWithCollection(m)
	}
}

func BenchmarkGetKeysWithoutPreAllocate(b *testing.B) {
	m := makeMap(1000000)
	b.ReportAllocs()

	for b.Loop() {
		_ = getKeysWithoutPreAllocate(m)
	}
}

// Benchmark for the Godash wrapper
func BenchmarkGetKeysWithLo(b *testing.B) {
	m := makeMap(1000000)
	b.ReportAllocs()

	for b.Loop() {
		_ = getKeysWithLo(m)
	}
}
