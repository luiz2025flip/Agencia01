---
agente: "TDD"
missao: "#005"
data: "2026-04-21"
---

# Validação de Transparência (TDD) - Missão #005

## 1. Cenário: Segurança do Inspecionador
**Funcionalidade:** Leitura de Prompts

  **Cenário:** Acesso autorizado a arquivos .md
    Dado que eu requisite `GET /inspect?agent=arquiteto`
    Quando o servidor processar a requisição
    Então o conteúdo de `prompts/arquiteto.md` deve ser retornado
    E não deve ser possível ler arquivos fora da pasta `prompts/` (Path Traversal Test)

## 2. Cenário: Diferenciação Visual
  **Cenário:** Cores semânticas no log
    Dado que o log contenha uma tag "TASK"
    Quando renderizado no dashboard
    Então o elemento deve possuir a classe `tag-task`
    E a cor deve ser laranja (`#ff9f43`)

---
> **Audit Note:** Testes definidos. **Coder Hacker**, execute o upgrade final!
