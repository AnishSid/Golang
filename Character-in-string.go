/*Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian
*/
package main

import (
"fmt"
"strings"
)

func main() {
  var str1 string
  fmt.Print("Enter a string: ")
  fmt.Scan(&str1)
  var str = strings.ToLower(str1)


  if strings.Contains(str, "a") && strings.HasSuffix(str, "n") && strings.HasPrefix(str, "i") {
    fmt.Println("Found!")
    } else {
      fmt.Println("Not Found!")
    }
  }

/*package main

import (
	"bufio"
	"os"
	"strings"
)

func isIANFound(someStr string) string {
	contains := strings.Contains(strings.ToLower(someStr), "i") &&
		strings.Contains(strings.ToLower(someStr), "a") &&
		strings.Contains(strings.ToLower(someStr), "n")
	if contains {
		return "Found!"
	}
	return "Not Found!"
}

func main() {
	var someLine string
	print("Enter some name:")
	userInput := bufio.NewReader(os.Stdin)
	someLine, err := userInput.ReadString('\n')
	if err != nil {
		print(err)
	}
	println(isIANFound(someLine))
}
*/
