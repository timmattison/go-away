package orange_slice

import (
	"reflect"
	"testing"
)

func TestSplitSliceIntoNumberOfChunksInt(t *testing.T) {
	tests := []struct {
		name           string
		numberOfChunks int
		originalSlice  []int
		want           [][]int
	}{
		{
			name:           "split int slice into valid chunks",
			numberOfChunks: 3,
			originalSlice:  []int{1, 2, 3, 4, 5, 6},
			want:           [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:           "split int slice into uneven chunks",
			numberOfChunks: 4,
			originalSlice:  []int{1, 2, 3, 4, 5, 6},
			want:           [][]int{{1, 2}, {3, 4}, {5}, {6}},
		},
		{
			name:           "empty slice",
			numberOfChunks: 3,
			originalSlice:  []int{},
			want:           [][]int{{}},
		},
		{
			name:           "number of chunks is 0",
			numberOfChunks: 0,
			originalSlice:  []int{1, 2, 3},
			want:           [][]int{{1, 2, 3}},
		},
		{
			name:           "number of chunks is negative",
			numberOfChunks: -1,
			originalSlice:  []int{1, 2, 3},
			want:           [][]int{{1, 2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitSliceIntoNumberOfChunks(tt.numberOfChunks, tt.originalSlice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSliceIntoNumberOfChunks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitSliceIntoNumberOfChunksString(t *testing.T) {
	tests := []struct {
		name           string
		numberOfChunks int
		originalSlice  []string
		want           [][]string
	}{
		{
			name:           "split string slice into more chunks than elements",
			numberOfChunks: 10,
			originalSlice:  []string{"a", "b", "c"},
			want:           [][]string{{"a"}, {"b"}, {"c"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitSliceIntoNumberOfChunks(tt.numberOfChunks, tt.originalSlice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSliceIntoNumberOfChunks() got = %v, want %v", got, tt.want)
			}
		})
	}
}
