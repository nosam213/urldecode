package main

import (
	"log"
	"net/url"
	"os"
)

func UrlDecode(x string) string {
	z, err := url.QueryUnescape(x)
	if err != nil { log.Fatal(err) }
	return z
}

func main() {
	if len(os.Args) > 1 {
		println(UrlDecode(os.Args[1]))
	} else {
		os.Exit(0)
	}
}
