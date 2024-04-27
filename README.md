# Newsletter Sender

- Go web framework - Echo
- DB - Postgresql
- DB ORM - SQLC
- DB Driver - pgx
- CMD - Cobra Cli

## Require
- Go1.22
- Docker - https://www.docker.com/
- Taskfile - https://taskfile.dev/
- SQLC - https://sqlc.dev/
- Mockery - https://github.com/vektra/mockery
- Mail Provider - eg. google, mailgun

## Run Project
1. Init DB with taskfile - `task setup-db` (make sure your port is avaliable)
2. Create config.yaml
3. Migrate DB - `task run -- migrate up` or `go run . migrate up`
4. Generate sqlc, mockery - `task gen`
5. Serve Http Server with cobra command - `task run -- serveApi` or `go run . serveApi`
6. Endpoint subscribe - /api/v1/subscribers/subscribe
7. Endpoint unsubscribe - /api/v1/subscribers/unSubscribe 
8. Send Newsletter to active subscriber - `task run -- sendNewsletter --header yourheader --body yourbody`


## config.yaml
```yml
server:
  port: 8080
  allowOrigins:
    - "*"
  bodyLimit: "10M" # MiB
  timeout: 30 # Seconds
  logLevel: DEBUG

database:
  host: localhost
  port: 5432
  user: postgres
  password: 123456
  dbname: es-assignment
  sslmode: disable
  schema: public
  maxOpenConns: 10
  connectTimeout: 5 # Seconds
  readTimeout: 30 # Seconds

mailer:
  host: smtp.yourhost.org
  port: 587
  username: your@username.com
  password: yourpassword
  from: from@mail.com
```
