package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal("Não foi possível importar o arquivo")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var count int

    for scanner.Scan() {
        count += len(scanner.Text())
    }

    fmt.Printf("Total: %d\n",count)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}