---
agente: "Arquiteto"
missao: "#005"
data: "2026-04-21"
---

# Arquitetura de Transparência Profunda - Missão #005

## 1. API de Inspeção (Sovereign Server V2)

O servidor deve suportar a rota seguinte:
- `GET /inspect?agent=nome_do_agente`
- **Lógica:**
    1. Verifica se o nome do agente existe no `squad.yml`.
    2. Lê o arquivo correspondente em `prompts/`.
    3. Retorna o conteúdo Markdown filtrado (OpSec).

## 2. Design do Modal de Inspeção
- **Estilo:** Glassmorphism profudo (`backdrop-filter: blur(20px)`).
- **Conteúdo:** Título do Agente, Tag de Função e um painel de visualização de código (Markdown).
- **Animação:** Fade-in suave com scale.

## 3. Esquema de Cores Semânticas (Logs)
- `.tag-proj`: color `#00f2ff` (Cyan)
- `.tag-miss`: color `#ff00ea` (Magenta)
- `.tag-task`: color `#ff9f43` (Orange)
- `.tag-done`: color `#3fb950` (Green)

---
> **Audit Note:** Planta aprovada. **TDD**, defina os contratos de segurança da API de inspeção.
