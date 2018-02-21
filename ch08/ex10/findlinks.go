package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/relsa/exercise-gopl.io/ch08/ex10/links"
)

var cancel = make(chan struct{})

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url, cancel)
	if err != nil {
		log.Print(err)
	}
	return list
}

func crawlUnseenLinks(unseenLinks <-chan string, worklist chan<- []string, wg *sync.WaitGroup) {
	defer func() { wg.Done() }()
	for {
		select {
		case <-cancel:
			return
		case link := <-unseenLinks:
			foundLinks := crawl(link)
			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case <-cancel:
					return
				default:
					worklist <- foundLinks
				}
			}()
		}
	}
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	var wg sync.WaitGroup

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
		log.Println("canceled")

		wg.Wait()
		close(unseenLinks)
		close(worklist)
	}()

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go crawlUnseenLinks(unseenLinks, worklist, &wg)
	}

	seen := make(map[string]bool)
	for {
		select {
		case list, ok := <-worklist:
			if !ok {
				return // exit
			}
			for _, link := range list {
				select {
				case <-cancel:
					continue
				default:
					if !seen[link] {
						seen[link] = true
						unseenLinks <- link
					}
				}
			}
		case <-cancel:
			for range worklist {
				// flush worklist
			}
			for range unseenLinks {
				// flush unseenLinks
			}
		}
	}
}
