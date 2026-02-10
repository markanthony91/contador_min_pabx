# Contador de Minutos PABX (FreePBX/Asterisk)

Este projeto visa centralizar e monitorar o tráfego de chamadas de 140 instâncias individuais de FreePBX operando em ambientes WSL (Windows 11).

## Objetivo
Mapear o volume de ligações (entrantes, saintes e internas) e o consumo de minutos em toda a rede de lojas, além de prover visibilidade em tempo real do status das chamadas.

## Stack Sugerida (Preliminar)
- **Coleta (Histórico):** Agente Go local (WSL) para extração de CDR via MySQL.
- **Coleta (Real-time):** Agente Go local (WSL) via AMI (Asterisk Manager Interface).
- **Transporte:** API REST ou InfluxDB Line Protocol.
- **Visualização:** Grafana ou Dashboard Centralizado em React.

## Estrutura do Projeto
- `gemini.md`: Brainstorming, arquitetura e plano detalhado.
- `scripts/`: Scripts de extração para as lojas.
- `server/`: Backend para consolidação dos dados.
