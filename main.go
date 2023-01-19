package main

import (
  "fmt"
  "time"
  "errors"

  "github.com/go-ping/ping"
)

const (
  INTERVAL = 5
  TARGET = "www.google.com"
)

var pinger *ping.Pinger
var startTime time.Time
var isDown = false
var err error

func main() {
  fmt.Println("Resolving target IP...")
  for err = errors.New("dummy"); err != nil; {
    fmt.Println(time.Now().Format("[15:04:05]"))

    pinger, err = ping.NewPinger(TARGET)
    if err != nil {
      fmt.Println(err)
      time.Sleep(time.Second)
    } else {
      fmt.Println("succeed")
    }
  }

  pinger.Count = 1
  startTime = time.Now()

  // start monitorin
  fmt.Printf("Starting with interval %v seconds\n", INTERVAL)
  for true {
    fmt.Println(time.Now().Format("[15:04:05]"))

    doPing()

    time.Sleep(INTERVAL * time.Second)
  }
}

func doPing() {
  err := pinger.Run() // Blocks until finished.
  if err != nil {
    if !isDown {
      startTime = time.Now()
      isDown = true
    }
    elapsed := time.Since(startTime).Seconds()

    fmt.Printf("error: %s\n", err.Error())
    fmt.Printf("down for %.f seconds\n", elapsed)
  } else {
    if isDown {
      startTime = time.Now()
      isDown = false
    }
    elapsed := time.Since(startTime).Seconds()

    stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
    //fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
    //fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
      //stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
    //fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
      //stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
    fmt.Printf("avg speed = %v\n", stats.AvgRtt)
    fmt.Printf("up for %.f seconds\n", elapsed)
  }
}
