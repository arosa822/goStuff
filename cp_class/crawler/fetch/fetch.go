// fetch prints the content found at each specified URL

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
    "strings"
)

func checkAndModString(s string) (inputUrl string) {
    if strings.HasPrefix(s, "http://") != true {
        fmt.Printf("Modifying input string %s...\n", s)
        inputUrl = "http://" + s
    } else {
        inputUrl = s
    }
    return inputUrl
}

func main() {
	for _, url := range os.Args[1:] {
        url = checkAndModString(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
