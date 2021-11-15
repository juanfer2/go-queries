# go-app-graphql

This project if a copy of thrullo in this

- Run development project
```
fresh -c other_runner.conf
```
### Database ORM
This project use [buffalo-pop](github.com/gobuffalo/buffalo-pop/v2), [Soda CLI](https://gobuffalo.io/en/docs/db/toolbox) and [Fizz](https://github.com/gobuffalo/fizz)

- Config database 
```
soda g config
```
```yml
development:
  dialect: postgres
  database: myNameDatabase
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5

test:
  url: {{envOr "TEST_DATABASE_URL" "postgres://postgresUser:postgresPassword@127.0.0.1:5432/myNameDatabaseTest?sslmode=disable"}}

production:
  url: {{envOr "DATABASE_URL" "postgres://postgresUser:postgresPassword@127.0.0.1:5432/myNameDatabaseProducction?sslmode=disable"}}
```
- Create Database
Development

```
soda create -a
```
* Test
```
soda create -e test
```
* Producction
```
soda create -e producction
```

- Generate Migrations 
```
soda generate fizz name_of_migration
```

- Run Migrations 
```
soda migrate
```
or
```
soda migrate up
```

- Down Migrations 
```
soda migrate down
```
### Golang console
This project use [gore](https://github.com/motemen/gore)

- Run console
```
gore -autoimport
```
