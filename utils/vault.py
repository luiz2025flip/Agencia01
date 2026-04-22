import base64
import os
import json
import time
from datetime import datetime
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

class SecretVault:
    """
    Sovereign Secret Vault V2 (Production Grade)
    Uses AES-GCM for authenticated encryption.
    Includes Audit Logging and Triple-ID Protocol support.
    """
    def __init__(self, master_key: str = None):
        self.master_key = master_key or os.getenv("VAULT_MASTER_KEY", "agencia_soverign_2026_default")
        self.key = hashlib.sha256(self.master_key.encode()).digest()
        self.aesgcm = AESGCM(self.key)
        self.audit_log_path = "/home/kali/Área de trabalho/AGENCIA01/scripts/logs/vault_audit.log"
        os.makedirs(os.path.dirname(self.audit_log_path), exist_ok=True)

    def _log(self, action: str, secret_name: str, success: bool):
        entry = {
            "timestamp": datetime.now().isoformat(),
            "action": action,
            "secret": secret_name,
            "success": success
        }
        with open(self.audit_log_path, "a") as f:
            f.write(json.dumps(entry) + "\n")

    def encrypt(self, secret_name: str, data: str) -> str:
        try:
            nonce = os.urandom(12)
            ciphertext = self.aesgcm.encrypt(nonce, data.encode(), None)
            token = base64.b64encode(nonce + ciphertext).decode()
            self._log("ENCRYPT", secret_name, True)
            return token
        except Exception:
            self._log("ENCRYPT", secret_name, False)
            raise

    def decrypt(self, secret_name: str, token: str) -> str:
        try:
            data = base64.b64decode(token)
            nonce = data[:12]
            ciphertext = data[12:]
            val = self.aesgcm.decrypt(nonce, ciphertext, None).decode()
            self._log("DECRYPT", secret_name, True)
            return val
        except Exception:
            self._log("DECRYPT", secret_name, False)
            raise

import hashlib
if __name__ == "__main__":
    vault = SecretVault()
    t = vault.encrypt("DB_PASSWORD", "admin123")
    print(f"Encrypted: {t}")
    print(f"Decrypted: {vault.decrypt('DB_PASSWORD', t)}")
