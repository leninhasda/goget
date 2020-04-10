package main

import (
	"flag"
	"fmt"
	"goget/goget"
	"goget/parser"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var parallelLimit int
	flag.IntVar(&parallelLimit, "parallel", 10, "Number of parallel requests")
	flag.Parse()

	if parallelLimit <= 0 {
		log.Fatalf("error: invalid number %d, `parallel` should be positive integer and more than 0", parallelLimit)
	}

	urls := parser.ArgsURL(os.Args[1:]) // ignore file name

	if len(urls) == 0 {
		log.Fatalf("error: invlid url(s) given")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	getter := goget.New(client)

	limit := make(chan struct{}, parallelLimit)
	wg := &sync.WaitGroup{}

	for _, u := range urls {
		limit <- struct{}{}
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup) {
			defer func() {
				<-limit
				wg.Done()
			}()

			md5, err := getter.GetMd5Hash(url)
			if err != nil {
				fmt.Println("error: failed")
				return
			}

			fmt.Println(url, md5)
		}(u, wg)
	}

	wg.Wait()
}
