package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/zxcfer/newz/crawler"
)

var header string

func init() {
	flag.StringVar(&header, "H", "", "Specify a header line. If none specified, then first line is used as a header line.")
	flag.Parse()
}

func StartCrawler() {
	out, err := crawler.Crawl()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}

// Generate a function that return square of a number
