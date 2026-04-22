# Arquitetura e Modelagem - Notificador Telegram (Checkpoint)

## 1. Visão Arquitetural (Arquiteto)
O objetivo do sistema é atuar como uma camada de webhook/API de notificação para a esteira do OpenSquad. Quando um checkpoint é alcançado e requer aprovação humana, a engine subjacente invocará nosso módulo `notify_system.py`, o qual fará uma requisição segura e assíncrona para a API Oficial do Telegram, enviando a mensagem descritiva para o dispositivo do usuário.

## 2. Diagrama de Fluxo (Mermaid)
```mermaid
sequenceDiagram
    participant OS as OpenSquad / Engine
    participant Ag as Agente (Ex: Validador PRD)
    participant NS as notify_system.py
    participant TG as Telegram API
    participant User as Dispositivo do Usuário

    OS->>Ag: Executa tarefa do Agente
    activate Ag
    Ag-->>OS: Retorna Status: CHECKPOINT REQUIRED
    deactivate Ag
    
    OS->>NS: call notify(agent_name, step, details)
    activate NS
    NS->>NS: Formata a mensagem com MarkdownV2
    NS->>NS: Captura BOT_TOKEN e CHAT_ID do .env
    NS->>TG: HTTP POST /sendMessage
    activate TG
    TG-->>NS: 200 OK (Message Data JSON)
    deactivate TG
    
    NS-->>OS: Retorno de Sub-processo (Sucesso)
    deactivate NS
    
    TG->>User: Push Notification Recebida!
    OS->>OS: Sistema Pausa aguardando input/aprovação
```

## 3. Integração (Schema do Payload Telegram)
O script Python deve construir o seguinte Payload (JSON) no Post para a URL:
`https://api.telegram.org/bot<TOKEN>/sendMessage`

```json
{
  "chat_id": 6092306332,
  "text": "🚨 *Agencia Squad - Checkpoint*\n\n*Agente:* Validador PRD\n*Ação:* Aguardando revisão humana do arquivo spec.md.\n*Status:* Parado.",
  "parse_mode": "MarkdownV2"
}
```

## 4. Estruturação do Código Base (Blueprint Python)

O arquivo principal será `notify_system.py` contendo uma classe orientada a objetos (ou uma função pura modular) chamada `TelegramNotifier`, que aceitará os seguintes parâmetros:
- `agent_name` (string)
- `step` (string)
- `details` (string)

**Dependências externas da linguagem:**
- Biblioteca padrão `urllib.request` ou `requests` para requisição HTTP externa. (A favor de menos dependências de pacotes, podemos usar a biblioteca nativa, mas o padrão será sugerido via `requests` caso seja aprovado). Vamos assumir o uso de `requests` via `pip`.

---
*Status: Aguardando Agente TDD para a elaboração do esquema de testes comportamentais.*
