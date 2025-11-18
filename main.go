package gokeys

import (
	"maps"
	"slices"

	"github.com/samber/lo"
)

func getKeysWithPreAllocate(m map[string]string) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func getKeysWithoutPreAllocate(m map[string]string) []string {
	keys := make([]string, 0)

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func getKeysWithCollection(m map[string]string) []string {
	return slices.Collect(maps.Keys(m))
}

func getKeysWithLo(m map[string]string) []string {
	return lo.Keys(m)
}
