package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(50, 24)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("texto da função")

	// pegando os valores dos dois retornos ao mesmo tempo
	resultadoSoma, resultadoSubtracao := calculosMatematicos(20, 40)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// pegando somente o retorno da soma
	retornoSoma, _ := calculosMatematicos(20, 40)
	fmt.Println(retornoSoma)

	// pegando somente retorno da subtracao
	_, retornoSubtracao := calculosMatematicos(20, 40)
	fmt.Println(retornoSubtracao)

}
