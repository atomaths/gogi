package main

import (
	"fmt"
	"io/ioutil"
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

func replace(data []byte) []byte {
	fmt.Println(data)
	src, _ := ioutil.ReadFile("src.html")
	return src
}

func main() {
	expr := "(?s)<!-- {{HEADER}} -->(.+?)<!-- {{/HEADER}} -->"
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}

	dst, _ := ioutil.ReadFile("dst.html")

	result := re.ReplaceAllFunc(dst, replace)

	ioutil.WriteFile("result.html", result, 0644)
}
