package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	/* TODO:
		- from cli supply website url
		- get website body
		- search through body for links
		- search through body for paths
		- search through body for js extensions, .js
		- print these to screen
	*/
	var url string = "https://www.garthhumphreys.com"
	//fmt.Println(url)

	//match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//fmt.Println(match)

	var re = regexp.MustCompile(`(http(s?):)([/|.|\w|\s|-])*\.(?:js)`)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	html, err := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	for i, match := range re.FindAllString(string(html), -1) {
		fmt.Println(match, "found at index", i)
	}

	//fmt.Println(string(html))
}