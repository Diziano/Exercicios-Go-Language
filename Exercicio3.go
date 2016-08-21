package main

import "fmt"
import "bufio"
import "os"

func main() {

	txt := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o texto: ")
	text, _ := txt.ReadString('\n')

	inverted := invert(text)

	fmt.Println("Seu texto invertido: "+inverted)
}

func invert(text string) string {

    txt := []rune(text)
    ret := []rune{}

    /* O -2 é porque começa em 0, e para não pegar a quebra de linha */
    for i := len(txt) - 2; i >= 0; i-- {
		ret = append(ret, txt[i])
    }

    return string(ret)
}
