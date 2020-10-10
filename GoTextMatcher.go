package GoTextMatcher

import "math"

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
