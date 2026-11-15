# 0003 · Dado sensível nunca passa pela IA

- **Status:** Proposto
- **Data:** 2026-11-15

## Contexto
O agente de IA lê arquivos, roda comandos e manda texto pra um provedor externo. Tudo que ele lê pode sair da máquina. Código pode ir: é o trabalho dele. Segredos (senhas, tokens, chaves), dados pessoais e dados de terceiros não podem ir, nem por descuido. Ao mesmo tempo, o projeto precisa desses segredos pra subir o ambiente e rodar os testes.

## Decisão

### A regra
Segredos, dados pessoais e dados de terceiros **nunca entram no contexto de um modelo**. Código, documentação e dados fictícios de teste podem entrar.

### Como é garantido
1. **O agente não enxerga o que não precisa.** O container do agente (ADR-0007) monta só o worktree da tarefa, sem `~/.ssh`, sem `.env` e sem a pasta de dados do Prumo.
2. **`.prumoignore`**: lista de caminhos que nenhuma etapa lê nem recebe como entrada, no formato do `.gitignore`. Vem com um padrão (`.env*` exceto `.env.example`, `*.pem`, `*.key`, `secrets/`, dumps de banco) e cada projeto acrescenta o seu.
3. **Diff que toca caminho ignorado é recusado** no gate, antes do commit.

### Variáveis de ambiente do projeto
- No repositório só entra o **`.env.example`**, com nomes e valores falsos. A IA pode ler esse arquivo.
- Os **valores reais** ficam na pasta de dados do Prumo, por workspace (`~/.local/share/prumo/<workspace>/segredos/`), fora do repositório e fora do container do agente.
- **Quem sobe o ambiente é o Prumo, não o agente.** O Prumo roda o `docker compose` do projeto passando os valores (`env_file` ou `secrets:` do Compose), e só os containers da aplicação recebem. O agente sabe que `DATABASE_URL` existe, mas não vê o valor; quando precisa do ambiente no ar, pede ao Prumo.
- O **cofre de segredos** (onde e como guardar os valores além de um arquivo local) fica pra um ADR próprio, depois do estudo da issue #42.

## Consequências
- Rodar o projeto com IA exige um `.env.example` completo. É uma boa prática de qualquer forma.
- O Prumo precisa de um comando pra subir e derrubar o ambiente do projeto, já que o agente não faz isso sozinho.
- O `.prumoignore` é uma camada a mais, não a única: o que protege de verdade é o agente não ter acesso ao arquivo.
- Testes que precisam de dado parecido com o real usam dados gerados (fictícios), nunca uma cópia da produção.

## Fontes
- OWASP, *Secrets Management Cheat Sheet* · https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html
- Docker, *Manage secrets securely in Docker Compose* · https://docs.docker.com/compose/how-tos/use-secrets/
- The Twelve-Factor App, *III. Config* · https://12factor.net/config
- LGPD (Lei 13.709/2018), tratamento de dados pessoais
