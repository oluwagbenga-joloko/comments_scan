// package main prints the 4 least used words in the url comments body from this
// "https://jsonplaceholder.typicode.com/comments"
package main

import (
	"fmt"

	c "github.com/oluwagbenga-joloko/comments_scan/comment_tool"
)

func main() {
	comments := c.GetComments("https://jsonplaceholder.typicode.com/comments")
	words := c.ProcessComments(comments)

	sortedWordCounts := c.SortWords(words)
	fmt.Println("four least Used words and count")
	for i := 0; i < 4; i++ {
		w := sortedWordCounts[i]
		fmt.Printf("\t%s : %v \n", w.Word, w.Count)
	}
}
