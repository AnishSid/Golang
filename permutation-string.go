
/*package main

import (
    "fmt"
)
 func join(ins []rune, c rune) (result []string) {
    for i := 0; i <= len(ins); i++ {
        result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
    }
    return
}

func permutations(testStr string) []string {
    var n func(testStr []rune, p []string) []string
    n = func(testStr []rune, p []string) []string{
        if len(testStr) == 0 {
            return p
        }else {
            result := []string{}
            for _, e := range p {
                result = append(result, join([]rune(e), testStr[0])...)
            }
            return n(testStr[1:], result)
        }
    }

    output := []rune(testStr)
    return n(output[1:], []string{string(output[0])})
}

func main() {
    d := permutations("ABCD")
    fmt.Println(d)
}
*/
package main

import "fmt"

func main() {
    data := "Abcd"
    dataRune := []rune(data)
    generatePermutation(dataRune, 0, len(dataRune)-1)
}

func generatePermutation(dataRune []rune, left, right int) {
    if left == right {
        fmt.Println(string(dataRune))
    } else {
        for i := left; i <= right; i++ {
            dataRune[left], dataRune[i] = dataRune[i], dataRune[left]
            generatePermutation(dataRune, left+1, right)
            dataRune[left], dataRune[i] = dataRune[i], dataRune[left]
        }
    }
}
