# PASSO A PASSO: RODAR AGENCIA NO REPLIT

## 1. Acesse
Vá em: https://replit.com

## 2. Novo Repl
Clique **+ New Repl** (canto superior direito)

## 3. Escolha Go
Selecione **Go** como linguagem

## 4. Nome
Digite: `agencia-go`

## 5. Copie Este Código
Substitua TUDO que estiver em `main.go` por:

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printMenu()
		return
	}

	cmd := args[0]
	
	switch cmd {
	case "tarefa":
		if len(args) < 2 {
			fmt.Println("Uso: agencia tarefa <mensagem>")
			return
		}
		msg := strings.Join(args[1:], " ")
		fmt.Printf("🤖 Executando: %s\n", msg)
		// Executa via OpenCode local
		exec.Command("opencode", "run", msg).Run()
		
	case "api":
		fmt.Println("🚀 Pipeline API...")
		runPipeline("api")
		
	case "build":
		fmt.Println("🔨 Build e deploy...")
		
	default:
		fmt.Printf("Comando '%s' não reconhecido\n", cmd)
		printMenu()
	}
}

func printMenu() {
	fmt.Println("AGENCIA - Orquestrador de Agentes")
	fmt.Println("================================")
	fmt.Println("")
	fmt.Println("Uso: agencia <comando>")
	fmt.Println("")
	fmt.Println("Comandos:")
	fmt.Println("  tarefa <msg>  - Executa tarefa")
	fmt.Println("  api           - Pipeline API")
	fmt.Println("  build        - Build & Deploy")
}

func runPipeline(pipeline string) {
	fmt.Println("Executando pipeline:", pipeline)
}
```

## 6. Run
Clique no botão **Run** (verde) ou pressione `Ctrl+Enter`

## 7. Teste
No console, digite:
```
tarefa criar hello world
```

---

**Feito!** 🎉