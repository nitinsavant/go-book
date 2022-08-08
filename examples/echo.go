package examples

import "fmt"
import "os"

func Echo() {
	var result string
	for i := 1; i < len(os.Args); i++ {
		result += os.Args[i] + " "
	}
	fmt.Println(result)
}
