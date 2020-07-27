package main
import "fmt"

func main() {
  for i := 1; i <= 100; i++{

    if i%4 == 0 && i%6 == 0 {
            fmt.Println("LinkedIn")
  } else if i%6 == 0 {
                fmt.Println("ln")
  } else if i%4 == 0 {
                    fmt.Println("Linked")
               }
}
}
