package cache

import (
	"fmt"
	"time"
)

type memCache struct {
	//最大内存
	maxMemorySize int64
	//最大内存字符串表示
	maxMemorySizeStr string
	//当前已使用内存
	currMemorySize int64
}

func NewMenCach() Cache {
	return &memCache{}
}

// size 1KB,100KB,1MB,2MB,1GB
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	fmt.Println(mc.maxMemorySize, mc.maxMemorySizeStr)
	return false
}
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	return false
}
func (mc *memCache) Get(key string) (interface{}, bool) {
	return nil, false
}
func (mc *memCache) Del(key string) bool {
	return false
}
func (mc *memCache) Exists(key string) bool {
	return false
}
func (mc *memCache) Flush() bool {
	return false
}
func (mc *memCache) Keys() int64 {
	return 0
}
