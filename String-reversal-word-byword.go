/*Reversing a string by word is a similar process. First, we convert the string into an array of strings where each entry is a word. Next we apply the normal reverse loop to that array. Finally, we smush the results back together into a string that we can return to the caller
*/

package main

import (
	"fmt"
	"strings"
)

func reverse_words(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverse_words("the sky is blue"))
}
