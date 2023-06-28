# Git Templates

Inspired by https://www.tabnews.com.br/guscsales/uma-maneira-de-organizar-suas-branches-commits-e-pull-requests

## Branch Naming

```
# Template
I<issue-id>/<feature-sumary>

# Example:
git checkout -b I100/create-post-api
```

## Commit template

```
# Template
- type(?scope): description
- type(?scope): description

scope is optional

# Example:
feat(posts): creating hook to integrate with posts API
test: add missing tests for posts hook
```

- **feat**: Um novo recurso para a aplicação, e não precisa ser algo grande, mas apenas algo que não existia antes e que a pessoa final irá acessar.
- **fix**: Correções de bugs
- **docs**: Alterações em arquivos relacionados à documentações
- **style**: Alterações de estilização, formatação etc
- **refactor**: Um codigo de refatoração, ou seja, que foi alterado, que tem uma mudança transparente para o usuário final, porém uma mudança real para a aplicação
- **perf**: Alterações relacionadas à performance
- **test**: Criação ou modificação de testes
- **chore**: Alterações em arquivos de configuração, build, distribuição, CI, ou qualquer outra coisa que não envolva diretamente o código da aplicação para o usuário final

## Pull Request Title

```
# Padrão:
[I-<issue-id>] type(scope): description

# Exemplo:
[I-100] feat(posts): creating hook to integrate with posts API
```

## Pull Request Body

```
## Type of change

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Chore (documentation, packages, or tests updates, nothing that affect the final user directly)
- [ ] Release (new version of the application - only for production)

## Description

...

## Screenshots

...

## Tasks

- [task-id](task-link) or N/A

## Checklist

- [ ] My changes have less than or equal 400 lines
- [ ] I have performed a self-review of my own code
- [ ] The existing tests and linter pass locally with my changes
- [ ] I have commented my code in hard-to-understand areas (if applicable)
- [ ] I have created tests for my fix or feature (if applicable)

## Dependencies

This pull request has a dependency on the following others:

- link-to-depency PR or N/A
```
