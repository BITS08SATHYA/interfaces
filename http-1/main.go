package main

import "fmt"

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.bing.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		fmt.Println(link)
	}

}
