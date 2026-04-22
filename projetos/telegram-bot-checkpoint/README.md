# Telegram Bot Checkpoint - Sub-projeto Isolado

Este projeto foi desenvolvido como uma ferramenta auxiliar para o `AGENCIA01`, mas mantido de forma totalmente desacoplada e isolada para evitar divergências na lógica central da "Agência".

## 🛠️ Tecnologias
- **Python 3.x**
- **library:** `requests` (para comunicação HTTP)
- **Framework de Teste:** `unittest` + `unittest.mock`

## 📂 Estrutura do Projeto
- `notify_system.py`: Módulo principal contendo a lógica de envio de mensagens para o Telegram.
- `test_notify_system.py`: Suíte de testes unitários com mocks para garantir resiliência.
- `01-PRD-estrategia.md`: Requisitos e visão do produto.
- `02-Arquitetura.md`: Desenho técnico e diagramas.
- `03-Contrato-TDD.md`: Definição de comportamentos para desenvolvimento guiado por testes.

## 🧪 Como Rodar os Testes
Para garantir que o sistema está funcionando sem realizar chamadas reais à API do Telegram:
```bash
python3 test_notify_system.py
```

## 🚀 Como Integrar (Uso Externo)
O sistema foi desenhado para ser invocado via CLI ou importação de módulo. Para integração via CLI que respeita o isolamento:
```bash
TELEGRAM_BOT_TOKEN="seu_token" TELEGRAM_CHAT_ID="seu_id" python3 notify_system.py
```

## 📝 Future Tasks / Roadmap Interno
- [ ] **Sanitização de MarkdownV2:** Implementar escape rigoroso de caracteres especiais exigidos pelo Telegram para evitar erros de parse em mensagens complexas.
- [ ] **Suporte a Retry:** Adicionar lógica de re-tentativa exponencial para casos de erro `429 (Too Many Requests)`.
- [ ] **Anexos:** Suporte para envio de logs ou artefatos (.md, .py) diretamente via bot quando um Checkpoint for atingido.
- [ ] **Interação:** Possibilidade de responder à mensagem no Telegram para "Aprovar" ou "Reprovar" o checkpoint remotamente.

---
**Nota de Auditoria:** O código foi validado via TDD e Auditoria de Qualidade em 21/04/2026. Todos os testes passaram.
