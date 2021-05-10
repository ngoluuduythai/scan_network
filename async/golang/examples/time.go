package main

import (
    "log"
    "time"
    "fmt"
)

func timeTrack(start time.Time, name string) {
        elapsed := time.Since(start)
        log.Printf("function %s took %s", name, elapsed)
}

func main() {
        defer timeTrack(time.Now(), "main")
        for i := 0; i < 100; i++ {
                fmt.Printf("I can count: %d\n", i)
        }
}


