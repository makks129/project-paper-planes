# project-paper-planes
Project Paper Planes

### Context
Written using TDD with functional tests.

TODO where it's running

## Stack

#### App
- go / [gin-gonic](https://github.com/gin-gonic/gin) (go is chosen only to get more exp with the language)
- docker compose
- [air](https://github.com/cosmtrek/air) for Gin live-reload
  - see `.air.toml`

#### Tests
Unit + Functional tests
- go test
- [testify](https://github.com/stretchr/testify)
- Custom MySQL mock with docker

#### CI/CD
- pre-commit hooks with [pre-commit](https://pre-commit.com/)
