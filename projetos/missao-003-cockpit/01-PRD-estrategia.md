---
mission_id: "#003"
title: "The Sovereign Cockpit - Interatividade e Controle Real"
status: "OPEN"
priority: "CRITICAL"
agente: "Estrategista"
data_criacao: "2026-04-21"
---

# PRD: Missão #003 - Do Monitor para o Cockpit

## 1. Visão Geral
Após o audit do Agente Cliente, ficou claro que um dashboard passivo não atende às necessidades de soberania. A Missão #003 focará em transformar a interface em um centro de comando interativo.

## 2. Objetivos Principais
- **Interatividade Agente-Usuário:** Permitir clicar em agentes para ver suas "entranhas" (prompts e logs recentes).
- **Controle de Sistema:** Adicionar botões de ação real (Emergency Stop, Refresh Health).
- **Filtros de Auditoria:** Implementar filtragem dinâmica no audit log (Error/Mission/System).
- **Destaque de Tarefa:** Centralizar a visualização da tarefa em execução para reduzir a carga cognitiva do usuário.

## 3. Requisitos Técnicos
- **R1: Server-Side Commands:** O monitor Python deve suportar o recebimento de comandos via pequenos arquivos de flag ou socket.
- **R2: UI Interativa:** Implementar event listeners no Dashboard para ações de clique.
- **R3: Traceability UI:** Mostrar o ID da missão em todos os logs gerados durante ela.

## 4. Evolução dos Prompts (Mestre de Prompts)
- Atualizar o `sentinel.md` e `monitor.md` para priorizar o feedback visual imediato ao usuário.

---
> **Audit Note:** Estrategista aprovou o plano. Arquiteto, sua vez de desenhar os controles.
