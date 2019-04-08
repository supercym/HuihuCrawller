package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

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
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}


	for {
		result := <- out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item}()
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier)  {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//循环等待
//解析完种子URL后得到的一堆request都在result里面，假定有100个request，
//在这100个request分发完之前，这个result不会变
//直到100个request分发完，result才更新为刚刚从out中取出的那个
//但是只有10个worker，这10个worker干完活后想把各自的result送进out，
//或许某一个送进去了，但是之前那个有100个request的result还没有消耗掉，
//所以engine不会去取out里面的东西，result也就不会更新为out里面的那个了
//
//
// 之前是worker共用一个输入，
// 在Submit中使用go routine之后，Submit函数只需开启一个go routine，然后就可以返回了，
//开go routine是很快的事情，它立刻就返回了
//Submit在很短时间内开启了100个go routine（大致这么理解），这100个go routine一直在监听in
//一旦in空着就会送request进去，然后这个go routine结束
//也因为Submit开启这100个go routine很快，所以这个含有100个request和item的result很快就被处理完了
//engine也能及时去取out里面新产生的result

