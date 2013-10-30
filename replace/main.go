package main

import (
	"fmt"
	"log"
	"regexp"
)

var data = `
<h3>test</h3>
<!-- {{HEADER}} -->
content
<!-- {{/HEADER}} -->
<h3>end</h3>
`

func main() {
	expr := "content"
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}

	res := re.ReplaceAllLiteralString(data, "replaced!")
	fmt.Println(res)
}
