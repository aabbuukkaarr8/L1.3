package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	workers := flag.Int("workers", 3, "number of workers")
	flag.Parse()

	dataCh := make(chan int)
	var wg sync.WaitGroup

	
	for i := 1; i <= *workers; i++ {
		wg.Add(1) 
		go func(id int) {
			defer wg.Done()
			for v := range dataCh {
				fmt.Printf("worker %d: %d\n", id, v)
			}
		}(i)
	}

	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	
	go func() {
		i := 0
		for {
			i++
			dataCh <- i
			time.Sleep(100 * time.Millisecond) 
		}
	}()

	<-stop           
	close(dataCh)   
	wg.Wait()      
}
