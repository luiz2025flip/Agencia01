#!/bin/bash
# Multica Telegram Notifier
# Roda em background: ./notify.sh &

BOT_TOKEN="8755791756:AAGGvfIXkm1HpeZIsG-2goVq3PLTBaV91h8"
CHAT_ID="6092306332"

send() {
    curl -s -X POST "https://api.telegram.org/bot${BOT_TOKEN}/sendMessage" \
        -d "chat_id=${CHAT_ID}" \
        -d "text=$1" \
        -d "parse_mode=HTML"
}

# Check for new working tasks
check_tasks() {
    cd /home/kali/multica/server
    
    # Get tasks that just started working
    WORKING=$(./multica issue list --status in_progress 2>/dev/null | grep -o '"title":"[^"]*"' | head -5)
    
    if [ -n "$WORKING" ]; then
        send "🔄 <b>Tasks em progresso:</b>\n${WORKING}"
    fi
}

# Send startup message
send "✅ <b>Notificador Multica Ativado!</b>\n\nVou notifying sobre:\n• Novas tasks\n• Tasks bloqueadas\n• Tasks concluídas\n\n/start - Verificar agora"

# Loop principal
while true; do
    check_tasks
    sleep 60  # Check every minute
done