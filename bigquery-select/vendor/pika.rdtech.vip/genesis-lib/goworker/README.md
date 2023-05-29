goworker
===

## Core

reference: https://github.com/ivpusic/grpool

## Method

* func NewPool(numWorkers int, jobQueueLen int) *Pool
* func DoJob(callback func([]interface{})) Job
* func DoJobParams(callback func([]interface{}), i ...interface{}) Job
* func (p *Pool) Wait()
* func (p *Pool) Release()
* func (p *Pool) GracefulStop()
* func (p *Pool) StopAddJob()
* func (p *Pool) WorkingJobCount() int64 // 正在執行中的任務數量
* func (p *Pool) JobQueueLen() int // 隊列中的任務數量

## Example

[Example link](expamles/base.go#L25)

讓pool可以伸縮
max pool
min pool
life time

如果pool沒有新的worker
那就再產一個新的 if worker num <= max
新的worker會有一個計時的life time
如果到時間也沒人用就會去死
有人用life time就會刷新一次 像redis的key一樣

