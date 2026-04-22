#!/usr/bin/env python3
"""
Módulo isolado do Notificador Telegram
Responsável por disparar alertas na esteira sem quebrar o loop central do AGENCIA01.
"""
import os
import requests
import logging

# Configuração simples de logging isolado para este projeto
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

class TelegramNotifier:
    """Implementa a integração amigável com o Telegram Bot."""
    
    def __init__(self, bot_token=None, chat_id=None):
        self.bot_token = bot_token or os.environ.get("TELEGRAM_BOT_TOKEN")
        self.chat_id = chat_id or os.environ.get("TELEGRAM_CHAT_ID")
        self.api_url = f"https://api.telegram.org/bot{self.bot_token}/sendMessage"

    def format_message(self, agent_name: str, step: str, details: str) -> str:
        """Formata o texto markdown com escape seguro se necessário (ou simples para testes atuais)."""
        msg = f"🚨 *Agencia Squad - Checkpoint*\n\n"
        msg += f"*Agente:* {agent_name}\n"
        msg += f"*Ação:* {step}\n"
        msg += f"*Detalhes:* {details}\n\n"
        msg += f"*Status:* Parado aguardando liberação."
        return msg

    def notify(self, agent_name: str, step: str, details: str) -> bool:
        """Envia a notificação HTTP e capta falhas sem propagar exceptions."""
        if not self.bot_token or not self.chat_id:
            logging.error("Telegram credentials missing (BOT_TOKEN ou CHAT_ID não encontrados).")
            return False

        text = self.format_message(agent_name, step, details)
        payload = {
            "chat_id": self.chat_id,
            "text": text,
            "parse_mode": "Markdown"
        }

        try:
            response = requests.post(self.api_url, json=payload, timeout=5)
            if response.status_code == 200:
                logging.info(f"Notificação enviada com sucesso para Checkpoint do agente '{agent_name}'.")
                return True
            else:
                logging.warning(f"Telegram API retornou status code {response.status_code}: {response.text}")
                return False
        except requests.exceptions.RequestException as e:
            logging.error(f"Erro de comunicação com o Telegram: {e}")
            return False

if __name__ == "__main__":
    # Teste de fumaça (somente roda se invocado diretamente e com as VARS contidas)
    # Exemplo: TELEGRAM_BOT_TOKEN="xxx" TELEGRAM_CHAT_ID="yyy" python3 notify_system.py
    notifier = TelegramNotifier()
    success = notifier.notify("TestAgent", "Teste de Isolamento", "Rodando script independente da engine.")
    print("Sucesso" if success else "Falha no envio ou variáveis não definidas")
