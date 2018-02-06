package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var f = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	f(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	f(r.MatchString("peach"))
	f(r.FindString("peach punch"))

	f(r.FindStringIndex("peach punch"))

	f(r.FindStringSubmatch("peach punch"))

	f(r.FindStringSubmatchIndex("peach punch"))

	f(r.FindAllString("peach punch pinch", -1))

	f(r.FindAllStringSubmatch("peach punch pinch", -1))

	f(r.FindAllString("peach punch pinch", 2))
	f(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	f(r)

	f(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	f(string(out))
}
