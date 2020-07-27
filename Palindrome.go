package main

import (
   "fmt"
   "unicode"
)

func isPalindrome(input string) bool {
  //base-case zero & 1 lenght input

    inputlength := len(input)

if inputlength == 0 || inputlength == 1 {
   return true
  }

  //invalidate palindrom assumption

  start := 0
  end := inputlength - 1

  for start < end {
    //get these as runes - golangs data-type type to represent Unicode

    lr, hr :=  rune(input[start]), rune(input[end])
    //move start end end inwards
    //invalid chars
    if !unicode.IsLetter(lr) {
      start++
      continue
    }
    if !unicode.IsLetter(hr) {
      end--
      continue
    }
    //check equality
    if unicode.ToLower(lr) != unicode.ToLower(hr) {
      return false
    }
    //advance
    start++
    end--
  }
  return true
}

func main() {
    fmt.Printf("race car! - %t\n", isPalindrome("race car!"))
}
