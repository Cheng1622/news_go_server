package spider

import (
	"fmt"
	"sync"
)

func SpiderNews() {
	wg := &sync.WaitGroup{}
	for k, v := range tabArr {
		wg.Add(1)
		go func(k string, v int) {
			defer wg.Done()
			get(k, v, 1)
		}(k, v)
	}

	wg.Wait()
	getlist()
	fmt.Println("spider is ok")
}
