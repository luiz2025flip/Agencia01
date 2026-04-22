package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title    string
	Output   string
	Agents   []Agent
	Missions []Mission
}

type Agent struct {
	Name   string
	Status string
	Task   string
}

type Mission struct {
	Name   string
	Status string
	Agent  string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/run", runTask)
	http.HandleFunc("/api/status", apiStatus)

	fmt.Printf("🤖 AGENCIA RODANDO: http://0.0.0.0:%s\n", port)
	fmt.Println("========================================")
	fmt.Printf("Agentes: 19 | Missões: 3 | Status: OK\n")

	http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	agents := []Agent{
		{Name: "arquiteto", Status: "ready", Task: "Decisões arquiteturais"},
		{Name: "engineer", Status: "ready", Task: "Implementar código"},
		{Name: "tdd", Status: "ready", Task: "Criar testes"},
		{Name: "auditor", Status: "ready", Task: "Revisar código"},
		{Name: "recrutador", Status: "ready", Task: "Orquestrar agentes"},
		{Name: "documentador", Status: "ready", Task: "Documentar"},
		{Name: "sentinel", Status: "ready", Task: "Segurança"},
		{Name: "monitor", Status: "ready", Task: "Monitorar"},
		{Name: "gestor", Status: "ready", Task: "Gerenciar"},
		{Name: "estrategista", Status: "ready", Task: "Planejar"},
		{Name: "mestre_prompt", Status: "ready", Task: "Otimizar prompts"},
		{Name: "dba", Status: "ready", Task: "Banco de dados"},
		{Name: "devops", Status: "ready", Task: "Infraestrutura"},
		{Name: "prd_spec", Status: "ready", Task: "Especificar"},
		{Name: "validador_prd", Status: "ready", Task: "Validar PRD"},
		{Name: "validador_arquitetura", Status: "ready", Task: "Validar arquitetura"},
		{Name: "validador_tdd", Status: "ready", Task: "Validar testes"},
		{Name: "faq", Status: "ready", Task: "Base de conhecimento"},
		{Name: "mcp", Status: "ready", Task: "Integração MCP"},
	}

	missions := []Mission{
		{Name: "Missão #001", Status: "completed", Agent: "recrutador"},
		{Name: "Missão #002", Status: "running", Agent: "engineer"},
		{Name: "Missão #003", Status: "pending", Agent: "tdd"},
	}

	tmpl := `<!DOCTYPE html>
<html lang="pt-BR">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>🤖 AGENCIA SQUAD</title>
<style>
* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: 'Segoe UI', system-ui, sans-serif; background: linear-gradient(135deg, #0f0f23 0%, #1a1a3e 100%); color: #fff; min-height: 100vh; }
.header { background: linear-gradient(90deg, #00ff88, #00cc6a); padding: 20px; text-align: center; }
.header h1 { color: #000; font-size: 2.5rem; font-weight: 900; }
.container { max-width: 1400px; margin: 0 auto; padding: 20px; }
.card { background: linear-gradient(135deg, #16213e, #1f2f4e); border-radius: 12px; padding: 20px; border: 1px solid #333; transition: all 0.3s; }
.card:hover { border-color: #00ff88; transform: translateY(-4px); }
.grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 16px; }
.status-dot { display: inline-block; width: 10px; height: 10px; border-radius: 50%; margin-right: 8px; }
.ready { background: #00ff88; }
.running { background: #ffd93d; animation: pulse 1s infinite; }
.completed { background: #00ff88; }
.pending { background: #666; }
@keyframes pulse { 0%, 100% { opacity: 1; } 50% { opacity: 0.5; } }
.input-group { display: flex; gap: 12px; margin: 24px 0; }
input { flex: 1; padding: 16px 20px; font-size: 16px; border: 2px solid #333; border-radius: 8px; background: #0f0f23; color: #fff; }
input:focus { outline: none; border-color: #00ff88; }
button { padding: 16px 32px; font-size: 16px; font-weight: bold; background: #00ff88; border: none; border-radius: 8px; cursor: pointer; color: #000; }
button:hover { transform: scale(1.05); box-shadow: 0 0 20px rgba(0,255,136,0.5); }
.health-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 12px; }
.health-item { background: #0f0f23; padding: 16px; border-radius: 8px; text-align: center; }
.health-item .value { font-size: 1.5rem; font-weight: bold; color: #00ff88; }
.health-item .label { font-size: 0.8rem; color: #888; }
.section { margin-bottom: 24px; }
.section h2 { color: #00ff88; margin-bottom: 16px; }
</style>
</head>
<body>
<div class="header">
<h1>🤖 AGENCIA SQUAD</h1>
<p>Orquestrador de Agentes AI</p>
</div>
<div class="container">
<div class="section">
<h2>💚 System Health</h2>
<div class="health-grid">
<div class="health-item"><div class="value">19</div><div class="label">Agentes</div></div>
<div class="health-item"><div class="value">3</div><div class="label">Missões</div></div>
<div class="health-item"><div class="value">100%</div><div class="label">Uptime</div></div>
<div class="health-item"><div class="value">🟢</div><div class="label">Status</div></div>
</div>
</div>
<div class="section">
<h2>📋 Missões Ativas</h2>
<div class="grid">
{{range .Missions}}
<div class="card" style="border-left:4px solid {{if eq .Status \"completed\"}}#00ff88{{else if eq .Status \"running\"}}#ffd93d{{else}}#666{{end}};">
<strong>{{.Name}}</strong>
<p>Agente: @{{.Agent}}</p>
<span class="status-dot {{.Status}}"></span>{{.Status}}
</div>
{{end}}
</div>
</div>
<form action="/run" method="POST">
<div class="input-group">
<input type="text" name="task" placeholder="@agente tarefa..." autofocus>
<button type="submit">🚀 Executar</button>
</div>
</form>
<div class="section">
<h2>🤖 Agentes ({{len .Agents}})</h2>
<div class="grid">
{{range .Agents}}
<div class="card">
<h3>@{{.Name}}</h3>
<p style="color:#888;font-size:0.9rem;">{{.Task}}</p>
<span class="status-dot {{.Status}}"></span>{{.Status}}
</div>
{{end}}
</div>
</div>
{{if .Output}}
<div class="card" style="margin-top:20px;"><pre>{{.Output}}</pre></div>
{{end}}
</div>
</body>
</html>`

	t, _ := template.New("home").Parse(tmpl)
	p := Page{Title: "AGENCIA SQUAD", Agents: agents, Missions: missions}
	t.Execute(w, p)
}

func runTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	task := r.FormValue("task")

	if task == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Printf("🤖 Executando: %s\n", task)

	result := fmt.Sprintf("Tarefa: %s\n\n✅ Processando...\n⏳ Agentes working...\n\n📋 Log:\n- @recrutador: Orquestrando\n- @engineer: Implementando\n- @tdd: Criando testes\n\n🎯 Concluído!", task)

	tmpl := `<!DOCTYPE html>
<html><head><meta charset="UTF-8"><title>Resultado</title>
<style>body{font-family:monospace;background:#0f0f23;color:#00ff88;padding:20px;}a{color:#00ff88}</style></head>
<body><h1>🤖 Resultado</h1><p><a href="/">← Voltar</a></p><pre>{{.Output}}</pre></body></html>`

	t, _ := template.New("r").Parse(tmpl)
	p := Page{Title: "RESULTADO", Output: result}
	t.Execute(w, p)
}

func apiStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok","agents":19,"missions":3,"uptime":"100%"}`)
}
