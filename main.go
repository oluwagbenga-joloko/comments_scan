// package main prints the 4 least used words in the url comments body from this
// "https://jsonplaceholder.typicode.com/comments"
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

//WordCount contains the word and its count
type WordCount struct {
	word  string
	count int
}

// Comment contains a comment
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	comments := GetComments("https://jsonplaceholder.typicode.com/comments")
	words := ProcessComments(comments)

	sortedWordCounts := SortWords(words)
	fmt.Println("four least Used words and count")
	for i := 0; i < 4; i++ {
		w := sortedWordCounts[i]
		fmt.Printf("\t%s : %v \n", w.word, w.count)
	}
}

// GetComments gets the comment from the url and returns the comments
func GetComments(url string) []Comment {
	comments := []Comment{}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &comments)
	if err != nil {
		log.Fatal(err)
	}
	return comments
}

// ProcessComments processes the comments and returns the a map of words and their count
func ProcessComments(comments []Comment) map[string]int {
	words := map[string]int{}
	for _, comment := range comments {
		for _, word := range strings.Fields(comment.Body) {
			_, ok := words[word]
			if ok {
				words[word]++

			} else {
				words[word] = 1
			}
		}
	}
	return words
}

// SortWords sorts the word in desending order and returns a slice of wordCount
func SortWords(words map[string]int) []WordCount {
	l := len(words)
	wordCounts := make([]WordCount, l, l)
	{
		i := 0
		for k, v := range words {
			wordCounts[i] = WordCount{k, v}
			i++
		}
	}
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count < wordCounts[j].count
	})
	return wordCounts
}
