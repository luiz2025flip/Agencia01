#!/usr/bin/env python3
import unittest
from unittest.mock import patch
import os
from notify_system import TelegramNotifier

class TestTelegramNotifier(unittest.TestCase):

    def setUp(self):
        # Configura as variaveis "mock" de ambiente essenciais da classe
        self.notifier = TelegramNotifier(bot_token="fancymocktoken", chat_id="123456789")

    @patch('notify_system.requests.post')
    def test_successful_notification(self, mock_post):
        # Simula resposta Positiva do Telegram (200 OK)
        mock_post.return_value.status_code = 200
        mock_post.return_value.json.return_value = {"ok": True}

        result = self.notifier.notify(agent_name="Engenheiro", step="Revisão Código", details="Aguardando liberação do PR")
        self.assertTrue(result)
        mock_post.assert_called_once()
    
    @patch('notify_system.requests.post')
    def test_failed_http_notification(self, mock_post):
        # Simula erro de retorno (Ex: token inválido retornado via json mas com erro 401)
        mock_post.return_value.status_code = 401

        result = self.notifier.notify(agent_name="Estrategista", step="Ideação", details="Token com erro")
        self.assertFalse(result)
        
    @patch('notify_system.requests.post')
    def test_connection_timeout_handling(self, mock_post):
        # Simula que a internet caiu e explodiu o requests
        import requests
        mock_post.side_effect = requests.exceptions.ReadTimeout

        # A execução da Engine NÃO DEVE morrer aqui e sim retornar False elegantemente
        result = self.notifier.notify(agent_name="TDD", step="Validação", details="Sem internet")
        self.assertFalse(result)

    def test_missing_credentials(self):
        # Testando quando falta .env, não deve mandar request 
        bad_notifier = TelegramNotifier(bot_token=None, chat_id=None)
        with patch('notify_system.requests.post') as mock_post:
            result = bad_notifier.notify("A", "B", "C")
            self.assertFalse(result)
            mock_post.assert_not_called()

if __name__ == '__main__':
    unittest.main()
