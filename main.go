package main

import (
	"log"
	"net/http"
	"time"

    "github.com/shirou/gopsutil/v4/mem"
    "github.com/shirou/gopsutil/v4/cpu"
)

func main() {
    http.HandleFunc("/events", sseHandler)

    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatalf("unable to start server: %s", err.Error())
    }
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")

    w.Header().Set("Access-Control-Allow-Origin", "*")

    memT := time.NewTicker(time.Second)
    defer memT.Stop()

    cpuT := time.NewTicker(time.Second)
    defer cpuT.Stop()

    for {
        select {
        case <- memT.C:
        // ..
        case <- cpuT.C:
        // ..
        }
    }
}

