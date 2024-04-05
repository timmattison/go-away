package maps

func SplitMapIntoNumberOfChunks[K comparable, V any](numberOfChunks int, originalMap map[K]V) []map[K]V {
	totalSize := len(originalMap)
	chunkSize := totalSize / numberOfChunks

	if totalSize%numberOfChunks != 0 {
		chunkSize++
	}

	chunks := make([]map[K]V, 0, numberOfChunks)

	for i := 0; i < numberOfChunks; i++ {
		chunks = append(chunks, make(map[K]V))
	}

	i := 0

	for key, value := range originalMap {
		index := i / chunkSize

		if index >= numberOfChunks {
			index = numberOfChunks - 1
		}

		chunks[index][key] = value
		i++
	}

	return chunks
}

func SplitMapWithMaxChunkSize[K comparable, V any](maxChunkSize int, originalMap map[K]V) []map[K]V {
	totalSize := len(originalMap)
	numberOfChunks := totalSize / maxChunkSize

	if totalSize%maxChunkSize != 0 {
		numberOfChunks++
	}

	chunks := make([]map[K]V, 0, numberOfChunks)
	currentChunk := make(map[K]V)
	i := 0

	for key, value := range originalMap {
		if i >= maxChunkSize {
			chunks = append(chunks, currentChunk)
			currentChunk = make(map[K]V)
			i = 0
		}

		currentChunk[key] = value
		i++

		if i == maxChunkSize || len(currentChunk) == (totalSize%maxChunkSize) && len(chunks) == numberOfChunks-1 {
			chunks = append(chunks, currentChunk)
			currentChunk = make(map[K]V)
			i = 0
		}
	}

	if len(currentChunk) > 0 && len(chunks) < numberOfChunks {
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
