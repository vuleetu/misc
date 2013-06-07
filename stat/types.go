package gostat

import (
    "github.com/vuleetu/web"
)

type Stat struct {
    Name string
    Description string
    Stream bool
    Fun Callback
}

type StatInfo struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Stream bool `json:"stream"`
}

type Callback func(*web.Context)

type processStatInfo struct {
    CPUUsage string `json:"cpu"`
    MemUsage string `json:"mem"`
    VirtMem string `json:"virt"`
    RssMem string `json:"rss"`
    Routine int `json:"routine"`
}
