# 0013 · Cofre de segredos: SOPS + age, entrega pelo Prumo e portão no sistema operacional

- **Status:** Proposto
- **Data:** 2026-11-16

## Contexto

O ADR-0003 decidiu o princípio: segredos nunca entram no contexto de um modelo, e quem sobe o ambiente é o Prumo, não o agente. Ficou em aberto onde os valores moram e como chegam aos containers sem passar pelo agente (#42).

Uma pesquisa aplicada, montando um ambiente local com agente de IA, Docker e vários repositórios, trouxe três fatos que mudam o desenho:

1. **A fronteira não é o cofre, é quem executa.** Se o segredo vira variável de ambiente num processo que o agente controla, um `env` ou `echo $X` mostra o valor. É a forma de vazamento mais citada em agentes de código[^1][^2][^3]. Guardar bem não basta; quem entrega precisa ser um processo que o agente não controla.
2. **Dá pra impor essa fronteira no sistema operacional, e não só por instrução.** O sandbox do Claude Code isola os comandos do agente com bubblewrap no Linux e no WSL2; com `sandbox.filesystem.denyRead`, o agente não lê a chave do cofre nem as credenciais da máquina[^4]. Docker é incompatível com esse sandbox, e a saída documentada é liberar comandos específicos com `sandbox.excludedCommands`[^4].
3. **Variável de ambiente no container também vaza.** Segredo passado por `environment:` aparece no `docker inspect`; passado por `secrets:` do Compose, vira arquivo em `/run/secrets/` e não aparece nas variáveis do container[^5].

Restrições: tudo livre (ADR-0005), rodando no computador de quem usa e sem servidor no C1 (ADR-0006).

## Decisão

Três peças, cada uma com um papel:

1. **Cofre: SOPS com age**[^6][^7]. Os valores reais de cada workspace ficam num arquivo criptografado na pasta de dados do Prumo (`~/.local/share/prumo/<workspace>/segredos/`, ADR-0003), fora do repositório. A chave do age fica no local padrão do usuário e nunca é montada no container nem lida pelo agente. Para ver e editar, `sops edit` abre o arquivo descriptografado no editor e criptografa de novo ao salvar.
2. **Porteiro: o próprio Prumo.** Só comandos do Prumo (`prumo ambiente subir`, `derrubar`, `logs`, `status`) descriptografam o cofre. Eles entregam os valores só aos serviços declarados na configuração do workspace, por `secrets:` do Compose, e devolvem a saída ao agente com os valores conhecidos mascarados, como os logs de CI fazem.
3. **Portão: isolamento do executor.** Cada adapter de executor (ADR-0007) declara como impede o agente de ler o cofre. No adapter do Claude Code: sandbox ligado, `denyRead` na chave do age, na pasta de segredos, em `~/.git-credentials`, `~/.docker` e `~/.ssh`, regra de permissão negando a ferramenta de leitura nesses caminhos (o sandbox cobre só comandos de terminal[^4]), `allowUnsandboxedCommands: false` e `excludedCommands` contendo só os comandos do porteiro. Docker não fica liberado para o agente; ele opera o ambiente só pelo Prumo.

Servidores de segredos (OpenBao, Infisical) ficam como adapters futuros, para quando o Prumo rodar numa VPS (ADR-0006), com o SOPS + age continuando como caminho nativo.

## Alternativas consideradas

| Opção | Licença | Servidor? | A favor | Contra |
| --- | --- | --- | --- | --- |
| SOPS + age | MPL-2.0 / BSD-3-Clause | Não | `sops exec-env` e `exec-file` entregam ao processo filho[^6]; texto criptografado legível por ferramenta comum | Sem interface gráfica; a chave do age vira o segredo a proteger |
| OpenBao[^8] | MPL-2.0 | Sim | Continuação livre do Vault; políticas, auditoria, segredos dinâmicos | Pesado para uma pessoa no próprio computador |
| Infisical[^9] | MIT (núcleo) | Sim (Postgres + Redis) | Interface web, `infisical run -- comando` | Mesmo peso de servidor |
| KeePassXC[^10] | GPL-3.0 | Não | Interface gráfica multiplataforma | A CLI pede a senha mestra a cada uso, então não serve como entregador automático; serve para guardar cópia da chave do age |
| Doppler, 1Password, Bitwarden Secrets Manager | Proprietários ou pagos | Nuvem | Já pensados para agentes | Fora do ADR-0005 |
| Só `.env` local (começo do ADR-0003) | — | Não | Simples | Texto puro no disco; qualquer leitura acidental expõe tudo |

## Consequências

- O agente consegue subir, derrubar e ler o estado do ambiente sem ver um valor sequer; vazamento passa a exigir contornar o isolamento, não um descuido.
- O `git push` deixa de funcionar para o agente, porque ele não lê as credenciais do git. Isso reforça o gate: entregar continua sendo ação humana.
- **Risco que fica:** a documentação do sandbox avisa que um comando liberado que executa arquivos graváveis pelo agente (um `compose.yml` dentro do worktree) pode ser alterado para extrair o segredo[^4]. Mitigações: segredos só para serviços declarados fora do repositório, e diff que toque arquivos de ambiente passa pelo gate antes de o ambiente subir. É uma barreira contra vazamento por descuido, não contra um agente tentando de propósito.
- O isolamento depende do agente: adapters sem sandbox equivalente (#13) precisam declarar isso no contrato, e o Prumo avisa quem usa.
- No Windows nativo o sandbox do Claude Code não existe[^4]; o runtime com isolamento passa a ser o WSL2 (#32).
- Exige um `.env.example` completo e a lista de serviços que recebem cada segredo na configuração do workspace.
- Perder a chave do age é perder o cofre: o guia de instalação precisa ensinar a fazer cópia dela (por exemplo, no KeePassXC).

## Fontes

[^1]: DZone, *Ways your AI coding agent exfiltrates secrets* · https://dzone.com/articles/ways-your-ai-coding-agent-exfiltrates-secrets
[^2]: Knostic, *Claude, Cursor and .env file secret leakage* · https://www.knostic.ai/blog/claude-cursor-env-file-secret-leakage
[^3]: Agent Patterns, *Secrets management for agents: credential injection* · https://www.agentpatterns.ai/security/secrets-management-for-agents/
[^4]: Anthropic, *Claude Code: sandboxing* · https://code.claude.com/docs/en/sandboxing
[^5]: Docker, *Manage secrets securely in Docker Compose* · https://docs.docker.com/compose/how-tos/use-secrets/
[^6]: SOPS, *Advanced usage* (`exec-env`, `exec-file`) · https://getsops.io/docs/usage/advanced/
[^7]: age · https://github.com/FiloSottile/age
[^8]: OpenBao · https://openbao.org
[^9]: Infisical · https://infisical.com
[^10]: KeePassXC, manual da CLI · https://man.archlinux.org/man/keepassxc-cli.1.en
