package dataexport

import (
	"currency-demo/model"
	"fmt"
	"sync"
)

func Export(accessLogChannel <-chan model.Data, wg *sync.WaitGroup) {
	var buffer []model.Data = make([]model.Data, 0)

	for {
		data, ok := <-accessLogChannel
		if !ok {
			break
		}

		buffer = append(buffer, data)

		if len(buffer) == 2000 {
			putToRedis(buffer)
		}

		// flush buffer
		buffer = nil

	}
	wg.Done()
}

func putToRedis(buffer []model.Data) {
	for _, v := range buffer {
		fmt.Println(v.IP, v.Bucket)
	}
}
