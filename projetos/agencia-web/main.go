package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"
)

type Agent struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Task   string `json:"task"`
}

type Mission struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Agent  string `json:"agent"`
}

type ChatMessage struct {
	User  string `json:"user"`
	Text  string `json:"text"`
	Time  string `json:"time"`
	Agent string `json:"agent,omitempty"`
}

type ChatResponse struct {
	Response string        `json:"response"`
	Agent    string        `json:"agent,omitempty"`
	Chat     []ChatMessage `json:"chat,omitempty"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Rotas
	http.HandleFunc("/", home)
	http.HandleFunc("/chat", chat)
	http.HandleFunc("/api/status", apiStatus)
	http.HandleFunc("/api/agents", listAgents)
	http.HandleFunc("/api/missions", listMissions)

	fmt.Printf("🤖 AGENCIA RODANDO: http://0.0.0.0:%s\n", port)
	fmt.Println("========================================")
	fmt.Printf("🌐 Chat: http://0.0.0.0:%s/chat\n", port)
	fmt.Printf("📊 API:  http://0.0.0.0:%s/api/status\n", port)

	http.ListenAndServe(":"+port, nil)
}

var chatHistory []ChatMessage

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

	data := struct {
		Agents   []Agent       `json:"agents"`
		Missions []Mission     `json:"missions"`
		Chat     []ChatMessage `json:"chat"`
	}{Agents: agents, Missions: missions, Chat: chatHistory}

	tmpl := getDashboardTemplate()
	t, _ := template.New("home").Parse(tmpl)
	t.Execute(w, data)
}

func chat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := r.FormValue("message")

	if message == "" {
		// Retorna chat history
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ChatResponse{Chat: chatHistory})
		return
	}

	fmt.Printf("💬 Chat: %s\n", message)

	// Adiciona mensagem do usuário
	msg := ChatMessage{
		User: "você",
		Text: message,
		Time: time.Now().Format("15:04"),
	}
	chatHistory = append(chatHistory, msg)

	// Detecta agente
	agent := detectAgent(message)

	// Processa comando
	response := processMessage(message, agent)

	// Adiciona resposta
	reply := ChatMessage{
		User:  "@" + agent,
		Text:  response,
		Time:  time.Now().Format("15:04"),
		Agent: agent,
	}
	chatHistory = append(chatHistory, reply)

	// Limita Histórico (últimas 50 mensagens)
	if len(chatHistory) > 50 {
		chatHistory = chatHistory[len(chatHistory)-50:]
	}

	// Retorna JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ChatResponse{
		Response: response,
		Agent:    agent,
		Chat:     chatHistory,
	})
}

func detectAgent(message string) string {
	lower := strings.ToLower(message)

	if strings.Contains(lower, "arquit") || strings.Contains(lower, "arquitetura") || strings.Contains(lower, "estrutura") {
		return "arquiteto"
	}
	if strings.Contains(lower, "engineer") || strings.Contains(lower, "cod") || strings.Contains(lower, "implement") || strings.Contains(lower, "crie") || strings.Contains(lower, "função") {
		return "engineer"
	}
	if strings.Contains(lower, "tdd") || strings.Contains(lower, "teste") || strings.Contains(lower, "spec") {
		return "tdd"
	}
	if strings.Contains(lower, "audit") || strings.Contains(lower, "revis") || strings.Contains(lower, "review") {
		return "auditor"
	}
	if strings.Contains(lower, "doc") || strings.Contains(lower, "document") || strings.Contains(lower, "readme") {
		return "documentador"
	}
	if strings.Contains(lower, "seguran") || strings.Contains(lower, "security") || strings.Contains(lower, "vuln") {
		return "sentinel"
	}
	if strings.Contains(lower, "monitor") || strings.Contains(lower, "status") || strings.Contains(lower, "métricas") {
		return "monitor"
	}
	if strings.Contains(lower, "gestao") || strings.Contains(lower, "gerenciar") || strings.Contains(lower, "projeto") {
		return "gestor"
	}
	if strings.Contains(lower, "estrateg") || strings.Contains(lower, "planej") || strings.Contains(lower, "roadmap") {
		return "estrategista"
	}
	if strings.Contains(lower, "prd") || strings.Contains(lower, "requisito") || strings.Contains(lower, "especific") {
		return "prd_spec"
	}
	if strings.Contains(lower, "recrut") || strings.Contains(lower, "orquest") || strings.Contains(lower, "coordena") {
		return "recrutador"
	}
	return "recrutador" // Default
}

func processMessage(msg, agent string) string {
	lower := strings.ToLower(msg)

	// Respostas por agente
	switch agent {
	case "arquiteto":
		return "📐 Analisando arquitetura...\n\nVou propor a melhor estrutura para seu projeto.\n\nQual o tipo de aplicação? (web, API, mobile, etc)"
	case "engineer":
		if strings.Contains(lower, "criar") || strings.Contains(lower, "função") {
			return "💻 Vou criar o código.\n\nQual a linguagem? (Go, Python, JS, etc)"
		}
		return "💻 Entendido! Vou implementar.\n\nPode dar mais detalhes?"
	case "tdd":
		return "🧪 criando testes...\n\nQuais funções/módulos deseja testar?"
	case "auditor":
		return "🔍 Revisando código...\n\nQuer que eu analise algum arquivo específico?"
	case "documentador":
		return "📝 Vou documentar.\n\nQual parte do projeto?"
	case "sentinel":
		return "🔐 Análise de segurança...\n\nVerificando vulnerabilidades..."
	case "monitor":
		return "📊 Sistema Monitorado:\n\n- Agentes: 19\n- Missões: 3\n- Uptime: 100%\n- Status: 🟢 OK"
	case "gestor":
		return "📋 Gestão de Projeto:\n\nMissões ativas:\n1. Missão #001 ✅\n2. Missão #002 🔄\n3. Missão #003 ⏳"
	case "estrategista":
		return "🎯 Planejamento:\n\nQual o objetivo estratégico?"
	case "prd_spec":
		return "📋 Especificação PRD:\n\nDescreva o requisito em detalhes."
	case "recrutador":
		return "🎯 Orquestrando agentes para: " + msg + "\n\nVou coordenar os agentes necessários."
	default:
		return "🤖 Entendido! " + msg + "\n\nComo posso ajudar?"
	}
}

func apiStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "ok",
		"agents":   19,
		"missions": 3,
		"uptime":   "100%",
		"chat":     len(chatHistory),
	})
}

func listAgents(w http.ResponseWriter, r *http.Request) {
	agents := []Agent{
		{Name: "arquiteto", Status: "ready"},
		{Name: "engineer", Status: "ready"},
		{Name: "tdd", Status: "ready"},
		{Name: "auditor", Status: "ready"},
		{Name: "recrutador", Status: "ready"},
		{Name: "documentador", Status: "ready"},
		{Name: "sentinel", Status: "ready"},
		{Name: "monitor", Status: "ready"},
		{Name: "gestor", Status: "ready"},
		{Name: "estrategista", Status: "ready"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(agents)
}

func listMissions(w http.ResponseWriter, r *http.Request) {
	missions := []Mission{
		{Name: "Missão #001", Status: "completed"},
		{Name: "Missão #002", Status: "running"},
		{Name: "Missão #003", Status: "pending"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(missions)
}

// Dashboard Template
func getDashboardTemplate() string {
	return `<!DOCTYPE html>
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
.grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 16px; }
.card { background: linear-gradient(135deg, #16213e, #1f2f4e); border-radius: 12px; padding: 20px; border: 1px solid #333; }
.card:hover { border-color: #00ff88; transform: translateY(-4px); transition: all 0.3s; }
.stat-dot { display: inline-block; width: 10px; height: 10px; border-radius: 50%; margin-right: 8px; }
.ready { background: #00ff88; }
.running { background: #ffd93d; }
.health-grid { display: grid; grid-template-columns: repeat(4, 1fr)); gap: 12px; }
.health-item { background: #0f0f23; padding: 16px; border-radius: 8px; text-align: center; }
.health-item .value { font-size: 1.5rem; font-weight: bold; color: #00ff88; }
.health-item .label { font-size: 0.8rem; color: #888; }
.section { margin-bottom: 24px; }
.section h2 { color: #00ff88; margin-bottom: 16px; }

/* Chat Styles */
.chat-container { background: #0f0f23; border-radius: 12px; padding: 20px; margin-top: 20px; }
.chat-messages { height: 400px; overflow-y: auto; margin-bottom: 16px; }
.chat-message { margin-bottom: 12px; padding: 12px; border-radius: 8px; }
.chat-message.user { background: #1a3a5a; margin-left: 20%; }
.chat-message.agent { background: #0f3460; margin-right: 20%; }
.chat-message .time { font-size: 0.7rem; color: #888; }
.chat-message .agent-name { color: #00ff88; font-weight: bold; }
.chat-input { display: flex; gap: 12px; }
.chat-input input { flex: 1; padding: 16px; font-size: 16px; border: 2px solid #333; border-radius: 8px; background: #16213e; color: #fff; }
.chat-input input:focus { outline: none; border-color: #00ff88; }
.chat-input button { padding: 16px 32px; font-size: 16px; font-weight: bold; background: #00ff88; border: none; border-radius: 8px; cursor: pointer; }
.chat-input button:hover { transform: scale(1.05); }

/* Tab Styles */
.tabs { display: flex; gap: 8px; margin-bottom: 16px; }
.tab { padding: 12px 24px; background: #16213e; border: none; border-radius: 8px; cursor: pointer; color: #fff; }
.tab.active { background: #00ff88; color: #000; }
.tab-content { display: none; }
.tab-content.active { display: block; }
</style>
</head>
<body>
<div class="header">
<h1>🤖 AGENCIA SQUAD</h1>
<p>Orquestrador de Agentes AI com Chat</p>
</div>
<div class="container">
<div class="tabs">
<button class="tab active" onclick="showTab('dashboard')">📊 Dashboard</button>
<button class="tab" onclick="showTab('chat')">💬 Chat</button>
<button class="tab" onclick="showTab('agents')">🤖 Agentes</button>
</div>

<!-- Dashboard Tab -->
<div id="dashboard" class="tab-content active">
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
<h2>📋 Missões</h2>
<div class="grid">
<div class="card"><strong>Missão #001</strong><p>Recrutador</p><span class="stat-dot ready"></span>completed</div>
<div class="card"><strong>Missão #002</strong><p>Engineer</p><span class="stat-dot running"></span>running</div>
<div class="card"><strong>Missão #003</strong><p>TDD</p><span class="stat-dot ready"></span>pending</div>
</div>
</div>
</div>

<!-- Chat Tab -->
<div id="chat" class="tab-content">
<div class="chat-container">
<div class="chat-messages" id="chatMessages">
{{range .Chat}}
<div class="chat-message {{if eq .User \"você\"}}user{{else}}agent{{end}}">
{{if ne .User "você"}}<span class="agent-name">{{.User}}</span><br>{{end}}
{{.Text}}
<span class="time">{{.Time}}</span>
</div>
{{end}}
</div>
<form class="chat-input" onsubmit="sendMessage(event)">
<input type="text" id="messageInput" placeholder="Digite sua mensagem..." autocomplete="off">
<button type="submit">🚀 Enviar</button>
</form>
</div>
</div>

<!-- Agents Tab -->
<div id="agents" class="tab-content">
<div class="grid">
{{range .Agents}}
<div class="card">
<h3>@{{.Name}}</h3>
<p style="color:#888;">{{.Task}}</p>
<span class="stat-dot {{.Status}}"></span>{{.Status}}
</div>
{{end}}
</div>
</div>
</div>

<script>
function showTab(tabId) {
document.querySelectorAll('.tab-content').forEach(t => t.classList.remove('active'));
document.querySelectorAll('.tab').forEach(t => t.classList.remove('active'));
document.getElementById(tabId).classList.add('active');
event.target.classList.add('active');
}

function sendMessage(e) {
e.preventDefault();
const input = document.getElementById('messageInput');
const message = input.value;
if (!message) return;

fetch('/chat?message=' + encodeURIComponent(message))
.then(r => r.json())
.then(data => {
const chatDiv = document.getElementById('chatMessages');
chatDiv.innerHTML = '';
data.chat.forEach(msg => {
const div = document.createElement('div');
div.className = msg.user === 'você' ? 'chat-message user' : 'chat-message agent';
div.innerHTML = (msg.user !== 'você' ? '<span class="agent-name">' + msg.user + '</span><br>' : '') + msg.text + '<span class="time">' + msg.time + '</span>';
chatDiv.appendChild(div);
});
chatDiv.scrollTop = chatDiv.scrollHeight;
});
input.value = '';
}

// Auto-refresh chat
setInterval(() => {
fetch('/chat')
.then(r => r.json())
.then(data => {
if (data.length > 0) {
const chatDiv = document.getElementById('chatMessages');
const html = data.map(msg => 
'<div class="' + (msg.user === 'você' ? 'chat-message user' : 'chat-message agent') + '">' +
(msg.user !== 'você' ? '<span class="agent-name">' + msg.user + '</span><br>' : '') + msg.text + '<span class="time">' + msg.time + '</span></div>'
).join('');
document.getElementById('chatMessages').innerHTML = html;
}
});
}, 3000);
</script>
</body>
</html>`
}
