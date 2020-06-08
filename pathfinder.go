package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
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

	siteURL := flag.String("siteURL", "https://juice-shop.herokuapp.com", "a website or webpage")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("expected 'site url'")
		os.Exit(1)
	}

	if *siteURL != "" {
		fmt.Println("site url:", *siteURL)
		findFiles(*siteURL)
	} else {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func findFiles(siteurl string) {
	url := siteurl
	// fmt.Println(url)

	// match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	// fmt.Println(match)

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
		if strings.Contains(match, url) {
			fmt.Println(match, "found at index", i)
		}
	}

	// fmt.Println(string(html))
}