package engine

import "fmt"

//调度器
type SimpleConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)                        //用于将Request送给worker的channel,来一个request就创建一个groutine.
	ConfigerMasterWorkerChan(chan Request) //配置workChan,实现也就是个赋值操作。
}

//传进来url
func (e *SimpleConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	//把in给到workerChan
	e.Scheduler.ConfigerMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	//现在是由scheduler去把request提交
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out

		//将拿到的值的数量和其值打印出来
		for _, item := range result.Items {
			fmt.Printf("这是拿到的第#%d个数据。拿到item的值是%s:\n", itemCount, item)
			itemCount++
		}

		//将拿到的seed给到scheduler去提交
		for _, item := range result.Requests {
			e.Scheduler.Submit(item)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			res, err := Worker(request)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}
