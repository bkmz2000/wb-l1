package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func worker(c chan string, stop chan os.Signal, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Stopping ", i)

	ok := true
	for ok {
		select {
		case s := <-c:
			fmt.Println("gorutine ", i, "got : ", s)
		case sig := <-stop:
			fmt.Println("gorutine ", i, "is closed")
			stop <- sig
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	data := make(chan string, 10)
	stop := make(chan os.Signal, 10)

	workersCount := 4

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(data, stop, i, &wg)
	}

	go func() {
		for {
			select {
			case data <- generateRandomString(10):
			case sig := <-stop:
				fmt.Println("sender is stoped")
				stop <- sig
				return
			}
		}
	}()

	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	wg.Wait()
	close(data)
	close(stop)

	fmt.Println("Terminated")
}
