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
)

func main(){    
    cmd := exec.Command("C:\\installed\\Wireshark\\tshark.exe", "-i", "Wi-Fi", "-Pf", "host 192.168.1.90", "-w", "d:\\tmp\\wireshark-out.pcapng") // -i \"Wi-Fi\" -Pf \"host 192.168.1.90\" -w d:\\tmp\\wireshark-out.pcapng

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