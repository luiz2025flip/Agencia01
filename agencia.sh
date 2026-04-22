#!/bin/bash
# AGENCIA01 - Orquestrador Automático (otimizado)
# Executa pipeline completo de agentes automaticamente

AGENCIA_DIR="/home/kali/Área de trabalho/AGENCIA01"
cd "$AGENCIA_DIR" || exit 1

# Cores
VERDE='\033[0;32m'
AZUL='\033[0;34m'
AMARELO='\033[1;33m'
RESET='\033[0m'

log() { echo -e "${VERDE}[$(date +%H:%M)]${RESET} $1"; }

# Pipeline pré-definidos
case "$1" in
    api)
        log "🚀 Pipeline API REST..."
        log "1/5 PRD Spec..."
        opencode run "@prd_spec crie especificação API RESTful para CRUD de usuários"
        log "2/5 Arquiteto..."
        opencode run "@arquiteto defina arquitetura REST com DDD"
        log "3/5 Engineer..."
        opencode run "@engineer implemente API RESTful em Node.js"
        log "4/5 TDD..."
        opencode run "@tdd crie testes unitários para API"
        log "5/5 Auditor..."
        opencode run "@auditor revise o código"
        log "✅ API concluída!"
        ;;
    
    tarefa|task)
        shift
        log "🚀 Executando: $*"
        opencode run "$*"
        log "✅ Concluído!"
        ;;
    
    pipeline)
        log "🚀 Pipeline completo..."
        log "1/8 PRD..."
        opencode run "@prd_spec $2"
        log "2/8 Validação PRD..."
        opencode run "@validador_prd valide"
        log "3/8 Arquiteto..."
        opencode run "@arquiteto defina arquitetura"
        log "4/8 Validação Arquitetura..."
        opencode run "@validador_arquitetura valide"
        log "5/8 Engineer..."
        opencode run "@engineer implemente"
        log "6/8 TDD..."
        opencode run "@tdd crie testes"
        log "7/8 Validação TDD..."
        opencode run "@validador_tdd valide"
        log "8/8 Auditor..."
        opencode run "@auditor revise"
        log "✅ Pipeline completo!"
        ;;
    
    *)
        echo "AGENCIA Orquestrador"
        echo "================"
        echo ""
        echo "Uso: $0 <comando> [args]"
        echo ""
        echo "Comandos:"
        echo "  api                 - Pipeline API REST completa"
        echo "  tarefa <msg>        - Executa tarefa direta"
        echo "  pipeline <msg>       - Pipeline completo"
        echo ""
        echo "Exemplos:"
        echo "  $0 api"
        echo "  $0 tarefa @engineer crie função login"
        echo "  $0 pipeline criar sistema de auth"
        exit 1
        ;;
esac