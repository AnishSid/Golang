// Code example for https://stackoverflow.com/questions/58175413/text-parsing-of-log-file-in-go
//Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object,
//creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.

package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

const input = `
2019-09-30T04:17:02 - REQUEST-A
2019-09-30T04:18:02 - REQUEST-C
2019-09-30T04:19:02 - REQUEST-B
2019-09-30T04:20:02 - REQUEST-A
2019-09-30T04:21:02 - REQUEST-A
2019-09-30T04:22:02 - REQUEST-B
`

var scanner = bufio.NewScanner(strings.NewReader(input))

// Here comes the magic: We create an anonymous group containing all
// non-whitespace characters up to the first blank. Since it is
// a group, we can easily extract it later down the road.
var re *regexp.Regexp = regexp.MustCompile(`^(\S*) - REQUEST-A$`)

func main() {
	var res []string
	for scanner.Scan() {
		// We use re.FindStringSubmatch here, as it actually kills two
		// birds with one stone: We check wether it is REQUEST-A
		// and the anonymous group of the regexp contains what we are looking for.
		if res = re.FindStringSubmatch(scanner.Text()); len(res) > 0 {
			fmt.Println(res[1])
		}
	}
}
