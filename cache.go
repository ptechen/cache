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

func init() {
	Timeout = 60
	deleteCache()
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
