package application

import (
	"fmt"
	"runtime"
	"time"

	"mono.thienhang.com/pkg/utils"
)

var (
	startTime = time.Now()
)

type AppStatus struct {
	Uptime       string `json:"Uptime"`
	NumGoroutine int    `json:"NumGoroutine"`

	// General statistics.
	MemAllocated string `json:"MemAllocated"`
	MemTotal     string `json:"MemTotal"`
	MemSys       string `json:"MemSys"`
	Lookups      uint64 `json:"Lookups"`
	MemMallocs   uint64 `json:"MemMallocs"`
	MemFrees     uint64 `json:"MemFrees"`

	// Main allocation heap statistics.
	HeapAlloc    string `json:"HeapAlloc"`
	HeapSys      string `json:"HeapSys"`
	HeapIdle     string `json:"HeapIdle"`
	HeapInuse    string `json:"HeapInuse"`
	HeapReleased string `json:"HeapReleased"`
	HeapObjects  uint64 `json:"HeapObjects"`

	// Low-level fixed-size structure allocator statistics.
	StackInuse  string `json:"StackInuse"`
	StackSys    string `json:"StackSys"`
	MSpanInuse  string `json:"MSpanInuse"`
	MSpanSys    string `json:"MSpanSys"`
	MCacheInuse string `json:"MCacheInuse"`
	MCacheSys   string `json:"MCacheSys"`
	BuckHashSys string `json:"BuckHashSys"`
	GCSys       string `json:"GCSys"`
	OtherSys    string `json:"OtherSys"`

	// Garbage collector statistics.
	NextGC       string `json:"NextGC"`
	LastGC       string `json:"LastGC"`
	PauseTotalNs string `json:"PauseTotalNs"`
	PauseNs      string `json:"PauseNs"`
	NumGC        uint32 `json:"NumGC"`
}

func GetAppStatus() AppStatus {
	var app AppStatus
	// app.Uptime = utils.TimeSincePro(startTime, language.Lang[config.GetLanguage()])

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	app.NumGoroutine = runtime.NumGoroutine()

	app.MemAllocated = utils.FileSize(m.Alloc)
	app.MemTotal = utils.FileSize(m.TotalAlloc)
	app.MemSys = utils.FileSize(m.Sys)
	app.Lookups = m.Lookups
	app.MemMallocs = m.Mallocs
	app.MemFrees = m.Frees

	app.HeapAlloc = utils.FileSize(m.HeapAlloc)
	app.HeapSys = utils.FileSize(m.HeapSys)
	app.HeapIdle = utils.FileSize(m.HeapIdle)
	app.HeapInuse = utils.FileSize(m.HeapInuse)
	app.HeapReleased = utils.FileSize(m.HeapReleased)
	app.HeapObjects = m.HeapObjects

	app.StackInuse = utils.FileSize(m.StackInuse)
	app.StackSys = utils.FileSize(m.StackSys)
	app.MSpanInuse = utils.FileSize(m.MSpanInuse)
	app.MSpanSys = utils.FileSize(m.MSpanSys)
	app.MCacheInuse = utils.FileSize(m.MCacheInuse)
	app.MCacheSys = utils.FileSize(m.MCacheSys)
	app.BuckHashSys = utils.FileSize(m.BuckHashSys)
	app.GCSys = utils.FileSize(m.GCSys)
	app.OtherSys = utils.FileSize(m.OtherSys)

	app.NextGC = utils.FileSize(m.NextGC)
	app.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	app.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	app.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	app.NumGC = m.NumGC

	return app
}

type SysStatus struct {
	CpuLogicalCore int
	CpuCore        int
	OSPlatform     string
	OSFamily       string
	OSVersion      string
	Load1          float64
	Load5          float64
	Load15         float64
	MemTotal       string
	MemAvailable   string
	MemUsed        string
}
