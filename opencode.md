# OpenCode - Histórico e Aprendizados

> Este documento registra a configuração, erros, aprendizados e decisões do projeto AGENCIA01.

---

## 📋 Histórico de Conversa

### Sessão Atual
- **Código de Sessão:** `opencode -s ses_261b63e01ffesVnHzXppyYCLz4`
- **Como continuar:** `opencode -s ses_261b63e01ffesVnHzXppyYCLz4`

### Contexto Inicial
- **Usuário:** Luiz (Kali Linux)
- **Objetivo:** Integrar agentes do Multica ao Antigravity via OpenSquad
- **Stack:** OpenCode + big-pickle (modelo gratuito)

### O Que Foi Feito Nesta Sessão

| Data | Ação | Resultado |
|------|------|----------|
| 2024-04-21 | Setup OpenSquad no AGENCIA01 | Strukturacriada |
| 2024-04-21 | Instaladas 3 skills | apify, image-ai-generator, resend |
| 2024-04-21 | Criados 19 agentes | Prompts profissionais |
| 2024-04-21 | Criado opencode.md | Documento de histórico |
| 2024-04-21 | Knowledge Harvest (Desktop) | Absorvidos protocolos de Agency_System e TaskMaster |
| 2026-04-21 | Missão #002: Monitor Real-time | Dashboard integrado via Python monitor |


### Decisões Tomadas

| Data | Decisão | Motivo |
|------|---------|--------|
| 2024-04-21 | Usar OpenSquad | Suporta nativamente Antigravity e OpenCode |
| 2024-04-21 | Usar big-pickle | Modelo gratuito, não precisa API key |
| 2024-04-21 | Criar prompts profissionais | Prompts iniciais eram básicos |
| 2024-04-21 | Criar 3 novos agentes | documentacao, mcp, faq |
| 2024-04-21 | Documentar sessão | Para continuar de onde paramos |

---

## 🛠 Configuração do Ambiente

### OpenCode Config
```json
{
  "$schema": "https://opencode.ai/config.json",
  "model": "big-pickle"
}
```

### Estrutura do Projeto
```
AGENCIA01/
├── squads/
│   └── agencia-squad/
│       ├── squad.yml
│       └── prompts/ (19 agentes)
├── skills/ (3 skills instaladas)
├── _opensquad/
└── .env
```

### Skills Instaladas
- `apify` - Web scraping e automação
- `image-ai-generator` - Geração de imagens via OpenRouter
- `resend` - Envio de emails via Resend

### Obsidian
- **Vault:** `/home/kali/Documents/Agencia_Vault`
- **Estrutura:**
  - `00_Central/` - Índice central
  - `10_Arquiteto/` - Agente Arquiteto
  - `20_CoderHacker/` - Agente Coder
  - `30_DevOps/` - Agente DevOps
  - `40_FAQ/` - Base de conhecimento
  - `99_Memos_Compartilhados/` - Memos

### MCP Configuração (.mcp.json)
```json
{
  "mcpServers": {
    "filesystem": {
      "command": "npx",
      "args": ["-y", "@modelcontextprotocol/server-filesystem", "/home/kali/Área de trabalho/AGENCIA01"]
    },
    "obsidian": {
      "command": "npx", 
      "args": ["-y", "@modelcontextprotocol/server-filesystem", "/home/kali/Documents/Agencia_Vault"]
    }
  }
}
```

---

## 📝 Agentes Criados (19+ total)

### Agentes de Desenvolvimento
| Agente | Arquivo | Descrição |
|--------|---------|-----------|
| arquiteto | arquiteto.md | Arquiteto de Software Sênior |
| engineer | engineer.md | Engenheiro de Software |
| tdd | tdd.md | Desenvolvedor TDD |
| prd_spec | prd_spec.md | Especificador de PRD |

### Agentes de Validação
| Agente | Arquivo | Descrição |
|--------|---------|-----------|
| validador_prd | validador_prd.md | Valida especificações |
| validador_arquitetura | validador_arquitetura.md | Valida arquitetura |
| validador_tdd | validador_tdd.md | Valida testes |
| auditor_codigo | auditor_codigo.md | Auditor de Código |

### Agentes de Suporte
| Agente | Arquivo | Descrição |
|--------|---------|-----------|
| documentador | documentador.md | Documentação técnica |
| documentacao | documentacao.md | Documentação global |
| mcp | mcp.md | Integração MCP |
| faq | faq.md | FAQ com aprendizado |
| mestre_prompt | mestre_prompt.md | Engenheiro de prompts |

### Agentes de Gestão
| Agente | Arquivo | Descrição |
|--------|---------|-----------|
| estrategista | estrategista.md | Planejamento estratégico |
| recrutador | recrutador.md | Orquestrador de agentes |
| gestor | gestor.md | Gestão de projeto |
| sentinel | sentinel.md | Segurança |
| monitor | monitor.md | Qualidade |

---

## ❌ Erros Cometidos

### 1. Prompts Iniciais Básicos Demais
**Problema:** Criei prompts genéricos sem profundidade técnica.

**Solução:** Re-escrevi todos os 16 agentes com prompts profissionais detalhados (~5KB cada).

**Aprendizado:** Prompts devem seguir estrutura padrão:
- Identidade clara
- Checklist de tarefas
- Templates reutilizáveis
- Exemplos de código

### 2. OpenSquad init Interativo Falhou
**Problema:** O comando `npx opensquad init` precisa de input interativo que foi fechado.

**Solução:** Criei estrutura manualmente + instalei skills uma a uma.

**Aprendizado:** Para ambientes não-interativos:
- Criar arquivos manualmente
- Usar `npx opensquad install <skill>` para cada skill

### 3. Multica vs OpenSquad Confusão
**Problema:** Achei que o Multica poderia ser servidor MCP.

**Realidade:**
- Multica: agentes consomem MCP servers
- OpenSquad: orchestra agentes na IDE
- Diferentes paradigmas

**Aprendizado:** Entender a arquitetura antes de propor soluções.

---

## ✅ Aprendizados

### OpenCode
1. **Modelos gratuitos:** big-pickle, gpt-5-nano funcionam sem API key
2. **Configuração:** Via `~/.config/opencode/opencode.json`
3. **Integração:** Já vem com várias ferramentas nativas

### OpenSquad
1. **Não funciona via CLI:** Precisa ser rodado dentro da IDE
2. **Comandos:** `/opensquad`, `/opensquad run <squad>`
3. **Skills:** Instalar com `npx opensquad install <nome>`
4. **Squad definition:** Arquivo YAML com agents e checkpoints

### Antigravity
1. **MCP config:** `~/.gemini/antigravity/mcp_config.json`
2. **Integração OpenSquad:** Suportado nativamente
3. **Modelos:** Suporta Gemini, Claude, e outros via proxy

### Telegram Bot
- **Token:** 8755791756:AAGGvfIXkm1HpeZIsG-2goVq3PLTBaV91h8
- **Chat ID:** 6092306332
- **Script:** `/home/kali/multica/notify.sh`

---

## 🔧 Com Úteis

### Iniciar Multica
```bash
cd /home/kali/multica && ./start.sh
```

### Testar OpenSquad no Antigravity
```bash
# No chat do Antigravity:
/opensquad run agencia-squad
```

### Listar Skills
```bash
cd /home/kali/Área\ de\ trabalho/AGENCIA01
npx opensquad skills
```

### Instalar Nova Skill
```bash
cd /home/kali/Área\ de\ trabalho/AGENCIA01
npx opensquad install <nome-da-skill>
```

---

## 📌 Pendências

- [ ] Testar orquestração no Antigravity
- [ ] Configurar MCPs adicionais no mcp.json
- [ ] Criar base inicial de FAQ
- [ ] Testar integração com Telegram

---

## 📚 Referências

- [OpenCode Docs](https://opencode.ai/docs/)
- [OpenSquad GitHub](https://github.com/renatoasse/opensquad)
- [MCP Servers](https://modelcontextprotocol.io/servers)
- [Antigravity Lab](https://antigravitylab.net)

---

## 🚀 Como Continuar Sessão

Para continuar de onde paramos na próxima vez:

```bash
# Abrir o projeto
cd /home/kali/Área\ de\ trabalho/AGENCIA01

# Continuar sessão específica
opencode -s ses_261b63e01ffesVnHzXppyYCLz4

# Ou simplesmente abrir o projeto
opencode .
```

---

## 📝 O Que Foi Adicionado Nesta Sessão

### Novos Agentes (3)
1. **documentacao.md** - Documentação global com versionamento e rastreabilidade
2. **mcp.md** - Integração com MCPs top de mercado
3. **faq.md** - FAQ com aprendizado contínuo de erros

### Updates Realizados
- Atualizado squad.yml com 19 agentes
- Criado opencode.md para histórico

---

## 🔄 Scripts de Automação

### boot_brain.py
**Local:** `projetos/second-brain-obsidian/boot_brain.py`

**Função:** Valida e cria estrutura do Obsidian Vault automaticamente
- Cria pastas se não existirem
- Gera índice mestre
- Respeita TDD (não faz overwrite)

**Uso:**
```bash
python3 projetos/second-brain-obsidian/boot_brain.py
```

---

## 📚 Arquivos de Conhecimento

### 01_Log_Boot_Zero.md
**Local:** `/home/kali/Documents/Agencia_Vault/40_FAQ/01_Log_Boot_Zero.md`

**Conteúdo:** Log de erro/otimização do boot
- Contexto anterior: gargalo e repetição manual
- Solução: boot_brain.py para automatizar estrutura
- Aprendizado: sempre rodar "Módulo de Boot" invisível

---

## 📌 Pendências Atualizadas

- [x] Testar orquestração no Antigravity - FUNCIONOU!
- [x] Configurar MCPs - Obsidian + Filesystem
- [x] Criar base inicial de FAQ
- [x] Squad executou com sucesso

---

## ✅ Execução Bem-Sucedida

### O que aconteceu:
1. Arquiteto criou 02-Arquitetura-e-MCP.md
2. TDD criou 03-Contrato-TDD.md
3. Coder Hacker criou boot_brain.py
4. Auditor executou e validou
5. Obsidian instalado e vault configurado

### Resultado:
- Vault em `/home/kali/Documents/Agencia_Vault`
- MCP configurado para ler Obsidian
- Scripts de automação funcionando

---

## 🚀 Como Ativar o Squad

### Comandos:
1. **Rodar o squad:**
```
/opensquad run agencia-squad
```

2. **Ver dashboard 2D:**
```
/opensquad dashboard
```

3. **Verificar estado:**
```
/opensquad status
```

---

## 📄 Arquivos de Referência

| Arquivo | Descrição |
|--------|----------|
| `opencode.md` | Este arquivo - Histórico completo |
| `GEMINI.md` | Diretrizes e regras do squad |
| `gimini.md` | Resumo para modelo |
| `.mcp.json` | Configuração MCP |

---

## 📋 Resumo das Atualizações (22 Abr 2026)

### O que foi feito:
1. ✅ Tor configurado e rodando (9050)
2. ✅ MCPs com proxy Tor
3. ✅ test_boot_brain.py criado (5/7 testes)
4. ✅ notify.sh copiado
5. ✅ Regra de economia adicionada aos prompts
6. ✅ Auditoria executada

### Arquivos novos:
- `scripts/notify.sh` - Telegram notifier
- `projetos/second-brain-obsidian/test_boot_brain.py` - Testes

### MCP Config (.mcp.json):
- filesystem (projeto)
- obsidian (vault)
- brave-search (Tor)
- fetch (Tor)

---

*Documento criado em: 2024-04-21*
*Última atualização: 2026-04-22*
*Próxima atualização: Ao final de cada sessão*
