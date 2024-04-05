package maps

func SplitMapIntoNumberOfChunks[K comparable, V any](numberOfChunks int, originalMap map[K]V) []map[K]V {
	totalSize := len(originalMap)

	if totalSize == 0 || numberOfChunks <= 0 {
		return []map[K]V{originalMap}
	}

	if numberOfChunks > totalSize {
		numberOfChunks = totalSize
	}

	baseChunkSize := totalSize / numberOfChunks
	extraElements := totalSize % numberOfChunks

	chunks := make([]map[K]V, numberOfChunks)
	for i := range chunks {
		chunks[i] = make(map[K]V)
	}

	i := 0
	for key, value := range originalMap {
		chunkIndex := i / (baseChunkSize + 1)
		if chunkIndex >= extraElements {
			chunkIndex = extraElements + (i-extraElements*(baseChunkSize+1))/baseChunkSize
		}

		chunks[chunkIndex][key] = value
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
