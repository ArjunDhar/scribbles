# execute commands via Go Lang on Windows

```
@author Arjun Dhar
@see https://golang.org/pkg/os/exec/#Command
```

A simple `Go` utility to execute a command. In this Exmaple a windows command.<br />

The **example** executes a [tashark](https://www.wireshark.org/docs/man-pages/tshark.html) utility installed to sniff packets over an IP.<br />

It assumes `tshark` is installed @ `C:\installed\Wireshark\`. **Change** path and flags based on the utility you want to use.

The utility clarifies running non CMD commands and also how to pipe the output **while running** without waiting to finish.

## Build
Build command: `go build -o my-runner.exe .\RunCommand.go`

## Run
Simply...

`myrunner.exe`

