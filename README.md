# pet project 1: simple auth service

## Getting started
1. mkdir go
2. cd go
3. git clone git@github.com:amenshenin/go_auth.git .
4. cd .devcontainer
5. cp .env.example .env
6. fill .env
7. docker compose up -d
8. docker exec -it ${COMPOSE_PROJECT_NAME}-app bash
9. nano /etc/hosts
10. Adds 127.0.0.1 ${VHOST}

## Run
1. docker exec -it ${COMPOSE_PROJECT_NAME}-app bash
2. go run cmd/main.go
