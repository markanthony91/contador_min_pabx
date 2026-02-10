# Brainstorming: Monitoramento Centralizado PABX (140 Lojas)

## 1. O Problema
- 140 instâncias isoladas de FreePBX (Asterisk) rodando em WSL.
- Falta de visibilidade consolidada sobre:
    - Total de minutos consumidos por loja/geral.
    - Relação entre chamadas atendidas e perdidas.
    - Tráfego interno (ramal-ramal) vs. Externo (Troncos).

## 2. Brainstorming & Insights

### Fontes de Dados (Asterisk)
O FreePBX armazena os registros no banco **MySQL/MariaDB**, especificamente na tabela `asteriskcdrdb.cdr`.
- **Campos Chave:** `calldate`, `src`, `dst`, `duration`, `billsec` (minutos cobráveis), `disposition` (ANSWERED, NO ANSWER, BUSY).

### Desafios do WSL no Windows 11
- **IP Dinâmico:** O WSL geralmente muda de IP após o reboot. Precisamos de um agente que "empurre" (push) os dados para um servidor central, em vez de um servidor tentando buscar (pull) de cada loja.
- **Recursos:** O script de coleta deve ser leve para não impactar a performance do PABX na loja.

### Estratégias de Coleta
1. **Agente Local (Go):** Um binário compilado que roda no Windows ou WSL. Ele se conecta ao banco de dados MySQL do FreePBX (porta 3306) para extrair os registros de chamadas (CDR) - **Foco em Auditoria e Minutos**.
2. **Monitoramento Real-time (AMI):** Interceptar eventos via Asterisk Manager Interface para dashboards ao vivo - **Foco em Status da Operação**.
3. **Identificação por IP:** Utilizaremos o IP fixo (`192.168.240.223`) como identificador inicial para os testes locais.

### Insights de Valor
- **Detecção de Ociosidade:** Identificar lojas que não estão fazendo/recebendo chamadas (possível falha técnica).
- **Ranking de Custos:** Identificar quais lojas gastam mais minutos para otimização de planos de telefonia.
- **Heatmap de Horários:** Entender os picos de atendimento para melhor escala de funcionários.
- **Painel em Tempo Real:** Ver quais ramais e lojas estão em chamada no exato momento.

## 3. Plano de Ação (Roadmap)

### Fase 1: Prova de Conceito (PoC) - Local
- [x] Definir linguagem (Go).
- [ ] Criar ferramenta de teste em Go para conectar no MySQL do PABX (`192.168.240.223`).
- [ ] Implementar listener AMI para eventos de chamada em tempo real.
- [ ] Validar extração dos campos CDR: `calldate`, `duration`, `billsec`, `disposition`.

### Fase 2: Arquitetura de Transporte
- [ ] Desenvolver uma API Central (Go ou Node.js) para receber os dados.
- [ ] Implementar autenticação básica ou Token por loja para segurança.
- [ ] Persistir os dados consolidados em um banco central (PostgreSQL ou SQLite).

### Fase 3: Visualização
- [ ] Configurar Dashboard no Grafana ou UI customizada em React.
- [ ] Criar filtros por Loja, Região (se aplicável) e Período.

### Fase 4: Deployment
- [ ] Criar um script de instalação automatizado (PowerShell/Shell script) para rodar nas 140 lojas.
- [ ] Configurar agendamento (Task Scheduler no Windows ou Crontab no WSL).

## 4. Próximos Passos Sugeridos
1. **Validar as credenciais:** Precisamos saber se temos acesso ao banco `asteriskcdrdb`.
2. **Definir Identificador:** Cada loja precisa de um ID único (ex: Loja001) para os dados não se misturarem.