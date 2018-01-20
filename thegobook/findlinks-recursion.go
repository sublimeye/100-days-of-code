package main

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	// printLinks(doc)
	printSummary(doc)
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

func printLinks(doc *html.Node) {
	fmt.Println("Document Links")
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func summary(elements map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	if n.FirstChild != nil {
		summary(elements, n.FirstChild)
	}

	if n.NextSibling != nil {
		summary(elements, n.NextSibling)
	}

	return elements
}

func printSummary(doc *html.Node) {
	fmt.Println("Document tags count")
	initMap := make(map[string]int)
	summaryMap := summary(initMap, doc)

	sValues := map[int]string{}
	sKeys := []int{}
	for key, val := range summaryMap {
		sValues[val] = key
		sKeys = append(sKeys, val)
	}

	sort.Ints(sKeys)

	for _, value := range sKeys {
		fmt.Println(sValues[value], value)
	}
}
