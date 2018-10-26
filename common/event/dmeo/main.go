package main

import (
	"github.com/ethereum/go-ethereum/event"
	"sync"
	"fmt"
)

type someEvent struct {
	I int
}

func main() {
	var feed event.Feed
	var wg sync.WaitGroup
	ch := make(chan someEvent)
	sub := feed.Subscribe(ch)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range ch {
			fmt.Printf("Received: %#v\n", event.I)
		}
		sub.Unsubscribe()
		fmt.Println("done")
	}()
	feed.Send(someEvent{5})
	feed.Send(someEvent{10})
	feed.Send(someEvent{7})
	feed.Send(someEvent{14})
	close(ch)
	wg.Wait()
}
