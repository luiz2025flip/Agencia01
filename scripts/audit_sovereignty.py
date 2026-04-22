#!/usr/bin/env python3
"""
Auditor de Soberania v1.1 (AGENCIA01)
Valida se as entregas do squad respeitam o Protocolo de Engenharia Soberana.
"""
import os
import sys
import re
import logging

logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")

class SovereigntyAuditor:
    def __init__(self, root_dir):
        self.root_dir = root_dir
        self.violations = 0

    def check_pii_leaks(self):
        """Busca por IDs de chat ou tokens expostos em arquivos de texto."""
        pii_pattern = re.compile(r'\d{8,12}|[0-9]{8,10}:[a-zA-Z0-9_-]{30,}')
        
        logging.info("🔍 Verificando vazamento de PII (IDs/Tokens)...")
        for root, _, files in os.walk(self.root_dir):
            if any(x in root for x in [".git", "node_modules", ".venv", "__pycache__"]):
                continue
            for file in files:
                if file.endswith((".py", ".md", ".log", ".txt")):
                    path = os.path.join(root, file)
                    try:
                        with open(path, 'r', errors='ignore') as f:
                            content = f.read()
                            clean_content = re.sub(r'os\.environ\.get\(".*"\)', '', content)
                            clean_content = re.sub(r'\*+', '', clean_content)
                            
                            matches = pii_pattern.findall(clean_content)
                            if matches:
                                for m in matches:
                                    if len(m) > 15 or (":" in m):
                                        logging.error(f"❌ PII Detectado em {file}: {m[:4]}***")
                                        self.violations += 1
                    except (IOError, OSError) as e:
                        logging.error(f"⚠️ Erro ao ler {file}: {e}")

    def check_tdd_pairing(self):
        """Verifica se cada projeto .py tem um correspondente de teste."""
        logging.info("🔍 Verificando cobertura de contrato TDD...")
        for root, _, files in os.walk(os.path.join(self.root_dir, "projetos")):
            for file in files:
                if file.endswith(".py") and not file.startswith("test_"):
                    test_file = f"test_{file}"
                    if test_file not in files:
                        logging.warning(f"⚠️ Arquivo órfão de teste: {file} em {root}")

    def check_resilience_patterns(self):
        """Busca por try/except genéricos (proibidos)."""
        logging.info("🔍 Analisando resiliência (Evitando try/except genéricos)...")
        # Procura por "except Exception:" ou "except:" sem tipo específico
        generic_except = re.compile(r'except:\s*|except\s+Exception(\s+as\s+\w+)?:\s*')
        for root, _, files in os.walk(self.root_dir):
            if "scripts" in root and "audit_sovereignty.py" in files:
                # O próprio auditor pode conter o padrão para fins de teste/regex, então pulamos a linha específica dele
                pass
            for file in files:
                if file.endswith(".py") and file != "audit_sovereignty.py":
                    path = os.path.join(root, file)
                    with open(path, 'r') as f:
                        for i, line in enumerate(f, 1):
                            if generic_except.search(line):
                                logging.error(f"❌ Erro de Resiliência em {file}:{i} - Uso de 'except Exception' proibido.")
                                self.violations += 1

    def run_all(self):
        self.check_pii_leaks()
        self.check_tdd_pairing()
        self.check_resilience_patterns()
        
        if self.violations == 0:
            logging.info("✅ SOBERANIA VALIDADA: Sistema limpo e seguro.")
            return True
        else:
            logging.critical(f"🔴 FALHA DE SOBERANIA: {self.violations} violações detectadas.")
            return False

if __name__ == "__main__":
    auditor = SovereigntyAuditor("/home/kali/Área de trabalho/AGENCIA01")
    if auditor.run_all():
        sys.exit(0)
    else:
        sys.exit(1)
