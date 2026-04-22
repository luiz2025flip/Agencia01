#!/bin/bash
# AGENCIA01 - Orquestrador Automático de Agentes
# Executa múltiplos agentes em sequência automaticamente

AGENCIA_DIR="/home/kali/Área de trabalho/AGENCIA01"
cd "$AGENCIA_DIR" || exit 1

# Cores
VERDE='\033[0;32m'
AZUL='\033[0;34m'
AMARELO='\033[1;33m'
RESET='\033[0m'

log() { echo -e "${VERDE}[$(date +%H:%M)]${RESET} $1"; }
info() { echo -e "${AZUL}[INFO]${RESET} $1"; }
warn() { echo -e "${AMARELO}[WARN]${RESET} $1"; }

# Tarefas pré-definidas
case "$1" in
    api)
        log "🚀 Criando API REST completa..."
        opencode run "@prd_spec crie especificação API de users" || warn "Erro no PRD"
        opencode run "@arquiteto defina arquitetura REST" || warn "Erro no Arquiteto"
        opencode run "@engineer implemente API de users" || warn "Erro no Engineer"
        opencode run "@tdd crie testes para API" || warn "Erro no TDD"
        opencode run "@auditor revise o código" || warn "Erro no Auditor"
        log "✅ API concluída!"
        ;;
    
    feature)
        log "🚀 Criando nova feature..."
        opencode run "@prd_spec crie especificação para: $2" || warn "Erro"
        opencode run "@arquiteto defina arquitetura" || warn "Erro"
        opencode run "@engineer implemente $2" || warn "Erro"
        opencode run "@tdd crie testes" || warn "Erro"
        log "✅ Feature concluída!"
        ;;
    
    teste)
        log "🚀 Criando testes..."
        opencode run "@tdd crie testes para: $2" || warn "Erro"
        opencode run "@validador_tdd valide testes" || warn "Erro"
        log "✅ Testes concluídos!"
        ;;
    
    revisar)
        log "🚀 Revisando código..."
        opencode run "@auditor revise o código" || warn "Erro"
        log "✅ Revisão concluída!"
        ;;
    
    docs)
        log "🚀 Criando documentação..."
        opencode run "@documentador documente o projeto" || warn "Erro"
        log "✅ Documentação concluída!"
        ;;
    
    completa)
        log "🚀 Execução completa..."
        opencode run "@prd_spec crie especificação" || warn "Erro"
        opencode run "@validador_prd valide" || warn "Erro"
        opencode run "@arquiteto defina arquitetura" || warn "Erro"
        opencode run "@validador_arquitetura valide" || warn "Erro"
        opencode run "@engineer implemente" || warn "Erro"
        opencode run "@tdd crie testes" || warn "Erro"
        opencode run "@validador_tdd valide" || warn "Erro"
        opencode run "@auditor revise" || warn "Erro"
        log "✅ Execução completa!"
        ;;
    
    *)
        if [ -z "$1" ]; then
            echo "Uso: $0 <comando> [args]"
            echo ""
            echo "Comandos disponíveis:"
            echo "  api            - Cria API REST completa"
            echo "  feature <msg>  - Cria nova feature"
            echo "  teste <msg>   - Cria testes"
            echo "  revisar       - Revisa código"
            echo "  docs         - Cria documentação"
            echo "  completa     - Execução full pipeline"
            echo ""
            echo "Agentes diretos:"
            echo "  run @arquiteto <tarefa>"
            echo "  run @engineer <tarefa>"
            exit 1
        else
            log "Executando agente: $@"
            opencode run "$@"
        fi
        ;;
esac