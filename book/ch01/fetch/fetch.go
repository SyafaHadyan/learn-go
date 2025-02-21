package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		respond, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", error)
			os.Exit(1)
		}
		body, error := ioutil.ReadAll(respond.Body)
		respond.Body.Close()
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, error)
			os.Exit(1)
		}
		fmt.Printf("%s", body)

	}
}
