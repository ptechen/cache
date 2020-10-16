package cache

import (
	"sync"
	"time"
)

var DataCache = new(sync.Map)

var Timeout int64

type Cache struct {
	Data      interface{}
	WriteTime int64
}

type Add struct {
	Key interface{}
	Val interface{}
}

func init() {
	Timeout = 60
	deleteCache()
}

func AddData(params []*Add) {
	wg := sync.WaitGroup{}
	wg.Add(len(params))
	for i := 0; i < len(params); i++ {
		go func(v *Add) {
			DataCache.Store(v.Key, &Cache {
				Data:      v.Val,
				WriteTime: time.Now().Unix(),
			})
		}(params[i])
		wg.Done()
	}
	wg.Wait()
}

func deleteCache() {
	go func() {
		timer := time.NewTicker(time.Minute)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				DataCache.Range(rangeDelete)
			}
		}
	}()
}

func rangeDelete(key interface{}, value interface{}) bool {
	curVal := value.(*Cache)
	if time.Now().Unix()-curVal.WriteTime > Timeout {
		DataCache.Delete(key)
	}
	return true
}
