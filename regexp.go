package main

import (
	"fmt"
	"regexp"
	"strings"
)

func testRegex() {
	expr := regexp.MustCompile("\\b\\w")

	transformation := func(s string) string {
		return strings.ToUpper(s)
	}

	text := "vitor bacelar de paula augusto"
	fmt.Println(transformation(text))
	fmt.Println(expr.ReplaceAllStringFunc(text, transformation))
}
