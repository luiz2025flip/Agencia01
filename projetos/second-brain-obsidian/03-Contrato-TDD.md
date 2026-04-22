# Contrato TDD / Behavior (Por Desenvolvedor TDD 🧪)

## Cenários BDD (Behavior-Driven Development)

**Funcionalidade:** Inicialização do Vault do Obsidian (Second Brain)
**Como:** Engine de Python (OS / Pathlib)
**Para:** Garantir que o Coder Hacker não crie lixo na máquina Kali.

```gherkin
Cenário: Configuração inicial limpa do cofre de conhecimento
  Dado que o diretório "/home/kali/Documents/Agencia_Vault" não existe no sistema
  Quando eu executo o script "boot_brain.py"
  Então as subpastas "10_Arquiteto", "20_CoderHacker", "30_DevOps", "40_FAQ", e "00_Central" devem ser criadas
  E um arquivo "00_Central/00_Index.md" deve ser gerado contendo o frontmatter YAML base
  E o script finaliza com "Saída de Sucesso 0"
```

```gherkin
Cenário: Tentativa de reinicialização sem destruição
  Dado que o diretório já existe e contém reflexões antigas dos agentes
  Quando eu executo o "boot_brain.py" novamente
  Então o script não deve sobrescrever os arquivos markdown que já estão lá
  E deve apenas validar a árvore e logar "Vault verificado e íntegro".
```

⚠️ O **Coder Hacker** só está autorizado a codificar seguindo EXATAMENTE essas amarras!
