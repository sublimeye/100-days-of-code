package main

import (
  "fmt"
  "net/http"
  "strconv"
  "time"
  "encoding/json"
  "os"
  "io/ioutil"
  "log"
)

const FIRST = 1

// const LAST = 1930
const LAST = 10

type result struct {
  index int
  res   *Xkcd
  err   error
}

type Xkcd struct {
  SafeTitle  string `json:"safe_title"`
  Title  string `json:"title"`
  Transcript, Img, Alt string
}

type Collection struct {
  Items []*Xkcd `json:"items"`
}

func initIndex() {
  buildIndex()
  // if !index => build
  // load index into memory
}

func buildIndex() {
  var urls []string
  var collection Collection
  start := time.Now()

  for i := 1; i <= LAST; i++ {
    urls = append(urls, "https://xkcd.com/"+strconv.Itoa(i)+"/info.0.json")
  }

  results := fetchAll(urls, LAST)
  for _, result := range results {
    collection.Items = append(collection.Items, result.res)
    //fmt.Printf("%v %.10s %.33s \n", result.index, result.res.Title, result.res.Transcript)
  }

  saveAsJson(collection, "xkcd-cache")
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func saveAsJson(data Collection, dir string) {
  if _, err := os.Stat(dir); err != nil {
    if os.IsNotExist(err) {
      os.Mkdir(dir, 0755)
    } else {
      log.Println(err)
    }
  }

  path := fmt.Sprint(dir, "/xkcd-index.json")
  os.Remove(path)

  //fmt.Println(data)

  b, err := json.Marshal(data)
  if err != nil { log.Println(err) }

  //fmt.Println(b)

  ioutil.WriteFile(path, b, 0644)
}

//func buildIndex2() {
//  start := time.Now()
//  ch := make(chan string)
//
//  for i := 1; i <= LAST; i++ {
//    url := "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
//    go fetch(url, ch) // start a goroutine
//  }
//
//  for i := 1; i <= LAST; i++ {
//    //fmt.Println(<-ch) // receive from channel
//  }
//
//  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
//  // boundaries hardcoded: 1-1930
//  // download concurrently (max of 10 concurrent connections)
//  // append to an array (don't sort)
//  // save as json with Items as array
//}

//func fetch(url string, ch chan<- string) {
//  start := time.Now()
//  resp, err := http.Get(url)
//
//  if err != nil {
//    ch <- fmt.Sprint(err) // send to channel
//    return
//  }
//
//  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
//  resp.Body.Close() // don't leak resources
//  if err != nil {
//    ch <- fmt.Sprintf("while reading %s: %v", url, err)
//    return
//  }
//  secs := time.Since(start).Seconds()
//  ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
//}

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

      var myresult Xkcd
      if jsonErr := json.NewDecoder(res.Body).Decode(&myresult); jsonErr != nil {
        res.Body.Close()
        fmt.Println(jsonErr)
        panic("JSON decoder unhandled error")
        //result := &result{i, nil, err}
      }

      result := &result{i, &myresult, err}

      res.Body.Close()

      // now we can send the result struct through the resultsChan
      resultsChan <- result

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

// Reads index from file into memory
func readIndex() {

}

func runXkcd() {
  fmt.Println("Running: xkcd")
  initIndex()
}
