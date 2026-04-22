import unittest
import os
from pathlib import Path
import shutil
from boot_brain import boot_squad_vault

class TestBootBrain(unittest.TestCase):
    def setUp(self):
        # Setup um diretório de teste provisório
        self.test_vault = Path("/tmp/Agencia_Vault_Test")
        if self.test_vault.exists():
            shutil.rmtree(self.test_vault)
        
        # Monkey patch do VAULT_BASE no boot_brain para não afetar o real
        import boot_brain
        self.original_vault = boot_brain.VAULT_BASE
        boot_brain.VAULT_BASE = self.test_vault

    def tearDown(self):
        # Restaura o original e limpa
        import boot_brain
        boot_brain.VAULT_BASE = self.original_vault
        if self.test_vault.exists():
            shutil.rmtree(self.test_vault)

    def test_vault_creation(self):
        """Valida se as pastas são criadas corretamente."""
        boot_squad_vault()
        
        expected_dirs = [
            "00_Central",
            "10_Arquiteto",
            "20_CoderHacker",
            "30_DevOps",
            "40_FAQ",
            "99_Memos_Compartilhados"
        ]
        
        for d in expected_dirs:
            self.assertTrue((self.test_vault / d).exists(), f"Pasta {d} não foi criada.")
            
        index_file = self.test_vault / "00_Central" / "00_Index.md"
        self.assertTrue(index_file.exists(), "Arquivo de índice não foi criado.")

    def test_no_overwrite_index(self):
        """Valida que o index não é sobrescrito se já existir."""
        self.test_vault.mkdir(parents=True)
        central = self.test_vault / "00_Central"
        central.mkdir()
        index_file = central / "00_Index.md"
        index_file.write_text("CONTEUDO ANTIGO")
        
        boot_squad_vault()
        
        self.assertEqual(index_file.read_text(), "CONTEUDO ANTIGO")

if __name__ == "__main__":
    unittest.main()
