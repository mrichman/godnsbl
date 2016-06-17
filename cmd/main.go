package main

import (
	"encoding/json"

	"fmt"
	"log"
	"os"
	"sync"

	"github.com/mrichman/godnsbl"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please supply a domain name or IP address.")
		os.Exit(1)
	}

	ip := os.Args[1]

	wg := &sync.WaitGroup{}
	results := make([]godnsbl.Result, len(godnsbl.Blacklists))
	for i, source := range godnsbl.Blacklists {
		wg.Add(1)
		go func(i int, source string) {
			defer wg.Done()
			rbl := godnsbl.Lookup(source, ip)
			if len(rbl.Results) == 0 {
				results[i] = godnsbl.Result{}
			} else {
				results[i] = rbl.Results[0]
			}
		}(i, source)
	}

	wg.Wait()

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&results); err != nil {
		log.Println(err)
	}
}
