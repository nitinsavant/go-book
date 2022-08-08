package examples

import (
	"fmt"
	"io"
	"net/http"
)
import "os"

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			return
		}

		text, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return
		}

		fmt.Printf("%s", text)
	}
}
