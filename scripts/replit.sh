#!/bin/bash
# Replit Agent via API
# Uso: ./replit.sh <comando>

REPLIT_API="https://replit.com/ajax/2024-09-01"

case "$1" in
    login)
        echo "Abra: https://replit.com/login"
        echo "Depois auth via OAuth"
        ;;
    list)
        echo "Listando repls..."
        curl -s "$REPLIT_API/repls" -H "Cookie: $(cat ~/.replit-cookie 2>/dev/null)"
        ;;
    run)
        shift
        echo "Executando em $1..."
        # Via URL: replit.com/@user/repl
        ;;
    *)
        echo "Uso: $0 <comando>"
        echo ""
        echo "Notas:"
        echo "- Replit MCP precisa OAuth"
        echo "- Para usar: configure em replit.com → Account → MCP"
        ;;
esac