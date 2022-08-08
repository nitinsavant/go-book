package examples

import "fmt"
import "os"

func Echo2() {
	var result string
	for _, arg := range os.Args[1:] {
		result += arg + " "
	}
	fmt.Println(result)
}
