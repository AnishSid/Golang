package main

import (
	"fmt"
	"os"
	"encoding/json"
	"bufio"
)

func main() {
	mymap := make(map[string]string)

	var name string
	var address string

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name: ")
	in.Scan()
	name = in.Text()

	fmt.Print("Enter address: ")
	in.Scan()
	address = in.Text()

	mymap["name"] = name
	mymap["address"] = address

	myjson, err := json.Marshal(mymap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(myjson))
}
