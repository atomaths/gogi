package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

func replace(src []byte) []byte {
	src, err := ioutil.ReadFile("replace.html")
	if err != nil {
		log.Print(err)
		return []byte{}
	}
	if src[len(src)-1] == '\n' {
		src = src[:len(src)-1]
	}
	return src
}

func main() {
	expr := "(?s)<!-- {{HEADER}} -->(.+?)<!-- {{/HEADER}} -->"
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}

	src, _ := ioutil.ReadFile("index.html")

	result := re.ReplaceAllFunc(src, replace)

	ioutil.WriteFile("result.html", result, 0644)
}
