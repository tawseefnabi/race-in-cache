package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	cache := run()
	cacheLen := len(cache.cache)
	pagesLen := (cache.pages).Len()
	fmt.Println("====", cacheLen, pagesLen)
}
