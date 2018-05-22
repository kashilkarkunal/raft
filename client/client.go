package main

import (
    "bufio"
    "fmt"
    "os"
)

func runCmd(cmd string) {
    fmt.Printf("cmd: %v\n", cmd)
}

func runPrompt() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Printf(">")
    for scanner.Scan() {
        cmd := scanner.Text()
        runCmd(cmd)
        fmt.Printf(">")
    }
}

func main() {
    runPrompt()
}
