package main

import (
        "fmt"

        "github.com/shirou/gopsutil/cpu"
        "github.com/shirou/gopsutil/host"
        "github.com/shirou/gopsutil/load"
        "github.com/shirou/gopsutil/mem"
)

func main() {
        v, _ := mem.VirtualMemory()
        h, _ := host.Info()
        l, _ := load.Avg()
        c, _ := cpu.Times(true)

        // almost every return value is a struct
        fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

        // convert to JSON. String() is also implemented
        fmt.Println(v.UsedPercent)

        // convert to JSON. String() is also implemented
        fmt.Println(h)
        fmt.Println(l)
        fmt.Println(c)
}
