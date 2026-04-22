# AGENCIA01 - OpenSquad Setup

Este projeto está configurado com o **OpenSquad** para orchestrar seus 16 agentes de IA.

## Agentes Configurados

| Agente | Função |Checkpoint|
|--------|--------|----------|
| arquiteto | Arquiteto de Software | ✓ |
| prd_spec | Especificador de PRD | ✓ |
| engineer | Engenheiro de Software | - |
| tdd | Desenvolvedor TDD | - |
| validador_prd | Validador de PRD | ✓ |
| validador_arquitetura | Validador de Arquitetura | ✓ |
| validador_tdd | Validador de TDD | ✓ |
| auditor_codigo | Auditor de Código | ✓ |
| documentador | Documentador | - |
| estrategista | Estrategista | ✓ |
| recrutador | Recrutador/Orquestrador | ✓ |
| faq | Especialista FAQ | - |
| gestor | Gestor de Projeto | ✓ |
| sentinel | Sentinel (Segurança) | ✓ |
| monitor | Monitor de Qualidade | - |
| mestre_prompt | Mestre de Prompts | - |

## Como Usar

### 1. Abra o projeto no OpenCode
```bash
cd /home/kali/Área\ de\ trabalho/AGENCIA01
opencode .
```

### 2. Use os comandos do OpenSquad
No terminal do OpenCode:

```bash
# Verificar squads disponíveis
/opensquad

# Rodar o squad
/opensquad run agencia-squad

# Ou criar um novo squad
/opensquad create "meu novo squad"
```

### 3. Configuração do Modelo
O OpenSquad usará automaticamente o modelo configurado no OpenCode (big-pickle).

## Estrutura do Projeto

```
AGENCIA01/
├── squads/
│   └── agencia-squad/
│       ├── squad.yml          # Definição do squad
│       └── prompts/           # Prompts de cada agente
│           ├── arquiteto.md
│           ├── engineer.md
│           ├── tdd.md
│           └── ...
├── _opensquad/                # Configurações internas
└── .env                      # Variáveis de ambiente
```

## Checkpoints

Agentes com ✓ pedem aprovação antes de prosseguir. Isso permite que você revise e aprove cada etapa do processo.

## Nota sobre OpenCode

Este projeto está configurado para usar **OpenCode** com o modelo **big-pickle** (gratuito), exatamente como no Multica.
