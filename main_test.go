package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetComments(t *testing.T) {
	expectedComments := []Comment{
		Comment{PostID: 1, ID: 2, Name: "alan", Email: "Lew@alysha.tv", Body: "non et atque\noccaecati deserunt"},
		Comment{PostID: 3, ID: 4, Name: "mill", Email: "mill@alysha.tv", Body: "unde odit nobis qui voluptatem\nquia"},
		Comment{PostID: 6, ID: 9, Name: "mark", Email: "mark@alysha.tv", Body: "nobis qui voluptatem\nquia voluptas"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expectedComments)
	}))
	defer server.Close()

	actualComments := GetComments(server.URL)

	if l := len(actualComments); l != 3 {
		t.Errorf("wrong number of comments: got %v want %v",
			l, 3)
	}

	for i, expectedComment := range expectedComments {
		actualComment := actualComments[i]
		if expectedComment.Body != actualComment.Body {
			t.Errorf("wrong comment body: got %v want %v",
				actualComment.Body, expectedComment.Body)
		}

	}
}

func TestProcessComments(t *testing.T) {
	comments := []Comment{
		Comment{PostID: 1, ID: 2, Name: "alan", Email: "Lew@alysha.tv", Body: "Hello from the\nriver the \nthe"},
		Comment{PostID: 3, ID: 4, Name: "mill", Email: "mill@alysha.tv", Body: "the boy from the river"},
		Comment{PostID: 6, ID: 9, Name: "mark", Email: "mark@alysha.tv", Body: "river river river swims the from"},
		Comment{PostID: 6, ID: 9, Name: "mark", Email: "mark@alysha.tv", Body: "boy boy boy boy boy boy boy"},
	}

	expectedWords := map[string]int{
		"Hello": 1,
		"from":  3,
		"the":   6,
		"river": 5,
		"boy":   8,
		"swims": 1,
	}
	actualWords := ProcessComments(comments)
	if lActual, lExpected := len(actualWords), len(expectedWords); lActual != lExpected {
		t.Errorf("wrong number of comments: got %v want %v",
			lActual, lExpected)
	}

	for k, v := range expectedWords {
		if aV := actualWords[k]; aV != v {
			t.Errorf("wrong count for word %s: got %v want %v",
				k, aV, v)
		}

	}
}
func TestSortWords(t *testing.T) {
	words := map[string]int{
		"Hello": 2,
		"from":  3,
		"max":   999,
		"the":   6,
		"river": 5,
		"boy":   8,
		"swims": 90,
		"one":   1,
	}
	expected := []WordCount{
		WordCount{"one", 1},
		WordCount{"Hello", 2},
		WordCount{"from", 3},
		WordCount{"river", 5},
		WordCount{"the", 6},
		WordCount{"boy", 8},
		WordCount{"swims", 90},
		WordCount{"max", 999},
	}
	actual := SortWords(words)
	if lActual, lExpected := len(actual), len(expected); lActual != lExpected {
		t.Errorf("wrong number of comments: got %v want %v",
			lActual, lExpected)
	}

	for i, expectedWordCount := range expected {
		actualWordCount := actual[i]
		if actualWordCount.count != expectedWordCount.count {
			t.Errorf("wrong WordCount count: got %v want %v",
				actualWordCount.count, expectedWordCount.count)
		}
		if actualWordCount.word != expectedWordCount.word {
			t.Errorf("wrong WordCount word: got %v want %v",
				actualWordCount.word, expectedWordCount.word)
		}
	}
}
