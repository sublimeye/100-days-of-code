package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const FIRST = 1
const LAST = 1930 // 1930
const CONCURRENCY = 10

const CACHE_DIR = "xkcd-cache"
const CACHE_FILE = "xkcd-index.json"

type result struct {
	index int
	res   *Xkcd
	err   error
}

type Xkcd struct {
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Url        string
}

// Collection of xkcd stories
type Collection struct {
	Items []*Xkcd `json:"items"`
}

func initIndex(fetch bool) Collection {
	if fetch {
		return buildIndex()
	}

	bIndex, err := readAndParseIndexFile()
	var index Collection

	if err == nil {
		if jerr := json.Unmarshal(bIndex, &index); jerr != nil {
			fmt.Println("Unable to parse index file. Building index from scratch")
			return buildIndex()
		}
		fmt.Println("Using index file")
		return index
	} else {
		fmt.Println("Error parsing index file", err)
		return buildIndex()
	}
}

func buildIndex() Collection {
	var urls []string
	var collection Collection
	start := time.Now()

	fmt.Println("fetching xkcd.com")
	for i := 1; i <= LAST; i++ {
		urls = append(urls, "https://xkcd.com/"+strconv.Itoa(i)+"/info.0.json")
	}

	results := fetchAll(urls, CONCURRENCY)
	for _, result := range results {
		if result.res != nil {
			collection.Items = append(collection.Items, result.res)
		}
	}

	saveAsJson(collection, CACHE_DIR)
	fmt.Printf("Fetched %d items in %.2fs\n", len(collection.Items), time.Since(start).Seconds())
	return collection
}

func saveAsJson(data Collection, dir string) {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir, 0755)
		} else {
			log.Println(err)
		}
	}

	path := fmt.Sprint(dir, "/", CACHE_FILE)
	os.Remove(path)

	//fmt.Println(data)

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println(b)

	ioutil.WriteFile(path, b, 0644)
}

func fetchAll(urls []string, concurrency int) []result {
	// buffered channel that will block at the concurrency limit
	semaphoreChan := make(chan struct{}, concurrency)

	// unbuffered channel -> will not block and collect http request results
	resultsChan := make(chan *result)

	defer func() {
		close(semaphoreChan)
		close(resultsChan)
	}()

	for i, url := range urls {

		// start a go routine with the index
		go func(i int, url string) {
			// this sends an empty struct into the semaphoreChan which
			// is basically saying add one to the limit, but when the
			// limit has been reached block until there is room
			semaphoreChan <- struct{}{}

			// send the request and put the response in a result struct
			// along with the index so we can sort them later along with
			// any error that might have happened
			res, err := http.Get(url)
			var item result
			var myresult Xkcd
			if res.StatusCode != 200 {
				fmt.Println("Bad status code", res.StatusCode, url)
				item = result{i, nil, err}
			} else if jsonErr := json.NewDecoder(res.Body).Decode(&myresult); jsonErr != nil {
				// res.Body.Close()
				// fmt.Println(jsonErr)
				// panic("JSON decoder unhandled error")
				item = result{i, nil, jsonErr}
			} else {
				myresult.Url = url
				item = result{i, &myresult, err}
			}

			res.Body.Close()

			// now we can send the result struct through the resultsChan
			resultsChan <- &item

			// once we're done it's we read from the semaphoreChan which
			// has the effect of removing one from the limit and allowing
			// another goroutine to start
			<-semaphoreChan

		}(i, url)
	}

	var results []result

	// start listening to resultsChan, once arrived append it to the result slice
	for {
		result := <-resultsChan
		results = append(results, *result)

		// stop when reached expected amount of urls
		if len(results) == len(urls) {
			break
		}
	}

	// we can sort here
	// sort.Slice()

	return results
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads index from file into memory
func readAndParseIndexFile() ([]byte, error) {
	const indexFilePath = CACHE_DIR + "/" + CACHE_FILE
	bIndex, err := ioutil.ReadFile(indexFilePath)
	if err != nil {
		fmt.Println("Index file was not found")
		return nil, err
	}

	if bIndex != nil {
		return bIndex, nil
	}

	return nil, errors.New("file is empty")
}

func searchIndex(text string, collection Collection) {
	var foundIndexes []int

	fmt.Printf("Searching for %q\n", text)
	for i, item := range collection.Items {
		if strings.Contains(item.Transcript, text) || strings.Contains(item.Title, text) {
			foundIndexes = append(foundIndexes, i)
			// 0 url title text
			fmt.Printf("%4.4d %.20s %.20s %.20s\n", i, item.Url, item.Title, item.Transcript)
		}
	}

	if len(foundIndexes) == 0 {
		fmt.Printf("No items found")
	}
}

func runXkcd() {
	text := flag.String("text", "", "text to search in xkcd index")
	fetch := flag.Bool("fetch", false, "force fetching index")
	flag.Parse()
	index := initIndex(*fetch)
	searchIndex(*text, index)
}
