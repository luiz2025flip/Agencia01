package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Page struct {
	Title  string
	Output string
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/run", runTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🌐 AGENCIA Web rodando em http://0.0.0.0:%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>AGENCIA - Orquestrador</title>
	<style>
		body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; background: #1a1a2e; color: #fff; }
		h1 { color: #00ff88; }
		.input-group { display: flex; gap: 10px; margin: 20px 0; }
		input { flex: 1; padding: 15px; font-size: 16px; border: none; border-radius: 8px; }
		button { padding: 15px 30px; background: #00ff88; color: #000; border: none; border-radius: 8px; cursor: pointer; font-weight: bold; }
		button:hover { background: #00cc6a; }
		.output { background: #16213e; padding: 20px; border-radius: 8px; white-space: pre-wrap; }
		.agents { display: flex; flex-wrap: wrap; gap: 10px; margin: 20px 0; }
		.agent { background: #0f3460; padding: 10px 15px; border-radius: 20px; font-size: 14px; }
	</style>
</head>
<body>
	<h1>🤖 AGENCIA</h1>
	<p>Orquestrador de Agentes AI</p>
	
	<div class="agents">
		<span class="agent">@arquiteto</span>
		<span class="agent">@engineer</span>
		<span class="agent">@tdd</span>
		<span class="agent">@auditor</span>
		<span class="agent">@documentador</span>
	</div>
	
	<form action="/run" method="POST">
		<div class="input-group">
			<input type="text" name="task" placeholder="@engineer crie função login..." autofocus>
			<button type="submit">Executar</button>
		</div>
	</form>
	
	{{if .Output}}
	<div class="output">{{.Output}}</div>
	{{end}}
</body>
</html>`

	t, err := template.New("page").Parse(tmpl)
	if err != nil {
		fmt.Fprint(w, "Erro:", err)
		return
	}

	p := Page{Title: "AGENCIA"}
	t.Execute(w, p)
}

func runTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	task := r.FormValue("task")

	if task == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Executa via opencode
	cmd := exec.Command("opencode", "run", task)
	cmd.Dir = "/home/kali/Área de trabalho/AGENCIA01"

	output, err := cmd.CombinedOutput()
	result := string(output)
	if err != nil {
		result = fmt.Sprintf("Erro: %v\n%s", err, result)
	}

	tmpl := `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>AGENCIA - Resultado</title>
	<style>
		body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; background: #1a1a2e; color: #fff; }
		h1 { color: #00ff88; }
		a { color: #00ff88; }
		.output { background: #16213e; padding: 20px; border-radius: 8px; white-space: pre-wrap; }
	</style>
</head>
<body>
	<h1>🤖 AGENCIA</h1>
	<p><a href="/">← Voltar</a></p>
	<div class="output">{{.Output}}</div>
</body>
</html>`

	t, _ := template.New("page").Parse(tmpl)
	p := Page{Title: "RESULTADO", Output: result}
	t.Execute(w, p)
}
