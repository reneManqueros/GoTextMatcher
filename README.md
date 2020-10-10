# Go Text Matcher
Go Text Matcher is a simple library with fuzzy text scoring/matching functions.

Currently it supports cosine similarity only.

My plan is to add support for:
* Levenshtein
* N-gram
* Hamming
* Others!

### Installation
```sh
$ go get github.com/reneManqueros/GoTextMatcher
```

### Usage
For now only Cosine distance is available:
```go
GoTextMatcher.CosineSimilarity([]byte("Hello gopher"), []byte("Hello gopher"))
```

Look at example/main.go

