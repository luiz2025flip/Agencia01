# Especificação de Produto (PRD): Second Brain dos Agentes (Obsidian Vault)

## 1. Visão Geral (Por Estrategista de Produto) 🧠
O objetivo deste projeto é dar **memória persistente local e isolada** para cada um dos agentes da Agência-Squad utilizando o ecossistema do **Obsidian** (arquivos Markdown locais).
Em vez de depender do limite de janela de contexto da LLM, os agentes lerão e escreverão em seus "cofres" (Vaults) do Obsidian via **Model Context Protocol (MCP - Filesystem)**. Cada agente terá seu próprio cérebro (pasta) e um cérebro compartilhado (pasta global) para passar o bastão.

## 2. Casos de Uso (Jornadas)
* **Onboarding do Agente:** Quando a Agência "acorda", o agente Arquitetura lê o `/Vault/Arquiteto/` para lembrar das decisões passadas.
* **Passagem de Bastão (Hand-off):** O Arquiteto salva o esquema YAML em `/Vault/Compartilhado/Especificacoes/`, onde o Coder Hacker usará como memória de trabalho.
* **Aprendizado do Agente FAQ:** O FAQ escreve erros solucionados em `/Vault/FAQ/Erros_Resolvidos/` automatizando o diário.

## 3. Escopo Funcional (Requisitos Core)
1. **Inicializador do Vault (Bootstrapper):** Um script Python automatizado que cria a hierarquia de pastas do Obsidian na máquina local se ela não existir.
2. **Setup do MCP (Filesystem):** Configuração oficial do `@modelcontextprotocol/server-filesystem` para que a IDE/Agentes interajam livremente com o Vault de `/home/kali/Documents/Agencia_Vault`.
3. **Padrão de Notas (YAML Frontmatter):** Toda anotação criada pelos agentes DEVE conter um Frontmatter rigoroso (metadados).

## 4. Estrutura Proposta do Cofre (Obsidian Vault)
```text
Agencia_SecondBrain/
├── 00_Central/
│   ├── Projetos_Ativos/
│   └── Logs_Orquestrador/
├── 10_Arquiteto/
│   └── ADRs_Aprendidos/
├── 20_CoderHacker/
│   ├── Snippets_Reutilizaveis/
│   └── Comandos_Favoritos/
├── 30_DevOps/
│   └── Infra_Config/
└── 40_FAQ/
    └── Base_Conhecimento/
```

## 5. Próximos Passos (Workflow)
* **Aguardando Cliente (Agente Cliente):** Validar se essa estrutura de pastas faz sentido do ponto de vista de atrito.
* **Aguardando Arquiteto:** Definir os schemas e bibliotecas Python exatas (Ex: `os`, `shutil` ou utilitários) para manipular os `.md`.
