package GoTextMatcher

import (
	"bytes"
	"math"
)

func CosineSimilarity(baseText []byte, comparisonText []byte) float64 {
	baseTextVector := make(map[byte]int)
	for _, character := range baseText {
		baseTextVector[character]++
	}

	comparisonTextVector := make(map[byte]int)
	for _, character := range comparisonText {
		comparisonTextVector[character]++
	}

	dotProduct := 0.0
	for k, v := range baseTextVector {
		dotProduct += float64(v) * float64(comparisonTextVector[k])
	}

	baseTextSum := 0.0
	for _, v := range baseTextVector {
		baseTextSum += math.Pow(float64(v), 2)
	}

	comparisonTextSum := 0.0
	for _, v := range comparisonTextVector {
		comparisonTextSum += math.Pow(float64(v), 2)
	}

	magnitude := math.Sqrt(baseTextSum) * math.Sqrt(comparisonTextSum)
	if magnitude == 0 {
		return 0.0
	}
	return dotProduct / magnitude
}

func HammingDistance(baseText []byte, comparisonText []byte) float64 {
	switch bytes.Compare(baseText, comparisonText) {
	case 0:
	case 1:
		temp := make([]byte, len(baseText))
		copy(temp, comparisonText)
		comparisonText = temp
	case -1:
		temp := make([]byte, len(comparisonText))
		copy(temp, baseText)
		baseText = temp
	}
	if len(baseText) != len(comparisonText) {
		return -1
	}
	count := 0
	for idx, byteFromBase := range baseText {
		byteFromComparison := comparisonText[idx]
		xor := byteFromBase ^ byteFromComparison
		for x := xor; x > 0; x >>= 1 {
			if int(x&1) == 1 {
				count++
			}
		}
	}
	if count == 0 {
		return 1
	}
	return float64(1) / float64(count)
}

func Levenshtein(baseText []byte, comparisonText []byte) float64 {
	matrix := initializeMatrix(len(baseText)+1, len(comparisonText)+1)
	for i := 0; i < len(baseText)+1; i++ {
		matrix[i][0] = i
	}
	for i := 0; i < len(comparisonText)+1; i++ {
		matrix[0][i] = i
	}
	for i := 0; i < len(baseText); i++ {
		for j := 0; j < len(comparisonText); j++ {
			edit := 0
			if baseText[i] != comparisonText[j] {
				edit = 1
			}
			matrix[i+1][j+1] = min(
				matrix[i][j+1]+1,
				matrix[i+1][j]+1,
				matrix[i][j]+edit,
			)
		}
	}
	distance := matrix[len(baseText)][len(comparisonText)]
	if distance == 0 {
		return 1
	}
	return float64(1) / float64(distance)
}

func min(items ...int) int {
	min := items[0]
	for _, elem := range items[1:] {
		if elem < min {
			min = elem
		}
	}
	return min
}

func initializeMatrix(rows int, columns int) [][]int {
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, columns)
	}
	return result
}
