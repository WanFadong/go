package model

import (
	"errors"
	"sync"

	"time"

	"github.com/qiniu/xlog.v1"
)

var ErrFinished = errors.New("produce finished")
var DEFAULT_INTERVAL = 100000

type ProducerConsumer struct {
	produce  func() (interface{}, error) // 断点放在produce/consume里面
	consume  func(interface{}) error     // 限速放在这里做
	num      int                         // 消费者数量
	interval int
	buf      chan interface{}
	stop     chan struct{}
}

func NewProducerConsumer(produce func() (interface{}, error), consume func(interface{}) error, num int, interval int) (pc *ProducerConsumer) {

	buf := make(chan interface{}, num)
	stop := make(chan struct{})
	if interval <= 0 {
		interval = DEFAULT_INTERVAL
	}
	pc = &ProducerConsumer{produce: produce, consume: consume, num: num, interval: interval, buf: buf, stop: stop}
	return
}

// 断点处理：
// 1. produce失败，produce记录失败的位置。consume会处理已经produce的数据，然后退出。
// 2. consume失败，consume记录失败的位置。produce退出。
// produce返回ErrFinished，表示结束
func (pc *ProducerConsumer) Run() {
	xl := xlog.NewWith("Run ProducerConsumer")
	wg := sync.WaitGroup{}
	wg.Add(1 + pc.num)
	go func() {
		defer wg.Done()

		num := 0
		start := time.Now()
		for {
			data, err := pc.produce()
			if err != nil {
				if err == ErrFinished {
					xl.Info("producer exit because of finished")
					close(pc.buf)
				} else {
					xl.Info("producer exit because of err")
					close(pc.buf)
				}
				return
			}

			select {
			case <-pc.stop:
				xl.Info("producer exit because of consumer err")
				return
			case pc.buf <- data:
				num++
				if num%pc.interval == 0 {
					xl.Infof("total num:%v, total time:%v", num, time.Now().Sub(start))
				}
				continue
			}
		}
	}()

	for i := 0; i < pc.num; i++ {
		go func(i int) {
			defer wg.Done()
			for {
				select {
				// 除非produce失败的时候，consume也失败了，否则会处理完所有produce出来的数据
				case <-pc.stop:
					xl.Info("consumer exit because of consumer err", i)
					return
				case data, ok := <-pc.buf:
					if !ok {
						xl.Info("consumer exit because of finished or producer err", i)
						// 结束
						return
					}
					err := pc.consume(data)
					if err != nil {
						xl.Info("consumer exit because of err", i)
						pc.consumeFail() // 出错
						return
					}
				}
			}
		}(i)
	}

	wg.Wait()
}

func (pc *ProducerConsumer) consumeFail() {
	select {
	case <-pc.stop:
		// 已经关闭
	default:
		close(pc.stop)
	}
}
