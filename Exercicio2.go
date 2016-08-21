package main

import "fmt"
import "math/rand"

func main() {

	var vetor[10]int

	for i:=0; i < 10; i++ {
	
		vetor[i] = rand.Intn(100)		

	}

	for i:=0; i < len(vetor); i++ {
		fmt.Println(vetor[i])
	}

}