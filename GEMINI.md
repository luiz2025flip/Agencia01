# Registro de Aprendizado: GEMINI (AntiGravity)

Este documento foi estabelecido como o cérebro central de contexto contínuo, onde ficam registrados os aprendizados em tempo real, decisões de fluxo do squad e anotações cruciais durante as sessões de desenvolvimento com o Usuário.

## Diretrizes e Decisões Fundamentais Tomadas
**As "Regras de Ouro" (Protocolo de Engenharia Soberana v7.9+):**

1. **Spec-First / Design-First (Rigidez Pré-Código):** 
   - **Regra:** Nenhuma linha de código deve ser escrita sem mapeamento completo (`.json`, `.yaml`, `.md`).
   - **Fluxo:** Estrategista -> Arquiteto -> TDD -> Validadores -> Auditor -> Código Real.

2. **Fluxo Linear e Síncrono (Soberania):**
   - O caminho crítico deve ser linear (Request -> LLM -> Response). Evitar threads paralelas que percam o contexto.

3. **Blindagem de Dados (OpSec):**
   - **PII Masking:** Proibido logar `chat_id`, `tokens` ou `senhas` in texto claro. Use máscaras (ex: `609*****32`).
   - **Fail-Fast:** Validação rigorosa de esquemas (Pydantic). Se o input for lixo, a execução para imediatamente.

4. **Resiliência e Economia (Efficient Intelligence):**
   - **Circuit Breaker:** Chamadas de IA com failover automático e timeouts estritos (5s).
   - **Semantic Cache:** Verificar hash no banco antes de gastar tokens via API.

5. **Paralelismo Inteligente:**
   - Validações operam assincronamente através de hooks, exceto Checkpoints aprovativos.


## Erros Detectados e Aprendizados (Log)

| Data / Sessão | Contexto do Erro | Solução Adotada (Aprendizado) |
| ------------- | ------------------ | ------------------------------- |
| `21 Abr 2026` | **Manifesto Corrompido:** No arquivo `squad.yml`, notou-se o agente `faq` instanciado e cadastrado duas vezes. Isso causaria duplicação de papéis ou laços em excesso na orquestração. | Removi ativamente a replicação, aprendendo a auditar as listas de perfis do time antes de iniciar grandes orquestrações. |
| 21 Abr 2026 | **Passividade do Dashboard:** O Agente Cliente reprovou a interface por ser apenas informativa. | Evoluir para "Cockpit Interativo". Interfaces não devem apenas mostrar dados, mas oferecer controle soberano imediato. |
| `21 Abr 2026` | **Transparência de Agentes:** Notou-se que o usuário se sente mais seguro quando pode ver o "pensamento" (prompt) do agente. | Implementado o **Deep Inspection Modal** no cockpit. |
| 21 Abr 2026 | **Protocolo Triple-ID:** Logs sem IDs de Projeto e Tarefa geram confusão em auditorias rápidas. | Criado o padrão `[Timestamp MS] [Proj ID] [Task ID]` para toda a telemetria. |
| 22 Abr 2026 | Missão #005: Deep Transparency | Concluída. /inspect rodando e Modais integrados. |
| 22 Abr 2026 | Ativação UI/UX Master | Agente criado e squad atualizado. |

## Roadmap de Aceleração Futura 
*(Identificado como defasagem para alcançarmos nossa independência 100%)*
- [x] Incorporação de agente para infraestrutura **SRE / DevOps** (Incorporado no squad.yml).
- [x] Criação de papel **UI / UX Master** (Design Semântico ativado).
- [x] Instanciação de um Especialista de **Banco de Dados (DBA)** (Incorporado no squad.yml).
- [ ] **Consumo TaskMaster:** Integrar lógica de DAG (Grafos) do projeto Desktop/TaskMaster para gerenciar squads não-lineares.
- [ ] **SecretVault v2:** Migrar gestão de segredos do Agencia_System para o AGENCIA01 usando criptografia AES-GCM.


---

## 🚀 Como Ativar o Squad

Para verificar o que aconteceu e ativar o squad, execute:

### Passo 1: Verificar Estado
1. Ler `opencode.md` para ver histórico
2. Ler `/home/kali/Documents/Agencia_Vault/40_FAQ/` para ver logs
3. Verificar se MCP está configurado em `.mcp.json`

### Passo 2: Ativar Squad
No chat do Antigravity, digitar:
```
/opensquad run agencia-squad
```

### Passo 3: Dashboard 2D (Opcional)
```
/opensquad dashboard
```

---

## 📋 Última Execução (21 Abr 2026)

### O que aconteceu:
1. **Arquiteto** criou `02-Arquitetura-e-MCP.md` com Mermaid
2. **TDD** criou `03-Contrato-TDD.md` com Gherkin
3. **Coder Hacker** criou `boot_brain.py` - script de boot
4. **Auditor** executou e validou o script
5. **Resultado:** Vault Obsidian configurado em `/home/kali/Documents/Agencia_Vault`

### Arquivos Criados:
- `projetos/second-brain-obsidian/boot_brain.py`
- `projetos/second-brain-obsidian/01-PRD-estrategia.md`
- `projetos/second-brain-obsidian/02-Arquitetura-e-MCP.md`
- `projetos/second-brain-obsidian/03-Contrato-TDD.md`

### MCP Configurado:
- `filesystem` - Projeto AGENCIA01
- `obsidian` - Vault em `/home/kali/Documents/Agencia_Vault`

---

> **Nota do AI (Gemini/AntiGravity):** Ler constantemente este arquivo antes do bootstrap de projetos novos dentro do open-squad para injetar o contexto e preservar a cultura "Spec-First".

---

## 🧠 Cérebro Obsidian

### Estrutura
```
Agencia_Vault/
├── 00_Central/     → Índice e estado global
├── 10_Arquiteto    → Arquitetura e specs
├── 20_CoderHacker  → Código e scripts
├── 30_DevOps       → Infra e deploy
├── 40_FAQ          → Aprendizados e logs
└── 99_Memos        → Notas rápidas
```

### Regra Principal
Cada agente DEVE salvar seu成果 na pasta correta e consultar as pastas dos outros agentes para entender o contexto.
