package main

import (
  "fmt"
  "time"
  "errors"
  "math"

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

  fmt.Printf("Starting monitor with interval %v seconds\n", INTERVAL)
  for true {
    fmt.Println(time.Now().Format("[15:04:05]"))

    doPing()

    time.Sleep(INTERVAL * time.Second)
  }
}

func doPing() {
  err := pinger.Run() // Blocks until finished.
  if err != nil {
    fmt.Printf("error: %s\n", err.Error())

    if !isDown {
      startTime = time.Now()
      isDown = true
    }
    elapsed := time.Since(startTime).Seconds()
    elapsedText := fmt.Sprintf("%.f seconds", math.Mod(elapsed, 60))
    if elapsed > 60.0 {
      elapsedText = fmt.Sprintf("%.f minutes %s", elapsed/60, elapsedText)
    }

    fmt.Printf("down for %s\n", elapsedText)
  } else {
    if isDown {
      startTime = time.Now()
      isDown = false
    }
    elapsed := time.Since(startTime).Seconds()
    elapsedText := fmt.Sprintf("%.f seconds", math.Mod(elapsed, 60))
    if elapsed > 60.0 {
      elapsedText = fmt.Sprintf("%.f minutes %s", elapsed/60, elapsedText)
    }

    stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
    fmt.Printf("avg speed = %v\n", stats.AvgRtt)
    fmt.Printf("up for %s\n", elapsedText)
  }
}
