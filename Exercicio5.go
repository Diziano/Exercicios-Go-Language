package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("access-log.txt")
    if err != nil {
        log.Fatal("Não foi possível importar o arquivo")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // map[chave]valor
    var ips = make(map[string]int)

    for scanner.Scan() {

        var split = strings.Split(scanner.Text(), " ")
        var ip = split[0] 

        ips[ip] = 1
    }

    fmt.Printf("Total de ips: %d\n",len(ips))

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}