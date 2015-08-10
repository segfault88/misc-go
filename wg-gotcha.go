package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doWork(i, wg)
	}

	log.Println("Starting to wait")
	wg.Wait()
	log.Println("Done")
}

func doWork(i int, wg sync.WaitGroup) {
	defer wg.Done()
	sleepTime := time.Duration(rand.Int31n(10000000))
	time.Sleep(sleepTime)
	log.Printf("%d slept for %d", i, sleepTime)
}
