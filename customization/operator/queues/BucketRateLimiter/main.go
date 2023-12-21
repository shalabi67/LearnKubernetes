package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"k8s.io/client-go/util/workqueue"
)

func main() {
	stopCh := make(chan string)
	timeLayout := "2006-01-02:15:04:05.0000"
	limiter := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	//limiter := workqueue.NewRateLimitingQueue(&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)})

	length := 20
	chs := make([]chan string, length)
	for i := 0; i < length; i++ {
		chs[i] = make(chan string, 1)
		go func(taskId string, ch chan string) {
			item := "Task-" + taskId + time.Now().Format(timeLayout)
			log.Println(item + " Added.")
			limiter.AddRateLimited(item) // 添加会根据When() 延迟添加到工作队列中

		}(strconv.FormatInt(int64(i), 10), chs[i])

		go func() {
			for {
				key, quit := limiter.Get()
				if quit {
					return
				}
				log.Println(fmt.Sprintf("%s process done", key))
				defer limiter.Done(key)

			}
		}()
	}
	<-stopCh
}
