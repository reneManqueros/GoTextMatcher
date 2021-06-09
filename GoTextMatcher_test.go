package GoTextMatcher

import (
	"testing"
)

const helloGopher = "Hello Gopher"
const goodbyeGopher = "Goodbye Gopher"
const longPhrase = "to be or not to be"

func TestCosineSimilarityMatch(t *testing.T) {
	similarity := CosineSimilarity([]byte(helloGopher), []byte(helloGopher))

	if similarity < 1 {
		t.Fatalf(`CosineSimilarity %s != %s`, helloGopher, helloGopher)
	}
}

func TestCosineSimilarityNoMatch(t *testing.T) {
	similarity := CosineSimilarity([]byte(helloGopher), []byte(goodbyeGopher))

	if similarity == 1 {
		t.Fatalf(`CosineSimilarity %s == %s`, helloGopher, helloGopher)
	}
}

func TestHammingDistanceMatch(t *testing.T) {
	distance := HammingDistance([]byte(helloGopher), []byte(helloGopher))
	if distance != 1 {
		t.Fatalf(`HammingDistance != 1, %s - %s`, helloGopher, helloGopher)
	}
}

func TestHammingDistanceNoMatch(t *testing.T) {
	distance := HammingDistance([]byte(helloGopher), []byte(goodbyeGopher))
	if distance == 1 {
		t.Fatalf(`HammingDistance != 1, %s - %s`, helloGopher, goodbyeGopher)
	}
}

func TestLevenshteinMatch(t *testing.T) {
	distance := Levenshtein([]byte(helloGopher), []byte(helloGopher))
	if distance != 1 {
		t.Fatalf(`Levenshtein != 1, %s - %s`, helloGopher, helloGopher)
	}
}

func TestLevenshteinNoMatch(t *testing.T) {
	distance := Levenshtein([]byte(helloGopher), []byte(goodbyeGopher))
	if distance == 1 {
		t.Fatalf(`Levenshtein == 1, %s - %s`, helloGopher, goodbyeGopher)
	}
}

func TestNGramsMatch(t *testing.T) {
	ngrams := Ngrams(longPhrase, 2)
	ngramCount, ok := ngrams["To be"]
	if ok == false {
		t.Fatalf("gram not found")
	}
	if ok == true && ngramCount != 2 {
		t.Fatalf("incorrect gram count")
	}
}
