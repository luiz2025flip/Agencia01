package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("AGENCIA - Orquestrador de Agentes")
		fmt.Println("===================================")
		fmt.Println("")
		fmt.Println("Uso: agencia <comando>")
		fmt.Println("")
		fmt.Println("Comandos:")
		fmt.Println("  tarefa <msg>  - Executa tarefa")
		fmt.Println("  api           - Pipeline API")
		fmt.Println("  pipeline      - Pipeline completo")
		os.Exit(1)
	}

	comando := args[1]

	switch comando {
	case "tarefa":
		if len(args) < 3 {
			fmt.Println("Uso: agencia tarefa <mensagem>")
			os.Exit(1)
		}
		fmt.Printf("Executando: %s\n", args[2])

	case "api":
		fmt.Println("🚀 Pipeline API...")

	case "pipeline":
		fmt.Println("🚀 Pipeline completo...")

	default:
		fmt.Printf("Comando desconhecido: %s\n", comando)
		os.Exit(1)
	}
}
