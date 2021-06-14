package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/danielAang/html_tree/link"
)

func main() {
	site := flag.String("site", "https://www.google.com", "Website to parse all links")
	flag.Parse()
	resp, err := http.Get(*site)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		r := strings.NewReader(string(body))
		links, err := link.Parse(r)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", links)
	} else {
		fmt.Println("Unable to get site:", site)
	}
}
