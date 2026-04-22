# 🧠 AGENCIA01 - Cérebro Obsidian

> Sistema de memória permanente do squad. Todos os agentes devem usar este sistema.

---

## Estrutura do Vault

```
/home/kali/Documents/Agencia_Vault/
├── 00_Central/      ← Índice e estado global
├── 10_Arquiteto/    ← Arquitetura e specs
├── 20_CoderHacker/  ← Código e scripts
├── 30_DevOps/      ← Infra e deploy
├── 40_FAQ/         ← Aprendizados e logs
└── 99_Memos/       ← Notas rápidas
```

---

## Como Cada Agente Usa

### 10_Arquiteto
- Salvar: arquitetura, ADRs, diagramas Mermaid, APIs, schemas
- Ler: specs do usuário, requisitos

### 20_CoderHacker
- Salvar: código (.py, .js, .ts), scripts, gists
- Ler: specs do Arquiteto, contratos TDD

### 30_DevOps
- Salvar: Dockerfiles, CI/CD, configs, deploy
- Ler: código do Coder

### 40_FAQ
- Salvar: logs de erros, soluções, aprendizados
- Ler: histórico de erros comuns

### 00_Central
- Salvar: índice, CHANGELOG, roadmap
- Ler: tudo

### 99_Memos
- Salvar: notas rápidas, feedback
- Ler: contexto recente

---

## Fluxo de Trabalho

1. Usuário define tarefa no Antigravity
2. Squad ativado com `/opensquad run agencia-squad`
3. Arquiteto → spec → 10_Arquiteto/
4. Coder → código → 20_CoderHacker/
5. DevOps → config → 30_DevOps/
6. FAQ → aprendido → 40_FAQ/
7. Central → Índice → 00_Central/

---

## Agentes e Pastas

| Agente | Pasta |
|--------|-------|
| arquiteto | 10_Arquiteto |
| engineer | 20_CoderHacker |
| tdd | 20_CoderHacker |
| prd_spec | 10_Arquiteto |
| validador_* | 40_FAQ |
| auditor_codigo | 40_FAQ |
| documentador | 10_Arquiteto |
| documentacao | 00_Central |
| mcp | Todas |
| faq | 40_FAQ |
| gestor | 00_Central |
| sentinel | 30_DevOps |
| monitor | 40_FAQ |
| estrategista | 10_Arquiteto |
| recrutador | 00_Central |
| mestre_prompt | 40_FAQ |

---

## 🛠️ Atualizações Recentes (22 Abr 2026)

### 1. Tor Configurado
- Tor rodando na porta 9050
- MCPs configurados com proxy Tor
- Teste de conexão: OK

### 2. MCPs com Tor
```json
{
  "brave-search": { "env": { "HTTPS_PROXY": "socks5://127.0.0.1:9050" }},
  "fetch": { "env": { "HTTPS_PROXY": "socks5://127.0.0.1:9050" }}
}
```

### 3. Testes TDD
- `test_boot_brain.py` criado
- 5/7 testes passando

### 4. Notificações Telegram
- `notify.sh` copiado para `/home/kali/Área de trabalho/AGENCIA01/scripts/`

### 5. ⚡ Regra de Economia de Recursos
Adicionada em todos os agentes:
- MCP: Prefira busca por texto
- FAQ: Priorize Obsidian → Docs → GitHub → Brave
- Engineer: Código > Imagens > Vídeos
- Arquiteto: Mermaid > ASCII > Imagens

---

## 📁 Arquivos do Projeto

```
AGENCIA01/
├── system-reminder.md     ← Este arquivo
├── opencode.md           ← Histórico completo
├── GEMINI.md             ← Regras do squad
├── gimini.md             ← Resumo para modelo
├── .mcp.json            ← Config MCP (Tor)
├── squads/
│   └── agencia-squad/
│       └── prompts/     ← 22 agentes
├── skills/              ← 3 skills
├── scripts/
│   └── notify.sh        ← Telegram
└── projetos/
    └── second-brain-obsidian/
        ├── boot_brain.py
        └── test_boot_brain.py
```

---

## Comandos Úteis

- `/opensquad run agencia-squad` - Ativar squad
- `/opensquad dashboard` - Ver dashboard 2D

---

*Atualizado: 2026-04-22*