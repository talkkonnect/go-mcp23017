package main

import (
    "log"
    "github.com/talkkonnect/go-mcp23017"
    "time"
)


func main() {
    d, err := mcp23017.Open(1, 0)
    if err != nil {
        log.Println(err)
    }

    defer d.Close()

    for i := 0; i < 16; i++ {
        err := d.PinMode(uint8(i), mcp23017.OUTPUT)
        if err != nil {
            log.Println(err)
        }
    }

    for {
        for i := 0; i < 16; i++ {
        log.Println("Turning On ",i)
            err = d.DigitalWrite(uint8(i), mcp23017.HIGH)
            if err != nil {
                log.Println(err)
            }
            time.Sleep(time.Millisecond * 500)
        }
        time.Sleep(time.Second)

        for i := 0; i < 16; i++ {
        log.Println("Turning Off ",i)
            err = d.DigitalWrite(uint8(i), mcp23017.LOW)
            if err != nil {
                log.Println(err)
            }
            time.Sleep(time.Millisecond * 500)
        }
    }
}


