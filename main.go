package main

import (
	"fmt"
)

const X string = "X"
const O string = "O"

func inicializarTabuleiro() [3][3]string {
	var tabuleiro [3][3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tabuleiro[i][j] = "|_|"
		}
	}
	return tabuleiro
}

func imprimirTabuleiro(tabuleiro [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(tabuleiro[i][j])
		}
		fmt.Println()
	}
}

func verificarCoordenada(tabuleiro [3][3]string, x, y int) bool {
	if tabuleiro[x][y] == "|_|" {
		return true
	}
	return false
}

func main() {
	tabuleiro := inicializarTabuleiro()
	jogadorAtual := X // Começa com o jogador X

	for {
		imprimirTabuleiro(tabuleiro)
		fmt.Printf("Jogador %s, insira as coordenadas (linha e coluna): ", jogadorAtual)

		var x, y int
		fmt.Scan(&x, &y)

		// Verifica se as coordenadas são válidas
		if x < 0 || x > 2 || y < 0 || y > 2 {
			fmt.Println("Coordenadas inválidas! Tente novamente.")
			continue
		}

		// Verifica se a célula está disponível
		if verificarCoordenada(tabuleiro, x, y) {
			// Marca a célula com o símbolo do jogador atual
			tabuleiro[x][y] = fmt.Sprintf("|%s|", jogadorAtual)

			// Alterna o jogador
			if jogadorAtual == X {
				jogadorAtual = O
			} else {
				jogadorAtual = X
			}
		} else {
			fmt.Println("Essa posição já está ocupada! Tente novamente.")
		}

		// Aqui você pode adicionar uma lógica para verificar se o jogo terminou
		// (ex.: vitória ou empate) e sair do loop.
	}
}
