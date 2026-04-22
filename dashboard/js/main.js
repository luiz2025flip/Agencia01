let currentData = null;

function init() {
    startTime();
    setInterval(updateStats, 1000);
}

async function updateStats() {
    try {
        const response = await fetch('data/status.json?t=' + new Date().getTime());
        const data = await response.json();
        currentData = data;
        
        renderSystemHealth(data.system);
        renderAgents(data.agents);
        renderMissions(data.missions, data);
        renderAuditLogs(data.missions, data);
        renderTaskBanner(data);
        
    } catch (e) {
        console.warn("Sincronizando Triple-ID Protocol...");
    }
}

function renderTaskBanner(data) {
    if (!data.active_task) return;
    document.getElementById('current-task-id').textContent = data.active_task.id;
    document.getElementById('current-task-name').textContent = data.active_task.name.toUpperCase();
}

function renderSystemHealth(system) {
    document.getElementById('cpu-bar').style.width = system.cpu + '%';
    document.getElementById('ram-bar').style.width = system.ram + '%';
    document.getElementById('cpu-val').textContent = system.cpu + '%';
    document.getElementById('ram-val').textContent = system.ram + '%';
    
    const torEl = document.getElementById('tor-status');
    torEl.textContent = system.tor ? 'ONLINE (TOR)' : 'OFFLINE';
    torEl.className = 'status-indicator ' + (system.tor ? 'status-online' : 'status-offline');
}

function renderAgents(agents) {
    const grid = document.getElementById('agents-grid');
    if (!agents) return;
    
    grid.innerHTML = agents.map(a => {
        let icon = "🤖";
        if (a.name.includes("arquiteto")) icon = "📐";
        if (a.name.includes("coder")) icon = "💻";
        if (a.name.includes("tdd")) icon = "🧪";
        if (a.name.includes("auditor")) icon = "🔍";
        if (a.name.includes("estrategista")) icon = "🧠";
        
        return `
            <div class="agent-card ${a.status}" onclick="inspectAgent('${a.name}')">
                <div class="icon">${icon}</div>
                <span class="name">${a.name.toUpperCase()}</span>
                <span class="role">${a.role}</span>
            </div>
        `;
    }).join('');
}

function renderMissions(missions, data) {
    const content = document.getElementById('mission-content');
    if (!missions || missions.length === 0) return;
    
    const activeMission = missions.find(m => m.status === "ACTIVE") || missions[missions.length - 1];
    document.querySelector('.mission-id').textContent = activeMission.id;
    
    content.innerHTML = `
        <div style="font-size: 0.7rem; color: var(--accent-magenta); margin-bottom: 5px;">PROJ: ${data.project_id}</div>
        <h4>${activeMission.id}: ${activeMission.title}</h4>
        <div class="steps">
            ${activeMission.steps.map(s => `
                <div class="step ${s.status}">
                    <span style="font-size: 0.6rem; opacity: 0.7;">[${s.id || '---'}]</span> 
                    ${s.status === 'done' ? '✓' : (s.status === 'doing' ? '⏳' : '○')} ${s.name}
                </div>
            `).join('')}
        </div>
    `;
}

function renderAuditLogs(missions, data) {
    const feed = document.getElementById('audit-feed');
    let allLogs = [];
    missions.forEach(m => {
        allLogs.push({ time: m.timestamp, tag: "PROJ", msg: `Switch: ${data.project_id}` });
        allLogs.push({ time: m.timestamp, tag: "MISS", msg: `Mission ${m.id} Activated` });
        (m.steps || []).forEach(s => {
            if (s.status === 'done') {
                allLogs.push({ time: "--:--:--", tag: "DONE", msg: `Task [${s.id}] Completed` });
            } else if(s.status === 'doing') {
                allLogs.push({ time: "--:--:--", tag: "TASK", msg: `Task [${s.id}] Priority` });
            }
        });
    });
    
    feed.innerHTML = allLogs.reverse().slice(0, 25).map(l => `
        <div class="log-item">
            <span class="time" style="font-size: 0.6rem;">[${l.time.includes('T') ? l.time.split('T')[1].split('.')[0] : l.time}]</span>
            <span class="tag" data-tag="${l.tag}">${l.tag}</span>: ${l.msg}
        </div>
    `).join('');
}

async function inspectAgent(name) {
    const modal = document.getElementById('agent-modal');
    const title = document.getElementById('modal-agent-name');
    const body = document.getElementById('modal-body');
    
    title.textContent = `AGENT INSPECTION: ${name.toUpperCase()}`;
    modal.classList.add('active');
    body.innerHTML = "<p>Buscando prompt de orquestração...</p>";
    
    try {
        const response = await fetch(`/inspect?agent=${name}`);
        if (!response.ok) throw new Error("Prompt not found");
        const content = await response.text();
        body.innerHTML = content || "Arquivo vazio.";
    } catch (e) {
        body.innerHTML = `<p style="color:red">Erro: ${e.message}</p>`;
    }
}

function closeModal() {
    document.getElementById('agent-modal').classList.remove('active');
}

function showToast(msg) {
    const container = document.getElementById('toast-container');
    const toast = document.createElement('div');
    toast.className = 'toast';
    toast.textContent = msg;
    container.appendChild(toast);
    setTimeout(() => toast.remove(), 3000);
}

async function sendCommand(cmd) {
    showToast(`Disparando ${cmd}...`);
    try {
        const response = await fetch('/', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ command: cmd, project_id: "AG_01_SQUAD", timestamp: new Date().toISOString() })
        });
        const result = await response.json();
        showToast(result.msg);
    } catch (e) {
        showToast("Conexão SVR falhou.");
    }
}

function startTime() {
    const timeEl = document.querySelector('.system-time');
    setInterval(() => {
        const now = new Date();
        timeEl.textContent = `SYSTEM TIME: ${now.toISOString().split('T')[1].split('Z')[0]}`;
    }, 50);
}

document.addEventListener('DOMContentLoaded', init);
window.onclick = function(event) {
    const modal = document.getElementById('agent-modal');
    if (event.target == modal) closeModal();
}
