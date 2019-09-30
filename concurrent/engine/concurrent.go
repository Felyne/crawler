package engine

import (
	"log"
	"time"

	"github.com/Felyne/crawler/concurrent/config"
)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan interface{}
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request //调度器分配一个chan作为worker的输入，可能多个worker共用一个输入
	Run()
}

//拆分出来便于createWorker()轻量级参数传递
type ReadyNotifier interface {
	WorkerReady(chan Request) //worker告知调度器输入已经准备好
}

//启动调度器，创建worker，接收worker的输出chan传来的结果再做处理
func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		tm := time.After(time.Duration(config.Timeout) * time.Second)
		var result ParseResult
		select {
		case <-tm:
			log.Printf("%ds timeout.No task execution!", config.Timeout)
			return
		case result = <-out:
		}
		//result := <-out

		for _, item := range result.Items {
			//ItemSaver消费item的速度远比生成快,所以同一时刻不会有太多的goroutine
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

// worker拿到一个输入chan ,放到调度器的worker队列等待调度
// 调度器拿出来发送一个request给它，任务完成后继续把输入chan放到worker队列等待调度
// worker公用一个输出chan
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

//去重
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
