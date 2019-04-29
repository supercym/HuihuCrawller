package engine

import (
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request)  {
	requests := make([]Request, 0)
	for _, req := range seeds {
		requests = append(requests, req)
	}

	count := 0
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
				count++
				log.Printf("Got item #%d: %+v", count, item)
		}
	}
}




