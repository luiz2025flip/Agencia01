---
agente: "TDD"
missao: "#002"
data: "2026-04-21"
---

# Contrato de Testes (TDD) - Missão #002

## 1. Cenário: Validação do Script de Monitoramento
**Funcionalidade:** Monitoramento de Saúde do Sistema

  **Cenário:** O script gera JSON válido
    Dado que o script `scripts/monitor_system.py` seja executado
    Quando ele concluir um ciclo de coleta
    Então o arquivo `dashboard/data/status.json` deve existir
    E o conteúdo deve ser um JSON válido
    E não deve conter caminhos como `/home/kali/`

  **Cenário:** Detecção do Tor
    Dado que o serviço Tor esteja rodando na porta 9050
    Quando o monitor verificar as métricas
    Então o campo `system.tor` no JSON deve ser `true`

## 2. Cenário: Performance do Frontend
  **Cenário:** Dashboard consome dados sem erro
    Dado que o Dashboard seja carregado
    Quando o `main.js` realizar o fetch do `status.json`
    Então os elementos de CPU e RAM no HTML devem ser atualizados
    E não deve haver mensagens de erro no console

## 3. Unit Tests (Proposta)
- `test_json_structure()`: Valida chaves obrigatórias.
- `test_opsec_filter()`: Garante que strings proibidas sejam filtradas.

---
> **Audit Note:** Contrato TDD selado. Passando a bola para o **Coder Hacker** para implementação.
