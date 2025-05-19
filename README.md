# 🌐 Site Status Monitor - Terminal

Um monitor de status de sites feito em Go, rodando via terminal. Ele verifica sites definidos em um arquivo `.txt`, registra o status com data e hora, e permite consultar o histórico de logs.

## ⚙️ Como funciona

1. Os sites são lidos do arquivo `listaSites.txt` (um por linha);
2. O programa faz requisições HTTP para cada site;
3. O status de cada site (online ou com erro) é exibido no terminal;
4. Cada resultado é registrado com data e hora no arquivo `logs.txt`;
5. É possível consultar os logs anteriores via menu.

## 📁 Estrutura esperada

- `main.go` → Código-fonte do projeto  
- `listaSites.txt` → Lista de sites a serem monitorados (um por linha)  
- `logs.txt` → Arquivo gerado com o histórico de status

## 🖥️ Menu de opções

- `1` → Iniciar monitoramento
- `2` → Exibir logs
- `3` → Sair

## ▶️ Como rodar

1. Certifique-se de ter o Go instalado;
2. Coloque os sites no `listaSites.txt`;
3. Compile e execute:

```bash
go run main.go
