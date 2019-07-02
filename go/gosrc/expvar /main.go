package main

import (
	"expvar"
	"fmt"
	"runtime"
	"time"
)

// 开始时间
var start = time.Now()

// calculateUptime 计算运行时间
func calculateUptime() interface{} {
	return time.Since(start).String()
}

// currentGoVersion 当前 Golang 版本
func currentGoVersion() interface{} {
	return runtime.Version()
}

// getNumCPUs 获取 CPU 核心数量
func getNumCPUs() interface{} {
	return runtime.NumCPU()
}

// getGoOS 当前系统类型
func getGoOS() interface{} {
	return runtime.GOOS
}

// getNumGoroutins 当前 goroutine 数量
func getNumGoroutins() interface{} {
	return runtime.NumGoroutine()
}

// GetCurrentRunningStats 返回当前运行信息
func GetCurrentRunningStats() {
	expvar.Do(func(kv expvar.KeyValue) {
		fmt.Println(kv.Key, kv.Value)
	})
}

func init() {
	expvar.Publish("运行时间", expvar.Func(calculateUptime))
	expvar.Publish("version", expvar.Func(currentGoVersion))
	expvar.Publish("cores", expvar.Func(getNumCPUs))
	expvar.Publish("os", expvar.Func(getGoOS))
}

func main() {
	GetCurrentRunningStats()
	//http.ListenAndServe(":1234", nil)
}
