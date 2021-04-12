/*

Build command: go build -o wire-shark-runner.exe .\RunCommand.go
Raw cmd to exec in our example: C:\installed\Wireshark\tshark.exe -i "Wi-Fi" -Pf "host 192.168.1.90" -w d:\tmp\wireshark-out.pcapng

@author Arjun Dhar
@see https://golang.org/pkg/os/exec/#Command
@see ...\Scripts\GO
*/

package main

import(
    "fmt"
    "os/exec"
    "bytes"

    "math/rand"
    "time"
    "strconv"

    "flag"
)

func random() int {
    min := 10
    max := 99999
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min + 1) + min
}

func main(){
    id := strconv.Itoa(random())
    var ip = flag.String("I", "192.168.1.109", "IP of printer")
    //cmd := exec.Command("C:\\installed\\Wireshark\\tshark.exe", "-i", "Wi-Fi", "-Pf", "host 192.168.1.90", "-w", "d:\\tmp\\wireshark-out-"+id+".pcapng") // -i \"Wi-Fi\" -Pf \"host 192.168.1.90\" -w d:\\tmp\\wireshark-out.pcapng
    cmd := exec.Command("C:\\Program Files\\Wireshark\\tshark.exe", "-i", "Ethernet", "-Pf", "host " + *ip, "-w", "C:\\Users\\25007084\\Desktop\\arjun\\ws-logs\\domino-tto-test-"+id+".pcapng") // -i \"Wi-Fi\" -Pf \"host 192.168.1.90\" -w d:\\tmp\\wireshark-out.pcapng

    var out bytes.Buffer
    cmd.Stdout = &out

    if err := cmd.Start(); err != nil { 
        fmt.Println("Error: ", err)
    }

    go func() {
        for {
            //bs := out.Bytes()
            bs := out.String()
            if len(bs) > 0 {
                //fmt.Printf("%s", string(bs))
                fmt.Printf("%s", bs)
            }
        }
    }()

    cmd.Wait()
}