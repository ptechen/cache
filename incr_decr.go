package cache

func StoreIncr(key interface{}, maxVal ...int) (ok bool) {
	cache, ok := load(key)
	if !ok {
		Store(&Add{
			Key:     key,
			Val:     1,
		})
		return true
	}

	v, ok := cache.Data.(int)
	if !ok {
		return false
	}

	if len(maxVal) > 0 {
		if maxVal[0] >= v {
			Store(&Add{
				Key:     key,
				Val:     v + 1,
			})
		} else {
			return false
		}

	} else {
		Store(&Add{
			Key:     key,
			Val:     v + 1,
			TimeOut: cache.Timeout,
		})
	}
	return true
}

func StoreDecr(key interface{}, mixVal ...int) (ok bool) {
	cache, ok := load(key)
	if !ok {
		Store(&Add{
			Key:     key,
			Val:     -1,
		})
		return true
	}

	v, ok := cache.Data.(int)
	if !ok {
		return false
	}

	if len(mixVal) > 0 {
		if mixVal[0] <= v {
			Store(&Add{
				Key:     key,
				Val:     v - 1,
			})
		} else {
			return false
		}

	} else {
		Store(&Add{
			Key:     key,
			Val:     v - 1,
			TimeOut: cache.Timeout,
		})
	}

	return true
}
