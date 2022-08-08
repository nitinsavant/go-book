package scratch

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{"nitin", "gajendra", "savant", "cassie", "utt"}

	results := processTasks(tasks)

	printResults(results)
}

func processTasks(tasks []string) []string {
	var results []string
	for _, task := range tasks {
		time.Sleep(1 * time.Second)
		results = append(results, task)
	}
	return results
}

func printResults(results []string) {
	for _, result := range results {
		fmt.Println(result)
	}
}
