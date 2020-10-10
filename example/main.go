package main

import (
	"github.com/reneManqueros/GoTextMatcher"
	"log"
)

func main() {
	// Exact match
	log.Println(GoTextMatcher.CosineSimilarity([]byte("Hello gopher"), []byte("Hello gopher")))

	// Close match
	log.Println(GoTextMatcher.CosineSimilarity([]byte("Hello World"), []byte("Hi World")))

	// Not so close:
	log.Println(GoTextMatcher.CosineSimilarity([]byte("Hello World"), []byte("Hola Mundo")))

	// Not even close:
	log.Println(GoTextMatcher.CosineSimilarity([]byte("qwerty"), []byte("asdfg")))
}
