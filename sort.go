package gointsorter

import (
	"sync"
	"time"
)

func Sort(arr []int) []int {
	var wg sync.WaitGroup
	var resCh = make(chan int)

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(arr[i]) * 10 * time.Millisecond)
			resCh <- arr[i]
		}(i)
	}

	go func() {
		wg.Wait()
		close(resCh)
	}()

	var res []int

	for n := range resCh {
		res = append(res, n)
	}

	return res
}
