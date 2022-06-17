package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    end := false

    defer func() {
        if err := recover(); err != nil {
            fmt.Println("[main] panic occurred:", err)
        }
    }()

    watchDog := func() {
        defer func() {
            if err := recover(); err != nil {
                fmt.Println("[watchdog] panic occurred:", err)
            }
        }()
        // Watchdog
        fmt.Println("Watchdog will wake in 10 sec and end it...")

        /*
         Notes on Blocking call here
         In real life this would be a network connection perhaps. Its important
         to also close the connection on whatever end trigger and handle any error here.
         Though the Watchdog will end the main function; the go routine will execute one last time.
         To ensure it does not, Additionally try to kill the source of the blocking call.

         Another counter consideration to watchdog can be that if we can kill the Blocking call,
         do we still need the Watchdog?
        
         Yes - Because, a watchdog locally watches over the main process. Any solution that involves
         context kill, will require global context handling which can make the code messy even if executable.         
        */
        time.Sleep(10 * time.Second)

        fmt.Println("WOOF !")
        end = true // Kill the loop itself
        wg.Done()
    }

    go watchDog()

    const mainWait = 11 // <----------------------------------------------- can change
    x := 0
    for !end {
        wg.Add(1)

        go func() {
            defer func() {
                if mainWait > 9 {
                    // CAUTION: Even if watchDog finishes, this can still be called
                    fmt.Println("defer called")
                }
            }()
            // Main function blocking; wait for something to complete
            time.Sleep(mainWait * time.Second) // Some blocking call for 2 secs
            if !end {                          // Needed to prevent multiple .Done() incase watchdog ends
                //end = true
                wg.Done()
                return
            }
            x++
            fmt.Printf("%d) blocking call over\n", x)
        }()

        wg.Wait()
    }

    fmt.Println("~~~ End loop ~~~")
    time.Sleep(mainWait * 2 * time.Second) // See what happens to main go function waiting
    fmt.Println("~~~ End program ~~~")
}