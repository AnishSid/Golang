package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func first50Char(s string) string {
	runes := []rune(s)
	return string(runes[0:50])
}

func main() {
	var fileName string
	nameSlice := make([]Name, 0)
	var nameObj Name

	fmt.Print("Enter filename: ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	filescan := bufio.NewScanner(file)
	for filescan.Scan() {
		s := strings.Split(filescan.Text(), " ")
		if len(s[0]) > 50 {
			s[0] = first50Char(s[0])
		}
		if len(s[1]) > 50 {
			s[1] = first50Char(s[1])
		}

		nameObj.fname, nameObj.lname = s[0], s[1]
		nameSlice = append(nameSlice, nameObj)
	}

	for _, v := range nameSlice {
		fmt.Println(v.fname, v.lname)
	}

	if err := filescan.Err(); err != nil {
		log.Fatal(err)
	}
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	firstName, lastName string
}

func main() {

	Data := make([]Name, 0, 100000)

	fmt.Println("Enter input file path: ")
	var filePath string
	_, _ = fmt.Scan(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File opening failed! Make sure correct path format!")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		name := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		Data = append(Data, Name{name[0], name[1]})
	}

	fmt.Println("File read successfully...")

	for _, name := range Data {
		fmt.Println(name.firstName, name.lastName)
	}

}
*/
