package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

//type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)

// URL deduplicate
func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

func createWorker(in chan Request, out chan ParseResult, r ReadyNotifier)  {
	go func() {
		for {
			r.WorkerReady(in)
			request := <- in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}


