---
agente: "TDD"
missao: "#004"
data: "2026-04-21"
---

# Validação de Rastreabilidade (TDD) - Missão #004

## 1. Cenário: Precisão de Horário e IDs
**Funcionalidade:** Auditoria Granular

  **Cenário:** Telemetria contém metadados obrigatórios
    Dado que o monitor gere um novo estado
    Quando o `status.json` for verificado
    Então o campo `project_id` deve ser `AG_01_SQUAD`
    E o campo `task.id` deve seguir o padrão `TASK_XXX`
    E o `timestamp` deve conter milissegundos (ISO 8601)

## 2. Cenário: Toast Notifications
  **Cenário:** Feedback de ação sem interrupção
    Dado que o usuário clique em um comando
    Quando a resposta do servidor chegar
    Então um elemento `.toast` deve surgir na tela
    E não deve haver o popup `alert()` nativo do navegador

---
> **Audit Note:** Contratos selados. **Coder Hacker**, hora do show. Implemente a Granularidade Triple-ID.
