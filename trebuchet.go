package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"log"
	"strings"
	"strconv"
)

var nonNumericRegex = regexp.MustCompile(`[^0-9]+`)

func clearString(str string) string {
	return nonNumericRegex.ReplaceAllString(str, "")
}

func main() {

	content, error := ioutil.ReadFile("./test.txt")
	if error != nil {

		log.Fatal(error)
	}

	str := string(content)

	lines := strings.Split(str, "\n")
	ans:= 0
	for _, line := range lines {
		if len(line) == 0{
			continue 
		}
		res:= clearString(line)
		res = string(res[0]) + string(res[len(res) - 1])
		i, err := strconv.Atoi(res)
    if err != nil {
        // ... handle error
        panic(err)
    }

		ans += i
	}

	fmt.Println(ans)
}