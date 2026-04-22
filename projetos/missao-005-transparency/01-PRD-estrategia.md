---
mission_id: "#005"
title: "Deep Transparency & Semantic Audit Log"
status: "OPEN"
priority: "CRITICAL"
agente: "Estrategista"
data_criacao: "2026-04-21"
---

# PRD: Missão #005 - Transparência Profunda e Auditoria Semântica

## 1. Visão Geral
Esta missão é o polimento final da UI/UX focada no Agente Cliente. O objetivo é remover as "caixas pretas" do sistema, permitindo que o usuário inspecione a lógica dos agentes (prompts) e identifique eventos instantaneamente através de cores semânticas.

## 2. Objetivos Principais
- **Semantic Logging:** Implementar cores distintas no Audit Log:
    - `PROJ`: Cyan (Soberania)
    - `MISS`: Magenta (Ação)
    - `TASK`: Orange (Execução)
    - `DONE`: Green (Sucesso)
- **Safe Prompt Viewer:** Criar um endpoint e modal para visualizar os arquivos `.md` de prompts dos agentes.
- **OpSec Audit:** O visualizador de prompts deve garantir que variáveis sensíveis (se houver) não sejam expostas se forem injetadas em tempo real.

## 3. Requisitos Técnicos
- **R1: Server Endpoint:** `sovereign_server.py` deve permitir leitura de arquivos em `squads/agencia-squad/prompts/`.
- **R2: Glass Modal UI:** Implementar o modal no Dashboard.
- **R3: Enhanced CSS:** Novas classes de cores para o log.

---
> **Audit Note:** Missão iniciada. Arquiteto, defina as rotas de inspeção.
