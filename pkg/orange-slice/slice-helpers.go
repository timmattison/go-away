package orange_slice

func SplitSliceIntoNumberOfChunks[T any](numberOfChunks int, originalSlice []T) [][]T {
	totalSize := len(originalSlice)

	if totalSize == 0 || numberOfChunks <= 0 {
		return [][]T{originalSlice}
	}

	if numberOfChunks > totalSize {
		numberOfChunks = totalSize
	}

	baseChunkSize := totalSize / numberOfChunks
	extraElements := totalSize % numberOfChunks

	chunks := make([][]T, numberOfChunks)

	currentChunk := 0
	currentChunkSize := 0

	for i, element := range originalSlice {
		if currentChunk < extraElements && currentChunkSize >= (baseChunkSize+1) ||
			currentChunk >= extraElements && currentChunkSize >= baseChunkSize {
			currentChunk++
			currentChunkSize = 0
		}

		chunks[currentChunk] = append(chunks[currentChunk], element)
		currentChunkSize++

		if i == totalSize-1 { // Last element
			break
		}
	}

	return chunks
}
