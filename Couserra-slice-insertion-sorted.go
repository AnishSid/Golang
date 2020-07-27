//https://stackoverflow.com/questions/50884144/how-to-fill-a-slice-with-scan-values
package main

import (
  "fmt"
  "sort"
  )


func input(mySlice []int, err error) []int {
    if err != nil {
        return mySlice
    }
    var d int
    n, err := fmt.Scanf("%d", &d)
    if n == 1 {
        mySlice = append(mySlice, d)
    }
    return input(mySlice, err)
}

func main() {
    fmt.Println("Enter input:")
    mySlice := input([]int{}, nil)
    sort.Ints(mySlice)
    fmt.Println("Sorted_slice:", mySlice)
}

/*
package main

import (
    "fmt"
    "sort"
    "strconv"
)

func main() {
    var s []int = make([]int, 3)
    var in string
    fmt.Println("Please enter an interger(X to exit):")
    for true {
        fmt.Scanln(&in)
        if in == "X"{
            break
        }
        ap,err:=strconv.Atoi(in)
        if err != nil {
            fmt.Println("Wrong input")
            continue
        }
        s = append(s, ap)
        sort.Ints(s[:])
        fmt.Println(s)
        fmt.Println("Please again enter an interger(X to exit):")
	}
}
*/
