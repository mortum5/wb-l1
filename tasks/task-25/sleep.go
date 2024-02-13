package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	wg.Done()

}

func sleepAfter(d time.Duration) {
	defer timeTrack(time.Now(), "sleep with time.After")
	<-time.After(d)
}

func sleepTimer(d time.Duration) {
	defer timeTrack(time.Now(), "sleep with time.Timer")
	<-time.NewTimer(d).C
}

func main() {
	d := time.Second * 5
	wg.Add(2)

	go sleepAfter(d)
	go sleepTimer(d)

	wg.Wait()
}
