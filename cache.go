package cache

import (
	"sync"
	"time"
)

var DataCache = new(sync.Map)

type Cache struct {
	Data      interface{}
	WriteTime int64
	Timeout   int64
}

type Add struct {
	Key     interface{}
	Val     interface{}
	TimeOut int64
}

func init() {
	deleteCache()
}

func Stores(params []*Add) {
	wg := sync.WaitGroup{}
	wg.Add(len(params))
	for i := 0; i < len(params); i++ {
		go func(v *Add) {
			DataCache.Store(v.Key, &Cache{
				Data:      v.Val,
				Timeout:   v.TimeOut,
				WriteTime: time.Now().Unix(),
			})
		}(params[i])
		wg.Done()
	}
	wg.Wait()
}

func Store(params *Add) {
	DataCache.Store(params.Key, &Cache{
		Data:      params.Val,
		WriteTime: time.Now().Unix(),
		Timeout:   params.TimeOut,
	})
}

func Load(key interface{}) (data interface{}, ok bool) {
	data, ok = DataCache.Load(key)
	if ok {
		return data.(*Cache).Data, ok
	}
	return nil, false
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
	if time.Now().Unix()-curVal.WriteTime > curVal.Timeout {
		DataCache.Delete(key)
	}
	return true
}
