#!/usr/bin/env python3
import os
import sys
from pathlib import Path

# --- AGENTE CODER HACKER (PYTHON) ---
# Missão BDD: Garantir pastas do Obsidian e index inicial sem overwrite.
# Otimizado e direto ao ponto.

VAULT_BASE = Path("/home/kali/Documents/Agencia_Vault")

DIRECTORIES = [
    "00_Central",
    "10_Arquiteto",
    "20_CoderHacker",
    "30_DevOps",
    "40_FAQ",
    "99_Memos_Compartilhados"
]

YAML_FRONTMATTER = """---
agente: "Orquestrador"
data_atualizacao: "AUTO"
tags:
  - memoria_central
  - root
tipo: "index"
relevancia: alta
---

# SECOND BRAIN - AGÊNCIA SQUAD
Este cofre central armazena a consciência paralela dos Agentes.
Diretório mapeado via MCP (Filesystem).
"""

def boot_squad_vault():
    print("[*] Iniciando sequencia de Boot do Obsidian Vault...")
    
    # Valida e cria a base do Vault
    if not VAULT_BASE.exists():
        print(f"[+] Criando Vault master em {VAULT_BASE}")
        VAULT_BASE.mkdir(parents=True, exist_ok=True)
    
    # Constroi subpastas com segurança
    for dir_name in DIRECTORIES:
        agent_dir = VAULT_BASE / dir_name
        if not agent_dir.exists():
            agent_dir.mkdir(exist_ok=True)
            print(f"  [+] Subsetor criado: {dir_name}")
            
    # Cria o arquivo Index Mestre (Apenas se nao existir - respeitando TDD)
    index_file = VAULT_BASE / "00_Central" / "00_Index.md"
    if not index_file.exists():
        index_file.write_text(YAML_FRONTMATTER, encoding="utf-8")
        print("  [+] Index Mestre ancorado.")
    else:
        print("  [-] Index Mestre encontado. Overwrite bloqueado pela diretriz TDD.")

    print(f"\n[SUCESSO] Vault mapeado. MCP Server pronto para plugar em '{VAULT_BASE}'.")

if __name__ == "__main__":
    try:
        boot_squad_vault()
        sys.exit(0)
    except (OSError, PermissionError) as e:
        print(f"[ERRO DE SISTEMA] Falha ao manipular o Vault: {e}")
        sys.exit(1)
    except KeyboardInterrupt:
        print("\n[!] Operação cancelada pelo usuário.")
        sys.exit(0)

