# Product Requirements Document (PRD) - Notificador Telegram (Checkpoint)

## 1. Visão Geral (Estrategista)
O projeto **Telegram Bot Checkpoint** visa integrar a esteira de orquestração do Agencia Squad (`opensquad`) com o Telegram do usuário. O principal objetivo é alertar o usuário (via push no celular) toda vez que a esteira parar em um "Checkpoint" aprovativo (ex: aguardando revisão de PRD, validação de Arquitetura ou Auditoria de Código).

## 2. Motivação e Problema
Atualmente, se o Squad é executado de forma paralela via processos longos, o usuário precisa ficar monitorando constantemente a IDE para saber se algum agente travou precisando de feedback humano. Com a notificação, damos "soberania assíncrona" ao Squad.

## 3. Escopo e Requisitos

### Requisitos Funcionais (RF)
- **RF01:** O sistema deve enviar uma mensagem para o Telegram do usuário (`chat_id`: 6092306332).
- **RF02:** A mensagem deve conter o nome do Agente que parou, a etapa do fluxo (ex: "Validação de Arquitetura") e um pequeno resumo do que precisa ser aprovado.
- **RF03:** Deve existir um script modular (`notify.py` ou `notify.sh`) que possa ser chamado facilmente por qualquer agente ou pelo framework do OpenSquad (via hooks).
- **RF04:** Utilizar o Token de Bot recém-criado (`8755791756:AAGGvfIXkm1HpeZIsG-2goVq3PLTBaV91h8`) guardado em variável de ambiente (não hardcoded no código final de produção).

### Requisitos Não Funcionais (RNF)
- **RNF01:** A integração deve ser assíncrona para não travar o loop principal do sistema.
- **RNF02:** O código deverá seguir restritamente as regras de **TDD**, com um `Contrato-TDD` acompanhando as definições antes que as rotinas de notificação sejam efetivamente criadas.
- **RNF03:** Os testes devem usar `mocks` para a API do Telegram para não realizar chamadas HTTP reais durante a auditoria local.

## 4. Entregáveis Esperados (Roadmap Spec-First)
1. **01-PRD-estrategia.md:** Este documento (Estrategista).
2. **02-Arquitetura.md:** Desenho do Fluxo, chamadas Webhook e modelagem JSON via Mermaid. (Arquiteto)
3. **03-Contrato-TDD.md:** Testes Gherkin simulando o sucesso e falha no envio. (Agente TDD)
4. **04-notify-system.py:** (Código final após auditoria).

---
*Status: Aguardando aprovação humana ou transição para o Arquiteto.*
