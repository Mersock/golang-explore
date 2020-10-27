package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://stackoverflow.com",
	}
	c := make(chan string)

	for _, link := range links {
		go chkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			chkLink(link, c)
		}(l)
	}
}

func chkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link

}
