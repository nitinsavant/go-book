package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	tasks := []string{"nitin", "gajendra", "savant", "cassie", "utt"}

	results := make(chan string)
	go processTasks(tasks, results)

	fmt.Println(time.Now().Sub(now))
	printResults(results)
}
func processTasks(tasks []string, results chan<- string) {
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go process(task, results, &wg)
	}

	wg.Wait()
	close(results)

}

func process(task string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	results <- task
}

func printResults(results <-chan string) {
	time.Sleep(1 * time.Second)
	for result := range results {
		fmt.Println(result)
	}
}
