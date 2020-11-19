package cache

import (
	"fmt"
	"testing"
)

func TestStoreIncr(t *testing.T) {
	ok := StoreIncr("test", 2)
	fmt.Println(ok)
	ok = StoreIncr("test", 2)
	fmt.Println(ok)
	ok = StoreIncr("test", 2)
	fmt.Println(ok)
	ok = StoreIncr("test", 2)
	fmt.Println(ok)
	ok = StoreIncr("test", 2)
	fmt.Println(ok)
	ok = StoreDecr("test")
	fmt.Println(ok)
	ok = StoreDecr("test")
	fmt.Println(ok)
	ok = StoreIncr("test", 2)
	fmt.Println(ok)
}
