---
agente: "Arquiteto"
missao: "#004"
data: "2026-04-21"
---

# Arquitetura de Rastreabilidade Granular - Missão #004

## 1. Novo Esquema de Log (JSON Schema)

```json
{
  "timestamp": "2026-04-21T23:35:00.123Z",
  "project_id": "AG_01_SQUAD",
  "mission_id": "#004",
  "task": {
    "id": "TASK_001",
    "name": "Implementação do Modal de Agente",
    "status": "in_progress"
  },
  "telemetry": {
    "system": { "cpu": 12.5, "ram": 70.1 },
    "agents": [
      {
        "name": "arquiteto",
        "last_action": "Design Schema V2",
        "last_task_id": "TASK_001",
        "prompt_ref": "squads/agencia-squad/prompts/arquiteto.md"
      }
    ]
  }
}
```

## 2. Componente de Notificações (Toast Engine)
- **Local:** `dashboard/js/toast.js` (a ser criado)
- **Visual:** Alinhado ao topo-direito, vidro fosco, bordas coloridas por tipo (success, error, info).

## 3. Modal de Inspeção de Agente
- **Gatilho:** Clique no Agent Card.
- **Conteúdo:** 
    - Card expandido com estatísticas reais.
    - Bloco de código mostrando `prompt_file` (via fetch).
    - Histórico de `task_id` processadas por este agente.

---
> **Audit Note:** Planta finalizada. **TDD**, crie as validações de ID e Precisão.
