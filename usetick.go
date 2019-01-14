// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import "time"
import "fmt"


func procticker(ticker *time.Ticker) {
	for t := range ticker.C {
		fmt.Printf("Tick at %s %d %d.%d\n", t, t.Unix(), t.UnixNano()/1000000000, t.UnixNano()%1000000000)
	}
}

func main() {

	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `range` builtin on the channel to iterate over
	// the values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	go procticker(ticker)

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.
	time.Sleep(2600 * time.Millisecond)
	ticker.Stop() 
	fmt.Println("Ticker to be reset")
	ticker = time.NewTicker(900 * time.Millisecond)
	go procticker(ticker)
	time.Sleep(5600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
