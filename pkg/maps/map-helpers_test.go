package maps

import (
	"reflect"
	"testing"
)

func TestEmptyMap(t *testing.T) {
	result := SplitMapIntoNumberOfChunks[string](3, map[string]string{})

	if len(result) != 1 {
		t.Errorf("Expected 1 chunks, got %d", len(result))
	}

	for _, chunk := range result {
		if len(chunk) != 0 {
			t.Errorf("Expected empty chunk, got %v", chunk)
		}
	}
}

func TestSingleChunk(t *testing.T) {
	inputMap := map[string]string{"a": "apple", "b": "banana"}
	result := SplitMapIntoNumberOfChunks[string](1, inputMap)
	if len(result) != 1 || len(result[0]) != len(inputMap) {
		t.Errorf("Expected single chunk with %d elements, got %d chunks with %d elements", len(inputMap), len(result), len(result[0]))
	}
}

func TestMultipleChunksEvenDistribution(t *testing.T) {
	inputMap := map[string]string{"a": "apple", "b": "banana", "c": "cherry", "d": "date"}
	result := SplitMapIntoNumberOfChunks[string](2, inputMap)
	if len(result) != 2 || len(result[0]) != 2 || len(result[1]) != 2 {
		t.Errorf("Expected 2 chunks with 2 elements each, got %d chunks with %d and %d elements", len(result), len(result[0]), len(result[1]))
	}
}

func TestMultipleChunksUnevenDistribution(t *testing.T) {
	inputMap := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	result := SplitMapIntoNumberOfChunks[string](2, inputMap)
	if len(result) != 2 || len(result[0]) < 1 || len(result[1]) < 1 {
		t.Errorf("Expected 2 chunks with at least 1 element each, got %d chunks with %d and %d elements", len(result), len(result[0]), len(result[1]))
	}
}

func TestNumberOfChunksGreaterThanMapSize(t *testing.T) {
	inputMap := map[string]string{"a": "apple", "b": "banana"}
	result := SplitMapIntoNumberOfChunks[string](5, inputMap)
	if len(result) != len(inputMap) {
		t.Errorf("Expected %d chunks, got %d", len(inputMap), len(result))
	}
	nonEmptyCount := 0
	for _, chunk := range result {
		if len(chunk) > 0 {
			nonEmptyCount++
		}
	}
	if nonEmptyCount != len(inputMap) {
		t.Errorf("Expected %d non-empty chunks, got %d", len(inputMap), nonEmptyCount)
	}
}

func TestNumberOfChunksOutput(t *testing.T) {
	tests := []struct {
		name                    string
		numberOfChunksRequested int
		numberOfChunksExpected  int
		originalMap             map[string]string
	}{
		{"EmptyMap", 3, 1, map[string]string{}},
		{"SingleChunk", 1, 1, map[string]string{"a": "apple"}},
		{"EvenDistribution", 2, 2, map[string]string{"a": "apple", "b": "banana"}},
		{"UnevenDistribution", 2, 2, map[string]string{"a": "apple", "b": "banana", "c": "cherry"}},
		{"ChunksGreaterThanSize", 5, 2, map[string]string{"a": "apple", "b": "banana"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitMapIntoNumberOfChunks[string](tt.numberOfChunksRequested, tt.originalMap)
			if len(result) != tt.numberOfChunksExpected {
				t.Errorf("For test '%s', requested %d chunks, expected %d chunks, got %d", tt.name, tt.numberOfChunksRequested, tt.numberOfChunksExpected, len(result))
			}
		})
	}
}

func TestSplitMapWithMaxChunkSize(t *testing.T) {
	tests := []struct {
		name         string
		maxChunkSize int
		originalMap  map[int]string
		want         []map[int]string
	}{
		{
			name:         "empty map",
			maxChunkSize: 2,
			originalMap:  map[int]string{},
			want:         []map[int]string{},
		},
		{
			name:         "single chunk",
			maxChunkSize: 5,
			originalMap: map[int]string{
				1: "a", 2: "b", 3: "c",
			},
			want: []map[int]string{
				{1: "a", 2: "b", 3: "c"},
			},
		},
		{
			name:         "multiple chunks",
			maxChunkSize: 2,
			originalMap: map[int]string{
				1: "a", 2: "b", 3: "c", 4: "d",
			},
			want: []map[int]string{
				{1: "a", 2: "b"},
				{3: "c", 4: "d"},
			},
		},
		{
			name:         "uneven chunks",
			maxChunkSize: 3,
			originalMap: map[int]string{
				1: "a", 2: "b", 3: "c", 4: "d",
			},
			want: []map[int]string{
				{1: "a", 2: "b", 3: "c"},
				{4: "d"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitMapWithMaxChunkSize(tt.maxChunkSize, tt.originalMap)
			flattenedGot := flattenMapSlices(got)
			flattenedWant := flattenMapSlices(tt.want)

			if !reflect.DeepEqual(flattenedGot, flattenedWant) {
				t.Errorf("Flattened maps do not match.\nGot: %v\nWant: %v", flattenedGot, flattenedWant)
			}
		})
	}
}

func flattenMapSlices[K comparable, V any](slices []map[K]V) map[K]V {
	flatMap := make(map[K]V)
	for _, m := range slices {
		for k, v := range m {
			flatMap[k] = v
		}
	}
	return flatMap
}
